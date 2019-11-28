package main

import (
	"log"

	"github.com/binance-chain/go-sdk/client/rpc"
	"github.com/binance-chain/go-sdk/common/types"
	"github.com/davecgh/go-spew/spew"
	"github.com/openlyinc/pointy"
)

func main() {
	// nodeAddr := "http://dataseed1.ninicoin.io" // Public
	nodeAddr := "http://45.76.119.188:27147" // Kai
	// nodeAddr := "127.0.0.1:27147"
	client := rpc.NewRPCClient(nodeAddr, types.ProdNetwork)
	if err := client.Start(); err != nil {
		log.Println(err)
	}

	// status, err := client.Status()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// spew.Dump(status)

	block, err := client.Block(pointy.Int64(51428800))
	if err != nil {
		log.Fatal(err)
	}
	spew.Dump(block)
}
