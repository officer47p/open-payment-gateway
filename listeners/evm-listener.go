package listeners

import (
	"errors"
	"fmt"
	"math/big"
	"open-payment-gateway/db"
	"open-payment-gateway/internal_notification"
	"open-payment-gateway/providers"
	"open-payment-gateway/types"
	"open-payment-gateway/utils"
	"strings"
	"sync"
	"time"
)

type EvmListenerConfig struct {
	Quitch              chan struct{}
	Wg                  *sync.WaitGroup
	Network             types.Network
	StartingBlockNumber int64
	AddressStore        db.AddressStore
	BlockStore          db.BlockStore
	TransactionStore    db.TransactionStore
	Notification        internal_notification.InternalNotification
	Provider            providers.EvmProvider
	WaitForNewBlock     time.Duration
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
			fmt.Println("Received stop signal")

			break BlockIterator

		default:
			fmt.Println("In the default")
			latestProcessedBlockNumber, err := l.Config.BlockStore.GetLatestProcessedBlockNumber()
			if err != nil {
				panic("Could not get the latest processed block number from database")
			}

			// Check if we need to skip some blocks if the starting block number is not equal to -1
			if l.Config.StartingBlockNumber > -1 && l.Config.StartingBlockNumber > latestProcessedBlockNumber {
				latestProcessedBlockNumber = l.Config.StartingBlockNumber
			}

			fmt.Printf("latest Processes Block Number is %d\n", latestProcessedBlockNumber)

			latestBlockNumber, err := l.Config.Provider.GetLatestBlockNumber()
			if err != nil {
				panic("Could not get the latest block number from provider")
			}

			fmt.Printf("latest Block Number is %d\n", latestBlockNumber)

			if latestBlockNumber > latestProcessedBlockNumber {
				// Iterate through blocks and process them
				// fmt.Println("We're behind")

				processingBlockNumber := latestProcessedBlockNumber + 1
				fmt.Printf("Processing Block Number is %d\n", processingBlockNumber)
				processingBlock, err := l.Config.Provider.GetBlockByNumber(processingBlockNumber)

				if err != nil {
					panic("Could not get block data from provider")
				}

				// TODO: Wrap ProcessBlock and SaveBlock in a database transaction
				if err := ProcessBlock(l.Config.Notification, l.Config.AddressStore, l.Config.TransactionStore, processingBlock); err != nil {
					panic("Could not process block")
				}

				if err := l.Config.BlockStore.SaveBlock(&processingBlock); err != nil {
					panic("Could not save processed block into the database")
				}

			} else {
				//  Wait for new blocks to be mined
				time.Sleep(l.Config.WaitForNewBlock)
				// fmt.Println("We've processed all the blocks, waiting for new blocks")

			}
		}
	}

	l.Config.Wg.Done()
	l.Config.Quitch <- struct{}{}
}

func (l *EvmListener) Stop() bool {
	l.Config.Quitch <- struct{}{}
	<-l.Config.Quitch
	return true
}

func ProcessBlock(notification internal_notification.InternalNotification, addressStore db.AddressStore, transactionStore db.TransactionStore, b types.Block) error {
	transactions := b.Transactions
	fmt.Println("Processing Block", b.BlockNumber)

	for _, t := range transactions {
		err := ProcessTransaction(notification, addressStore, transactionStore, t)
		if err != nil {
			return err
		}
	}

	return nil
}

func ProcessTransaction(notification internal_notification.InternalNotification, addressStore db.AddressStore, transactionStore db.TransactionStore, tx types.Transaction) error {
	if tx.To == "" {
		return nil
	}

	txType, err := GetTransactionType(addressStore, tx)
	if err != nil {
		return err
	}

	if txType != "" {
		tx.TxType = txType
		n := big.Int{}
		ethValue, ok := n.SetString(tx.Value, 10)
		if !ok {
			return errors.New("could not convert wei to ETH")
		}
		tx.Value = utils.WeiToEther(ethValue).String()
		err := transactionStore.SaveTransaction(&tx)
		if err != nil {
			return err
		}

		// NOT IMPLEMENTED
		err = notification.Notify("TRANSACTION_DETECTED", fmt.Sprintf("Received Transaction of type %s from %s to %s with the value of %s Ether\n", txType, tx.From, tx.To, tx.Value))
		if err != nil {
			panic("NOT IMPLEMENTED")
		}

	}

	return nil
}

func GetTransactionType(addressStore db.AddressStore, tx types.Transaction) (string, error) {
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

func IsDepositTransaction(addressStore db.AddressStore, from string, to string) (bool, error) {
	exists, err := addressStore.AddressExists(strings.ToLower(to))
	if err != nil {
		return false, err
	}
	return exists, nil
}

func IsWithdrawalTransaction(addressStore db.AddressStore, from string, to string) (bool, error) {
	exists, err := addressStore.AddressExists(strings.ToLower(from))
	if err != nil {
		return false, err
	}
	return exists, nil
}
