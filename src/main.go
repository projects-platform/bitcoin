package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"./core"
)

var blockchain *core.Blockchain

func run() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/write", WriteHandler)
	http.ListenAndServe(":3000", nil)
}

func IndexHandler(w http.ResponseWriter, req *http.Request) {
	v, err := json.Marshal(blockchain)
	if err != nil {
		fmt.Printf(err.Error())
	}

	io.WriteString(w, string(v))
}

func WriteHandler(w http.ResponseWriter, req *http.Request) {
	blockchain.AppendBlock(req.FormValue("data"))
	io.WriteString(w, "完成新区块的写入")
}

func main() {
	blockchain = core.NewBlockchain()
	run()
}
