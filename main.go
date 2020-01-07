package main

import (
	"net/http"

	"github.com/matic/postgres"
	"github.com/matic/util"
)

func main() {
	http.HandleFunc("/transactions/{AccountId}", postgres.GetTransactions)
	http.ListenAndServe(":8000", nil)

	ethClient, _ := util.GetEthClient()
	pgClient, _ := util.GetPostgresClient()

	postgres.CacheBlocks(pgClient, ethClient)
	postgres.SubscribeAndSaveBlocks(pgClient, ethClient)
}
