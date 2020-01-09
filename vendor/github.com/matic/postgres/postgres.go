package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
	"github.com/matic/ether"
	"github.com/matic/util"
)

// GetTransactions -
func GetTransactions(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	accountID := vars["AccountId"]
	fmt.Fprintln(w, "AccountId - ", accountID)

	pgClient, err := util.GetPostgresClient()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := pgClient.Query("SELECT txHash, blockNo, fromAdr, toAdr FROM transaction WHERE fromAdr=$1", accountID)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	transactionsList := []util.Transaction{}

	for rows.Next() {

		transaction := util.Transaction{}

		err = rows.Scan(&transaction.TransactionHash, &transaction.BlockNumber, &transaction.From, &transaction.To)
		if err != nil {
			panic(err)
		}
		transactionsList = append(transactionsList, transaction)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	resp := util.Response{}
	resp.Transactions = transactionsList
	resp.Count = len(transactionsList)

	json.NewEncoder(w).Encode(resp)
}

func CacheBlocks(pgClient *sql.DB, ethClient *ethclient.Client) {

	blockHeight, err := ether.GetBlocksHeight(ethClient)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 10000; i++ {

		block, err := ether.GetBlockDetails(ethClient, blockHeight)
		if err != nil {
			log.Fatal(err)
		}

		blockHeight.Sub(blockHeight, big.NewInt(1))

		transactionsList := ether.GetTransactionsDetails(ethClient, *block)

		saveTransactions(pgClient, transactionsList)
	}
}

func SubscribeAndSaveBlocks(pgClient *sql.DB, ethClient *ethclient.Client) {

	headers := make(chan *types.Header)

	sub, err := ethClient.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	var transactionsList = []util.Transaction{}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Println(header.Hash().Hex()) // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f

			block, err := ethClient.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(block.Hash().Hex())        // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
			fmt.Println(block.Number().Uint64())   // 3477413
			fmt.Println(block.Nonce())             // 130524141876765836
			fmt.Println(len(block.Transactions())) // 7

			transactionsList = ether.GetTransactionsDetails(ethClient, *block)

			fmt.Println(transactionsList)
			saveTransactions(pgClient, transactionsList)
		}
	}
}

func saveTransactions(pgClient *sql.DB, transactionsList []util.Transaction) {

	for _, transaction := range transactionsList {
		// fmt.Println(transaction)

		sqlStatement := `INSERT INTO transaction (txHash, blockNo, fromAdr, toAdr) VALUES ($1, $2, $3, $4)`
		_, err := pgClient.Exec(sqlStatement, transaction.TransactionHash, transaction.BlockNumber, transaction.From, transaction.To)
		if err != nil {
			panic(err)
		}
	}
}
