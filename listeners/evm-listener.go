package listeners

import (
	"fmt"
	"open-payment-gateway/db"
	"open-payment-gateway/providers"
	"open-payment-gateway/types"
)

type EvmListener struct {
	network             types.Network
	currency            types.Currency
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
			fmt.Printf("Processing Block is %+v\n", processingBlock)

			if err != nil {
				panic("Could not get the block data from provider")
			}

			if err := ProcessBlock(processingBlock); err != nil {
				panic("Could not process block")
			}

			if err := l.db.SaveBlock(&processingBlock); err != nil {
				panic("Could not save the processed block into the database")
			}

		} else {
			//  Wait for new blocks to be mined
			// fmt.Println("We've processed all the blocks, waiting for new blocks")

		}
	}
}

func NewEvmListener(n types.Network, c types.Currency, cid int64, sbn int64, db db.DB, p providers.EvmProvider) *EvmListener {
	return &EvmListener{network: n, currency: c, chainId: cid, startingBlockNumber: sbn, db: db, provider: p}
}

func ProcessBlock(b types.Block) error {
	return nil
}
