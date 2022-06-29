package api

import (
	"crypto-ltd-svc/config"
	"crypto-ltd-svc/service"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var svc service.Manager

func NewCryptoSVC() {
	// initialize service layer.
	svc = service.NewManager()

	// sync prices for ever 1 min in real-time into in-memory
	go func() {
		for {
			time.Sleep(1 * time.Minute)

			fmt.Println("syncing prices with in-memory")

			// sync price with inmemory
			svc.SyncPrices()
		}
	}()

	fmt.Printf("Started Crypto svc on port  %v\n", config.AppPort)

	// initialize routes
	routes()
}

func routes() {
	r := mux.NewRouter()

	r.HandleFunc("/v1/currency/{symbol}", GetSymbolPrice).Methods(http.MethodGet)
	r.HandleFunc("/v1/currency/all", GetAllSymbolPrices).Methods(http.MethodGet)

	http.ListenAndServe(config.AppPort, r)
}
