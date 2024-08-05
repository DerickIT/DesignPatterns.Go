package easyblockchain

type Blockchain struct {
	blocks []*Block
}

func (bc *Blockchain) GetBlocks() []*Block {
	return bc.blocks
}

func (bc *Blockchain) AddBlock(data string) {
	pewvBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, pewvBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)

}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
