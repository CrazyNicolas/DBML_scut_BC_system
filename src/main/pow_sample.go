package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math/big"
	"time"
)

// 设置困难度为2
const target_bit = 16

// 区块对象
type Block struct {
	Height    int64
	Data      []byte
	Nonce     int64
	Timestamp int64
	Hash      []byte
	Prev_hash []byte
}

// 工作量证明对象
type Pow struct {
	Cur_block *Block
	target    *big.Int
}

func (p *Pow) PrepareData(nonce int64) []byte {
	data := bytes.Join(
		[][]byte{
			p.Cur_block.Prev_hash,
			p.Cur_block.Data,
			Int64ToBytes(nonce),
			Int64ToBytes(p.Cur_block.Timestamp),
			//这个其实就是困难度
			Int64ToBytes(int64(target_bit)),
			Int64ToBytes(p.Cur_block.Height),
		}, []byte{})
	return data
}

func (p *Pow) IsValid() bool {
	var hashInt *big.Int
	hashInt.SetBytes(p.Cur_block.Hash)

	if p.target.Cmp(hashInt) == 1 {
		return true
	}
	return false
}

func (p *Pow) Run() ([]byte, int64) {
	// 应该是一个循环 不断迭代nonce
	nonce := 0
	// 存储新生成的哈希用的两个版本 一个大整数 一个字节
	var hashInt *big.Int = big.NewInt(0)
	var hash [32]byte
	for {
		// 这就进入了循环之中，并且开始不断的求hash使得能够跳出循环
		data := p.PrepareData(int64(nonce))
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		// 下面我们要进行比较 看看是否有可能达到目标
		if p.target.Cmp(hashInt) == 1 {
			break
		}
		nonce += 1
	}
	return hash[0:0], int64(nonce)

}

func NewPow(block *Block) *Pow {
	// 首先初始化一个target值
	hashInt := big.NewInt(1)

	// 然后把他移位到16位难度水平
	hashInt = hashInt.Lsh(hashInt, 256-target_bit)

	return &Pow{
		block,
		hashInt,
	}
}

func New_Block(prev_hash []byte, data string, height int64) *Block {
	block := &Block{
		height,
		[]byte(data),
		0,
		time.Now().Unix(),
		nil,
		prev_hash,
	}

	pow := NewPow(block)

	hash, nonce := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

func GenisBlock() *Block {
	return New_Block([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, "Hello world!", 1)
}

//相应的应该有一个区块链的结构来承载新建的区块链
type BlockChain struct {
	Blocklist []*Block
}

func (bc *BlockChain) AddNewBlock(prev_hash []byte, data string, height int64) {
	block := New_Block(prev_hash, data, height)
	bc.Blocklist = append(bc.Blocklist, block)
}

func CreateNewBlockChain() *BlockChain {
	block := GenisBlock()
	return &BlockChain{[]*Block{block}}
}

func Int64ToBytes(i int64) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, i)
	if err != nil {
		fmt.Println("Transfer Err: ", err)
	}
	return buf.Bytes()
}

func main() {
	blockchain := CreateNewBlockChain()

	blockchain.AddNewBlock(blockchain.Blocklist[len(blockchain.Blocklist)-1].Hash, "send 10 rmb to Jack", blockchain.Blocklist[len(blockchain.Blocklist)-1].Height+1)
	blockchain.AddNewBlock(blockchain.Blocklist[len(blockchain.Blocklist)-1].Hash, "send 200 rmb to Lily", blockchain.Blocklist[len(blockchain.Blocklist)-1].Height+1)
	blockchain.AddNewBlock(blockchain.Blocklist[len(blockchain.Blocklist)-1].Hash, "send 80 rmb to Mzy", blockchain.Blocklist[len(blockchain.Blocklist)-1].Height+1)
	blockchain.AddNewBlock(blockchain.Blocklist[len(blockchain.Blocklist)-1].Hash, "send 100 rmb to Mommy", blockchain.Blocklist[len(blockchain.Blocklist)-1].Height+1)

	fmt.Println(blockchain)
	fmt.Println(blockchain.Blocklist)
}
