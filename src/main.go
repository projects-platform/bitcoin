package main

import (
	"./core"
	"fmt"
)

func main() {

	bc := core.NewBlockchain()

	bc.AppendBlock("Send 1 bit")
	bc.AppendBlock("Send 2 bit")

	for _, value := range bc.Blocks {
		fmt.Printf("%x  %x  %s\n",value.Hash, value.PrevBlockHash,value.Data)
	}
}
