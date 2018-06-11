package BLC

import (
	"time"
	"crypto/sha256"
	"bytes"
	"strconv"
	"fmt"
)

type Block struct {
	//时间戳
	Timestame int64
	//交易数据
	Data []byte
	//上一个区块hash
	Prehash []byte
	//区块hash
	Hash []byte
	//
	Nonce int64
}
//
//func (block *Block) SetHash() {
//	//时间戳转byte数组
//	timestameBytes := []byte(strconv.FormatInt(block.Timestame,10))
//	//将区块转byte数组
//	blockData := bytes.Join([][]byte{
//		timestameBytes,
//		block.Data,
//		block.Prehash},[]byte{})
//	//计算hash值
//	hash := sha256.Sum256(blockData)
//	block.Hash = hash[:]
//}
func  CreateBlock(data string,preHash []byte)(*Block) {
	timestame := time.Now().Unix()
	block := &Block{Timestame:timestame,Data:[]byte(data),Prehash:preHash}
	//挖矿
	PandingHash(block)
	return block
}
func (block *Block)PrintBlock()  {
	fmt.Println("timestamp:",block.Timestame)
	fmt.Printf("Prehash:%x\n",block.Prehash)
	fmt.Printf("Data:%s\n",block.Data)
	fmt.Printf("Hash:%x\n",block.Hash)
	fmt.Println("=====================")
}
func CreateGenesisBlock()*Block  {
	//长度为32bits的byte数组,刚好256位与sha256得到的hash值长度对应
	var preHash = [32]byte{}
	block := CreateBlock("genesis block",preHash[:])
	return block
}