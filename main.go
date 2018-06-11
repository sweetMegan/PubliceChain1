package main

import "zhqGo/PublicChain/BLC"

func main() {
	//创建区块链
	blc := BLC.CreateGenesisBlockchain()
	//添加新区块
	blc.Addblock("first block")
	//添加新区块
	blc.Addblock("second block2")
	//打印区块链
	blc.PrintBlockchain()

}
