package block

import (
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Data     string
	Hash     string
	PrevHash string
}

func (b *Block) calculateHash() {
	b.Hash = fmt.Sprintf("%x", sha256.Sum256([]byte(b.Data+b.PrevHash)))
}

func (b *Block) newBlock(data string) *Block {
	newBlock := &Block{Data: data, PrevHash: b.Hash}
	newBlock.calculateHash()
	return newBlock
}

func initBlock(data string) *Block {
	b := &Block{Data: data, PrevHash: ""}
	b.calculateHash()
	return b
}
