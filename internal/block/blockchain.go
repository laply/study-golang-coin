package block

import (
	"fmt"
	"sync"
)

var b *blockchain
var once sync.Once

type blockchain struct {
	blocks []Block
}

func GetBlockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{
				[]Block{*initBlock("Genesis Block")},
			}
		})

		fmt.Println("Genesis Block Created")
	}

	return b
}

func (bc *blockchain) ListBlocks() string {
	blcokInfo := ""
	for _, info := range bc.blocks {
		blcokInfo += fmt.Sprintf("[%s] %s / %s \n\n", info.Data, info.Hash, info.PrevHash)
	}
	return blcokInfo
}

func (bc *blockchain) GetBlocks() []Block {
	return bc.blocks
}

func (bc *blockchain) AddBlock(data string) {
	if len(bc.blocks) == 0 {
		bc.blocks = append(bc.blocks, *initBlock(data))
		return
	}

	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := prevBlock.newBlock(data)
	bc.blocks = append(bc.blocks, *newBlock)
}
