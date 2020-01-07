package util

import (
	"database/sql"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

// Transaction structure
type Transaction struct {
	From            string `json:"from"`
	To              string `json:"to"`
	BlockNumber     string `json:"blockNumber"`
	TransactionHash string `json:"transactionHash"`
}

// GetEthClient -  Function to connect to a ethereum client and return it
func GetEthClient() (*ethclient.Client, error) {

	client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return client, nil
}

// GetPostgresClient -  Function to connect to a postgres client and return it
func GetPostgresClient() (*sql.DB, error) {

	psqlInfo := "host=localhost port=5050 user=test password=test dbname=eth sslmode=disable"
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return db, nil
}
