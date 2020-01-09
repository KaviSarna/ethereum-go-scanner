package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matic/postgres"
	"github.com/matic/util"
)

func main() {

	ethClient, _ := util.GetEthClient()
	pgClient, _ := util.GetPostgresClient()

	postgres.CacheBlocks(pgClient, ethClient)
	postgres.SubscribeAndSaveBlocks(pgClient, ethClient)
	
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from server")
	})
	router.HandleFunc("/transactions/{AccountId}", postgres.GetTransactions).Methods("GET")
	log.Fatal(http.ListenAndServe(":8010", router))
}
