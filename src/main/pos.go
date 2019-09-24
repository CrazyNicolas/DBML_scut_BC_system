package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math/big"
	"time"
)

//定义区块  以下数据域的数量是对的 只不过 数据域的类型与真实的类型相差甚远
type Block_POS struct {
	Height    int64
	TimeStamp string
	// 这个区块是有谁验证的
	Validator string
	// 当前区块哈希
	Hash string
	// 父区块的哈希
	Prev_Hash string
	//交易数据
	Data int
}

// 简化一下目标 先把区块链近似看作是一个数组
var BlockChain_POS []Block

// 先写一个创建区块的办法 这是在选定由谁来记账以后进行的操作
func GenerateBlock(data int, oldblock Block_POS, address string) (Block_POS, error) {
	// 首先构建一个区块
	var newBlock Block_POS

	newBlock.Data = data
	newBlock.Height = oldblock.Height + 1
	newBlock.TimeStamp = time.Now().String()
	newBlock.Prev_Hash = oldblock.Hash
	newBlock.Validator = address

	// 上面的数据与都可以直接获得 但是当前我们要构建的区块的hash不是那么好获得需要计算,调用CaiculateBlock函数
	newBlock.Hash = CalculateBlock(newBlock)
	return newBlock, nil
}

// 下面来写这个计算函数 很多帖子将计算函数拆成两部分 个人认为没有必要
func CalculateBlock(block Block_POS) string {
	// 应该先把这个block的字段先拼接起来

	//接下来我们应该将这个拼接起来的字符串进行哈希 返回结果
}
