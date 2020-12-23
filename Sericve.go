package main

import (
	"BcRPCCode04/entity"
	"BcRPCCode04/utils"
	"encoding/json"
	"fmt"

	//"encoding/json"
)

func GetDifficulty()interface{}{
	rpcBytes:=hanshu1(utils.GETDIFFICULTY)
	return hanshu2(rpcBytes)
}

func GetBlockCount()interface{}{
	rpcBytes:=hanshu1(utils.GETBLOCKCOUNT)
	return hanshu2(rpcBytes)
}

func GetBestBlockHash()interface{}{
	rpcBytes:=hanshu1(utils.GETBESTBLOCKHASH)
	return hanshu2(rpcBytes)
}

func GetBlockChainInfo()entity.RPCGetBlockChainInfo{
	rpcBytes:=hanshu1(utils.GETBLOCKCHAININFO)
	result :=hanshu2(rpcBytes)
    //fmt.Println(result)
	result2,err := json.Marshal(result)
	//fmt.Println(string(result2) )
	  result1 := entity.RPCGetBlockChainInfo{}
	if err != nil {
      fmt.Println("报错1")
      panic(err.Error())
	}
    err = json.Unmarshal(result2,&result1)

	if err != nil {
		fmt.Println("报错2")
		panic(err.Error())
	}
	//fmt.Println(result1)
	return  result1
}

func GetBlockHashByHeight(height int)interface{}{
	rpcBytes:=hanshu1(utils.GETBLOCKHASH,height)
	return hanshu2(rpcBytes)
}

func GetNewAddress(label string, ADDRESS_TYPE string)interface{}{
	switch (ADDRESS_TYPE) {
	case utils.LEGACY:
		ADDRESS_TYPE = "legacy";
		break;
	case utils.P2SH_SEGWIT:
		ADDRESS_TYPE = "p2sh-segwit";
		break;
	case utils.BECH32:
		ADDRESS_TYPE = "bech32";
		break;
	default:
		fmt.Println("请输入LEGACY，P2SH_SEGWIT，BECH32中的一个")
		break
		/*LEGACY = "LEGACY"
		P2SH_SEGWIT = "P2SH_SEGWIT"
		BECH32 ="BECH322"*/
	}
	rpcBytes:=hanshu1(utils.GETNEWADDRESS,label,ADDRESS_TYPE )
	return hanshu2(rpcBytes)
}

