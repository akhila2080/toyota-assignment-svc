package main

import (
	"crypto-ltd-svc/api"
	_ "crypto-ltd-svc/config"
)

func main() {
	// start service
	api.NewCryptoSVC()
}
