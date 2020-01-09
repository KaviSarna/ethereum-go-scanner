package ether

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/matic/util"
)

func GetBlocksHeight(client *ethclient.Client) (*big.Int, error) {

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	number := header.Number
	fmt.Println("Latest Block Number - ", number.String())

	return number, nil
}

func GetBlockDetails(client *ethclient.Client, blockNumber *big.Int) (*types.Block, error) {

	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	fmt.Println(block.Number().Uint64())     // 5671744
	fmt.Println(block.Difficulty().Uint64()) // 3217000136609065
	fmt.Println(block.Hash().Hex())          // 0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9
	fmt.Println(len(block.Transactions()))

	return block, nil
}

func GetTransactionsDetails(client *ethclient.Client, block types.Block) []util.Transaction {

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	var transactionsList = []util.Transaction{}

	for _, tx := range block.Transactions() {
		// fmt.Println(tx.Hash().Hex())        // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
		// fmt.Println(tx.Value().String())    // 10000000000000000
		// fmt.Println(tx.Gas())               // 105000
		// fmt.Println(tx.GasPrice().Uint64()) // 102000000000
		// fmt.Println(tx.Nonce())             // 110644
		// fmt.Println(tx.Data())              // []
		fmt.Println("Transaction to - ", tx.To().Hex())          // 0x55fE59D8Ad77035154dDd0AD0388D09Dd4047A8e

		msg, err := tx.AsMessage(types.NewEIP155Signer(chainID))
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println(msg.From().Hex()) // 0x0fD081e3Bb178dc45c0cb23202069ddA57064258

		transaction := util.Transaction{}
		transaction.From = msg.From().Hex()
		transaction.To = tx.To().Hex()
		transaction.BlockNumber = strconv.Itoa(int(block.Number().Int64()))
		transaction.TransactionHash = tx.Hash().Hex()
                
		fmt.Println("Transaction - ", transaction)
		transactionsList = append(transactionsList, transaction)
	}

	return transactionsList
}
