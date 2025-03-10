package evm

import (
	"errors"
	"log"
	"open-payment-gateway/pkg/eventbus"
	"strings"
	"sync"
	"time"
)

type EvmListenerConfig struct {
	Quitch           chan struct{}
	Wg               *sync.WaitGroup
	Network          Network
	AddressStore     AddressStore
	BlockStore       BlockStore
	TransactionStore TransactionStore
	Notification     eventbus.InternalNotification
	Provider         EvmProvider
	WaitForNewBlock  time.Duration
}

type EvmListener struct {
	Config *EvmListenerConfig
}

func NewEvmListener(config *EvmListenerConfig) *EvmListener {
	return &EvmListener{
		Config: config,
	}
}

func (l *EvmListener) Start() {

BlockIterator:
	for {
		select {
		case <-l.Config.Quitch:
			log.Print("Received stop signal")

			break BlockIterator

		default:
			latestProcessedBlockNumber, err := l.Config.BlockStore.GetLatestProcessedBlockNumber()
			if err != nil {
				log.Fatal("Could not get the latest processed block number from database")
			}

			// Check if we need to skip some blocks if the starting block number is not equal to -1
			if l.Config.Network.StartingBlockNumber > -1 && l.Config.Network.StartingBlockNumber > latestProcessedBlockNumber {
				latestProcessedBlockNumber = l.Config.Network.StartingBlockNumber
			}

			log.Printf("latest Processes block number: %d\n", latestProcessedBlockNumber)
			latestBlockNumber, err := l.Config.Provider.GetLatestBlockNumber()
			if err != nil {
				log.Fatal("Could not get the latest block number from provider")
			}

			log.Printf("latest Block Number is %d\n", latestBlockNumber)

			if latestBlockNumber > latestProcessedBlockNumber {
				// Iterate through blocks and process them
				processingBlockNumber := latestProcessedBlockNumber + 1
				log.Printf("Processing block %d\n", processingBlockNumber)
				processingBlock, err := l.Config.Provider.GetBlockByNumber(processingBlockNumber)

				if err != nil {
					log.Fatal("Could not get block data from provider")
				}

				if err := ProcessBlock(l.Config.Notification, l.Config.AddressStore, l.Config.TransactionStore, processingBlock); err != nil {
					log.Fatal("Could not process block")
				}

				if err := l.Config.BlockStore.SaveBlock(&processingBlock); err != nil {
					log.Fatal("Could not save processed block into the database")
				}

			} else {
				log.Println("Waiting for new blocks to be mined")
				time.Sleep(l.Config.WaitForNewBlock)
			}
		}
	}

	// Removing goroutine from wait group
	l.Config.Wg.Done()
	// Sending a signal to the quit channel indicating we've exited the loop
	l.Config.Quitch <- struct{}{}
}

func (l *EvmListener) Stop() bool {
	// Sending quit signal to the listener goroutine
	l.Config.Quitch <- struct{}{}
	// Waiting for the listener to exit the for loop
	<-l.Config.Quitch
	return true
}

func ProcessBlock(notification eventbus.InternalNotification, addressStore AddressStore, transactionStore TransactionStore, b Block) error {
	transactions := b.Transactions
	log.Printf("Processing block %d\n", b.BlockNumber)

	for _, t := range transactions {
		err := ProcessTransaction(notification, addressStore, transactionStore, t)
		if err != nil {
			return err
		}
	}

	return nil
}

func ProcessTransaction(notification eventbus.InternalNotification, addressStore AddressStore, transactionStore TransactionStore, tx Transaction) error {
	if tx.To == "" {
		return nil
	}

	txType, err := GetTransactionType(addressStore, tx)
	if err != nil {
		return err
	}

	if txType != "" {
		tx.TxType = txType
		weiValue, ok := StringToBigInt(tx.Value)
		if !ok {
			return errors.New("could not convert string to bigint wei")
		}
		tx.Value = WeiToEther(weiValue).String()
		err := transactionStore.SaveTransaction(&tx)
		if err != nil {
			return err
		}

		log.Printf("Received Transaction of type %s from %s to %s with the value of %s Ether\n", txType, tx.From, tx.To, tx.Value)
		n, err := eventbus.NewTransactionNotification{
			BlockNumber: tx.BlockNumber,
			BlockHash:   tx.BlockHash,
			Network:     tx.Network,
			Currency:    tx.Currency,
			TxHash:      tx.TxHash,
			TxType:      tx.TxType,
			Value:       tx.Value,
			From:        tx.From,
			To:          tx.To,
		}.ToJSON()

		if err != nil {
			return err
		}
		// NOT IMPLEMENTED
		err = notification.Notify("TRANSACTION_DETECTED", n)
		if err != nil {
			log.Printf("Failed to send notification to notification service: %+v", err)
		} else {
			transactionStore.UpdateBroadcasted(tx.TxHash, true)
		}

	}
	return nil
}

func GetTransactionType(addressStore AddressStore, tx Transaction) (string, error) {
	isDeposit, err := IsDepositTransaction(addressStore, tx.From, tx.To)
	if err != nil {
		return "", err
	}
	isWithdrawal, err := IsWithdrawalTransaction(addressStore, tx.From, tx.To)
	if err != nil {
		return "", err
	}

	if isDeposit && isWithdrawal {
		if tx.From == tx.To {
			return "self", nil
		}

		return "internal", nil
	}
	if isDeposit {
		return "deposit", nil
	}
	if isWithdrawal {
		return "withdrawal", nil
	}

	return "", nil
}

func IsDepositTransaction(addressStore AddressStore, from string, to string) (bool, error) {
	exists, err := addressStore.AddressExists(strings.ToLower(to))
	if err != nil {
		return false, err
	}
	return exists, nil
}

func IsWithdrawalTransaction(addressStore AddressStore, from string, to string) (bool, error) {
	exists, err := addressStore.AddressExists(strings.ToLower(from))
	if err != nil {
		return false, err
	}
	return exists, nil
}
