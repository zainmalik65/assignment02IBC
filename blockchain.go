package assignment02IBC

import (
	"crypto/sha256"
	"fmt"
	
)

func CalculateBalance(userName string, chainHead *Block) int {
	var balance int=0
	tempBlock:=chainHead
	for tempBlock!=nil {
		if value, found := tempBlock.Spender[userName]; found {
			balance = balance - value
		} else {

		}
		if value2, found2 := tempBlock.Receiver[userName]; found2 {
			balance = balance + value2
		} else {

		}
		tempBlock=tempBlock.PrevPointer
	}
	return  balance

}

func CalculateHash(inputBlock *Block) string {
	hashes := sha256.New()
	hashes.Write([]byte(fmt.Sprintf("%v", inputBlock.Spender)))
	hashes.Write([]byte(fmt.Sprintf("%v", inputBlock.Receiver)))
	hashes.Write([]byte(fmt.Sprintf("%v", inputBlock.PrevHash)))
	return fmt.Sprintf("%x", hashes.Sum(nil))
}


func InsertBlock(spendingUser string, receivingUser string, miner string, amount int, chainHead *Block) *Block {
	var newBlock Block=Block{nil,nil ,nil,"nil","nil"}

	tempBlock:=chainHead


	if miner == "Satoshi" {


		if tempBlock==nil {
			spenderBalance := CalculateBalance(spendingUser, chainHead)
			if (spenderBalance >= amount) {
				newBlock.Spender = make(map[string]int)
				newBlock.Receiver = make(map[string]int)
				newBlock.Receiver[miner] = 100
				newBlock.PrevPointer = tempBlock
				newBlock.PrevHash = "nil"
				newBlock.CurrentHash = CalculateHash(&newBlock)
				fmt.Println(miner, " successfully mines a block")
			} else {
				fmt.Println(spendingUser, " has insufficient balance for this transaction")
				return chainHead
			}

		} else {
			if spendingUser!="" {
				spenderBalance := CalculateBalance(spendingUser, chainHead)

				if (spenderBalance > amount) {
					newBlock.Spender = make(map[string]int)
					newBlock.Receiver = make(map[string]int)
					newBlock.Spender[spendingUser] = amount
					newBlock.Receiver[receivingUser] = amount
					newBlock.Receiver[miner] = 100
					newBlock.PrevPointer = tempBlock
					newBlock.PrevHash = tempBlock.CurrentHash
					newBlock.CurrentHash = CalculateHash(&newBlock)
					fmt.Println(miner, " successfully mines a block")
				} else {
					fmt.Println(spendingUser, " has insufficient balance for this transaction")
					return chainHead
				}
			}else
			{
				newBlock.Spender=make(map[string]int)
				newBlock.Receiver=make(map[string]int)
				newBlock.Receiver[miner]=100
				newBlock.PrevPointer=chainHead
				newBlock.PrevHash=chainHead.CurrentHash
				newBlock.CurrentHash=CalculateHash(&newBlock)
				fmt.Println(miner," successfully mines a block")
			}


		}


	} else{
		fmt.Println(miner,"has not permission to mine block")
		return chainHead
	}



	return  &newBlock
}




func ListBlocks(chainHead *Block) {
	genesisBlock :=chainHead

	fmt.Println("Current blocks in the BlockChain are:")
	count:=1
	for genesisBlock!=nil{
		fmt.Println("Block ",count,":")

		fmt.Println("Spender Map:",genesisBlock.Spender)
		fmt.Println("Reciver Map:",genesisBlock.Receiver)

		genesisBlock=genesisBlock.PrevPointer
		count=count+1
	}
}

func VerifyChain(chainHead *Block) {
	blockChainHead:=chainHead
	prevBlock:=blockChainHead.PrevPointer
	checkInt:=true
	for prevBlock!=nil{

		if blockChainHead.PrevHash!=prevBlock.CurrentHash {
			checkInt=false

			break
		}
		blockChainHead=prevBlock
		prevBlock=prevBlock.PrevPointer


	}
	if checkInt==true {
		fmt.Println("Blockchain is intact")
	}else{
		fmt.Println("Blockchain is modified")
	}
}

