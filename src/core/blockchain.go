package core

// 定义区块链结构
type Blockchain struct {
	Blocks []*Block
}

/**
 * 添加区块方法
 * @param {string} data 区块数据
 */
func (bc *Blockchain) AppendBlock(data string) {
	// 获取上一个区块的数据
	prevBlock := bc.Blocks[len(bc.Blocks)-1]

	// 生成新的区块
	newBlock := NewBlock(data, prevBlock.Hash)

	// 生成新的区块链
	bc.Blocks = append(bc.Blocks, newBlock)
}
