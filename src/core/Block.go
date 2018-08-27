package core

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// 定义区块结构
type Block struct {
	Timestamp    int64  // 区块生成的时间戳
	Data         []byte // 区块存储的数据
	PrvBlockHash []byte // 前一个区块的 Hash 值
	Hash         []byte // 区块自身 Hash 值，用于校验区块数据有效
}

// 生成 Hash 值方法
func (block *Block) SetHash() {
	// 时间戳转为 字节数组
	timestamp := []byte(strconv.FormatInt(block.Timestamp, 10))

	// 拼接所有字节
	headers := bytes.Join([][]byte{block.PrvBlockHash, block.Data, timestamp}, []byte{})

	// 使用 sha256 转为新的 Hash 值
	block.Hash = sha256.Sum256(headers)[:]
}

/**
 * 创建新的区块
 * @param {string} data 区块数据
 * @param {[]byte} prvBlockHash 上一个区块的 Hash 值
 */
func NewBlock(data string, prvBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prvBlockHash, []byte{}}
	block.SetHash()
	return block
}
