package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)
   
func main() {
	url := "https://testnet.hashio.io/api"

	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	
	/*
		Some block numbers that have this problem
		- 3454401
		- 3730571
		- 3454410
	*/
	block, err := client.BlockByNumber(context.Background(), big.NewInt(3454401))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(block)
	
}