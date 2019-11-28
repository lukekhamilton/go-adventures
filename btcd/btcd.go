package main

import (
	"log"

	"github.com/btcsuite/btcd/rpcclient"
	"github.com/davecgh/go-spew/spew"
)

// Ref https://en.bitcoin.it/wiki/API_reference_(JSON-RPC)#JSON-RPC
// docker image: https://github.com/fflo/docker-bitcoin
// or: https://github.com/ruimarinho/docker-bitcoin-core
func main() {
	// Connect to local bitcoin core RPC server using HTTP POST mode.
	connCfg := &rpcclient.ConnConfig{
		Host:         "localhost:8332",
		User:         "bitcoin",
		Pass:         "password",
		HTTPPostMode: true, // Bitcoin core only supports HTTP POST mode
		DisableTLS:   true, // Bitcoin core does not provide TLS by default
	}
	// Notice the notification parameter is nil since notifications are
	// not supported in HTTP POST mode.
	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		log.Println("err1")
		log.Fatal(err)
	}
	defer client.Shutdown()

	// Get the current block count.
	blockCount, err := client.GetBlockCount()
	if err != nil {
		log.Println("err2")
		log.Fatal(err)
	}
	log.Printf("Block count: %d", blockCount)

	hash, err := client.GetBlockHash(0)
	if err != nil {
		log.Println("err3")
		log.Fatal(err)
	}
	spew.Dump(hash)

	block, err := client.GetBlock(hash)
	if err != nil {
		log.Println("Failed to Getblock: ", err)
	}

	spew.Dump(block)
}
