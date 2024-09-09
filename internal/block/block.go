package block

import (
	"crypto/sha256"
	"fmt"
)

type block struct {
	data     string
	hash     string
	prevHash string
}

func (b *block) calculateHash() {
	b.hash = fmt.Sprintf("%x", sha256.Sum256([]byte(b.data+b.prevHash)))
}

func (b *block) newBlock(data string) *block {
	newBlock := &block{data: data, prevHash: b.hash}
	newBlock.calculateHash()
	return newBlock
}

func initBlock(data string) *block {
	b := &block{data: data, prevHash: ""}
	b.calculateHash()
	return b
}
