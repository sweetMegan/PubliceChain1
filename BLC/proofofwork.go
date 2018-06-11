package BLC

import (
	"bytes"
	"strconv"
	"crypto/sha256"
	"math/big"
	"fmt"
)

const kTargetDifficult  = 16
func PandingHash(block *Block) {
	for  {
		//将区块转byte数组
		timestameBytes := []byte(strconv.FormatInt(block.Timestame,10))
		nonce := []byte(strconv.FormatInt(block.Nonce,10))
		blockData := bytes.Join([][]byte{
			timestameBytes,
			block.Data,
			block.Prehash,
			nonce},[]byte{})
		hash := sha256.Sum256(blockData)
		fmt.Printf("\r%x",hash)
		//判断hash是否有效
		if IsValidHash(hash[:]) {
			//找到有效hash，跳出循环
			block.Hash = hash[:]
			fmt.Println()
			break
		}else {
			//hash值无效，累加nonce，继续挖矿
			block.Nonce++
		}
	}
}
//判断hash是否合法
func IsValidHash(hash []byte)bool  {
	target := big.NewInt(1)
	target.Lsh(target,256-kTargetDifficult)
	var hashInt big.Int
	hashInt.SetBytes(hash)
	if target.Cmp(&hashInt)== 1 {
		return true
	}
	return false
}