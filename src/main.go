package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func calcHash(data string) string {
	hashByte := sha256.Sum256([]byte(data))
	hashData := hex.EncodeToString(hashByte[:])
	return hashData
}

func main() {
	fmt.Printf("%s, %s", "test", calcHash("test"))
}
