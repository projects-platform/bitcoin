package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"./core"
)

var blockChain *core.Blockchain

func main() {
	blockChain = core.NewBlockchain()

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		data, err := json.Marshal(blockChain)
		if err != nil {
			log.Panic("Error")
		}
		io.WriteString(w, string(data))
	})

	http.HandleFunc("/write", func(w http.ResponseWriter, req *http.Request) {
		blockChain.AppendBlock(req.FormValue("data"))
		io.WriteString(w, "区块写入完成")
	})

	http.ListenAndServe(":3000", nil)
}
