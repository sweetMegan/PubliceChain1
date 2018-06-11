package BLC

type Blockchain struct {
	//链上所有数据应该是保存到本地数据库中，暂时先用切片简单实现
	blocks []*Block
}

func  CreateGenesisBlockchain()(*Blockchain) {
	//创建创世区块
	block := CreateGenesisBlock()
	//创建区块链
	blc := new(Blockchain)
	blc.blocks = append(blc.blocks,block)
	return blc
}
func (blc *Blockchain) Addblock(data string)  {
	//获取链上最后一个区块的hash值
	lastBlock := blc.blocks[len(blc.blocks)-1]
	//将最后一个区块的hash值，传给新区块，做为prehash
	block := CreateBlock(data,lastBlock.Hash)
	blc.blocks = append(blc.blocks,block)
}
func (blc *Blockchain) PrintBlockchain() {
	for _, block := range blc.blocks {
		block.PrintBlock()
	}
}