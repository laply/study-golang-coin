package block

import (
	"fmt"
	"sync"
)

var b *blockchain
var once sync.Once

type blockchain struct {
	blocks []block
}

func GetBlockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{
				[]block{*initBlock("Genesis Block")},
			}
		})

		fmt.Println("Genesis Block Created")
	}

	return b
}

func (bc *blockchain) ListBlocks() {
	for _, info := range bc.blocks {
		fmt.Printf("[%s] %s / %s \n", info.data, info.hash, info.prevHash)
	}
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
