package blockchain

import (
	"fmt"
	"sync"

	"github.com/byungsujeong/gocoin/db"
	"github.com/byungsujeong/gocoin/utils"
)

type blockchain struct {
	NewestHash string `json:"newestHash"`
	Height     int    `json:"height"`
}

var b *blockchain
var once sync.Once

func (b *blockchain) restore(data []byte) {
	utils.FromBytes(b, data)
}

func (b *blockchain) persit() {
	db.SaveBlockchain(utils.ToBytes(b))
}

func (b *blockchain) AddBlock(data string) {
	block := createBlock(data, b.NewestHash, b.Height+1)
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.persit()
}

func (b *blockchain) Blocks() []*Block {
	var blocks []*Block
	hashCursor := b.NewestHash
	for {
		block, _ := FindBlock(hashCursor)
		blocks = append(blocks, block)
		if block.PrevHash != "" {
			hashCursor = block.PrevHash
		} else {
			break
		}
	}
	return blocks
}

func Blockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{"", 0}
			// search for checkpoint on the db
			checkpoint := db.Checkpoint()
			if checkpoint == nil {
				b.AddBlock("Genesis Block")
			} else {
				// restore b from bytes
				b.restore(checkpoint)
			}
		})
	}
	fmt.Println(b.NewestHash)
	return b
}

// func (b *Block) calculateHash() {
// 	hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
// 	b.Hash = fmt.Sprintf("%x", hash)
// }

// func getLastHash() string {
// 	totalBlocks := len(GetBlockchain().blocks)
// 	if totalBlocks == 0 {
// 		return ""
// 	}
// 	return GetBlockchain().blocks[totalBlocks-1].Hash
// }

// func createBlock(data string) *Block {
// 	newBlock := Block{data, "", getLastHash(), len(GetBlockchain().blocks) + 1}
// 	newBlock.calculateHash()
// 	return &newBlock
// }

// func (b *blockchain) AddBlock(data string) {
// 	b.blocks = append(b.blocks, createBlock(data))
// }

// func GetBlockchain() *blockchain {
// 	if b == nil {
// 		once.Do(func() {
// 			b = &blockchain{}
// 			b.AddBlock("Genesis Block")
// 		})
// 	}
// 	return b
// }

// func (b *blockchain) AllBlocks() []*Block {
// 	return b.blocks
// }

// var ErrNotFound = errors.New("block not found")

// func (b *blockchain) GetBlock(height int) (*Block, error) {
// 	if height > len(b.blocks) {
// 		return nil, ErrNotFound
// 	}
// 	return b.blocks[height-1], nil
// }
