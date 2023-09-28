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
	db                  db.DB
	provider            providers.EvmProvider
}

func (l EvmListener) Start() {

	for {
		latestProcessedBlockNumber, err := l.db.GetLatestProcessedBlockNumber()
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

			if err := ProcessTransactions(l.db, processingBlock); err != nil {
				panic("Could not process block")
			}

			if err := l.db.SaveBlock(&processingBlock); err != nil {
				panic("Could not save processed block into the database")
			}

		} else {
			//  Wait for new blocks to be mined
			// fmt.Println("We've processed all the blocks, waiting for new blocks")

		}
	}
}

func NewEvmListener(n string, c string, cid int64, sbn int64, db db.DB, p providers.EvmProvider) *EvmListener {
	return &EvmListener{network: n, currency: c, chainId: cid, startingBlockNumber: sbn, db: db, provider: p}
}

func ProcessTransactions(d db.DB, b types.Block) error {
	transactions := b.Transactions
	fmt.Println("Processing Block", b.BlockNumber)

	for _, t := range transactions {
		err := ProcessTransaction(d, t)
		if err != nil {
			return err
		}
	}

	return nil
}

func ProcessTransaction(d db.DB, tx types.Transaction) error {
	if tx.To == "" {
		return nil
	}

	txType, err := GetTransactionType(d, tx)
	if err != nil {
		return err
	}

	if txType != "" {
		fmt.Printf("Received Transaction of type %s from %s to %s with the value of %s Ether\n", txType, tx.From, tx.To, tx.Value)
	}

	return nil
}

func GetTransactionType(d db.DB, tx types.Transaction) (string, error) {
	isDeposit, err := IsDepositTransaction(d, tx.From, tx.To)
	if err != nil {
		return "", err
	}
	isWithdrawal, err := IsWithdrawalTransaction(d, tx.From, tx.To)
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

func IsDepositTransaction(d db.DB, from string, to string) (bool, error) {
	exists, err := d.AddressExists(strings.ToLower(to))
	if err != nil {
		return false, err
	}
	return exists, nil
}

func IsWithdrawalTransaction(d db.DB, from string, to string) (bool, error) {
	exists, err := d.AddressExists(strings.ToLower(from))
	if err != nil {
		return false, err
	}
	return exists, nil
}
