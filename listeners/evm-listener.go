package listeners

import (
	"fmt"
	"open-payment-gateway/db"
	"open-payment-gateway/providers"
	"open-payment-gateway/types"
	"strings"
)

type EvmListener struct {
	network             string
	currency            string
	chainId             int64
	startingBlockNumber int64
	addressStore        db.AddressStore
	blockStore          db.BlockStore
	provider            providers.EvmProvider
}

func (l EvmListener) Start() {

	for {
		latestProcessedBlockNumber, err := l.blockStore.GetLatestProcessedBlockNumber()
		if err != nil {
			panic("Could not get the latest processed block number from database")
		}

		// Check if we need to skip some blocks if the starting block number is not equal to -1
		if l.startingBlockNumber > -1 && l.startingBlockNumber > latestProcessedBlockNumber {
			latestProcessedBlockNumber = l.startingBlockNumber
		}

		fmt.Printf("latest Processes Block Number is %d\n", latestProcessedBlockNumber)

		latestBlockNumber, err := l.provider.GetLatestBlockNumber()
		if err != nil {
			panic("Could not get the latest block number from provider")
		}

		fmt.Printf("latest Block Number is %d\n", latestBlockNumber)

		if latestBlockNumber > latestProcessedBlockNumber {
			// Iterate through blocks and process them
			// fmt.Println("We're behind")

			processingBlockNumber := latestProcessedBlockNumber + 1
			fmt.Printf("Processing Block Number is %d\n", processingBlockNumber)
			processingBlock, err := l.provider.GetBlockByNumber(processingBlockNumber)

			if err != nil {
				panic("Could not get block data from provider")
			}

			if err := ProcessTransactions(l.addressStore, processingBlock); err != nil {
				panic("Could not process block")
			}

			if err := l.blockStore.SaveBlock(&processingBlock); err != nil {
				panic("Could not save processed block into the database")
			}

		} else {
			//  Wait for new blocks to be mined
			// fmt.Println("We've processed all the blocks, waiting for new blocks")

		}
	}
}

func NewEvmListener(network string,
	currency string,
	chainId int64,
	startingBlockNumber int64,
	addressStore db.AddressStore,
	blockStore db.BlockStore,
	provider providers.EvmProvider,
) *EvmListener {
	return &EvmListener{
		network:             network,
		currency:            currency,
		chainId:             chainId,
		startingBlockNumber: startingBlockNumber,
		addressStore:        addressStore,
		blockStore:          blockStore,
		provider:            provider,
	}
}

func ProcessTransactions(addressStore db.AddressStore, b types.Block) error {
	transactions := b.Transactions
	fmt.Println("Processing Block", b.BlockNumber)

	for _, t := range transactions {
		err := ProcessTransaction(addressStore, t)
		if err != nil {
			return err
		}
	}

	return nil
}

func ProcessTransaction(addressStore db.AddressStore, tx types.Transaction) error {
	if tx.To == "" {
		return nil
	}

	txType, err := GetTransactionType(addressStore, tx)
	if err != nil {
		return err
	}

	if txType != "" {
		fmt.Printf("Received Transaction of type %s from %s to %s with the value of %s Ether\n", txType, tx.From, tx.To, tx.Value)
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
