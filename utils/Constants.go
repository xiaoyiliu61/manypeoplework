package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"manypeoplework/entity"
	"net/http"
	"time"
)

const RPCURL = "http://127.0.0.1:8332"
const RPCUSER ="user"
const RPCPASSWORD = "pwd"


const GETBLOCKCOUNT= "getblockcount"
const GETDIFFICULTY= "getdifficulty"
const GETNEWADDRESS= "getnewaddress"
const GETBLOCK= "getblock"
const GETBLOCKHASH = "getblockhash"
const GETBLOCKCHAININFO= "getblockchaininfo"
const GETBESTBLOCKHASH = "getbestblockhash"
const GETTXOUTSETINFO= "gettxoutsetinfo"
const GETBALANCES= "getbalances"
const GETMEMPOOLINFO= "getmempoolinfo"
const GETNETWORKINFO= "getnetworkinfo"

const  (
LEGACY = "LEGACY"
P2SH_SEGWIT = "P2SH_SEGWIT"
BECH32 ="BECH322"
)



func GetDifficulty()interface{}{
	rpcBytes:=hanshu1(GETDIFFICULTY)
	return hanshu2(rpcBytes)
}

func GetBlockCount()interface{}{
	rpcBytes:=hanshu1(GETBLOCKCOUNT)
	return hanshu2(rpcBytes)
}

func GetBestBlockHash()interface{}{
	rpcBytes:=hanshu1(GETBESTBLOCKHASH)
	return hanshu2(rpcBytes)
}

func GetBlockChainInfo()entity.RPCGetBlockChainInfo{
	rpcBytes:=hanshu1(GETBLOCKCHAININFO)
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
	rpcBytes:=hanshu1(GETBLOCKHASH,height)
	return hanshu2(rpcBytes)
}

func GetNewAddress(label string, ADDRESS_TYPE string)interface{}{
	switch (ADDRESS_TYPE) {
	case LEGACY:
		ADDRESS_TYPE = "legacy";
		break;
	case P2SH_SEGWIT:
		ADDRESS_TYPE = "p2sh-segwit";
		break;
	case BECH32:
		ADDRESS_TYPE = "bech32";
		break;
	default:
		fmt.Println("请输入LEGACY，P2SH_SEGWIT，BECH32中的一个")
		break
		/*LEGACY = "LEGACY"
		P2SH_SEGWIT = "P2SH_SEGWIT"
		BECH32 ="BECH322"*/
	}
	rpcBytes:=hanshu1(GETNEWADDRESS,label,ADDRESS_TYPE )
	return hanshu2(rpcBytes)
}



func hanshu1(Method string,params1 ...interface{})([]byte) {
	rpcReq :=entity.RPCRequest{}
	rpcReq.Id = time.Now().Unix()
	rpcReq.Jsonrpc = "2.0"
	rpcReq.Method = Method      // GETBLOCKCOUNT//获取当前节点的数量
	rpcReq.Params = params1
	//对结构体类型进行序列化
	rpcBytes,err :=json.Marshal(&rpcReq)  //结构体一定要加&  深拷贝，引用类型
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return rpcBytes;
}

func hanshu2(rpcBytes []byte)interface{}{
	//2,发送一个post请求
	//client:客户端，客户端用于发起请求
	client := http.Client{}
	//实例化一个请求
	//POST一定要大写                                                                             ！！！！！！！！！！！！！
	request,err  :=http.NewRequest("POST",RPCURL,bytes.NewBuffer(rpcBytes))//NewBuffer,建立一个缓存区，进行读操作和写操作
	if err != nil {                                                                          //!!!!!!!!!!!!!!!!!!!!!!
		fmt.Println(err.Error())
		return err.Error()
	}
	//给post请求添加请求头
	//key --> value
	request.Header.Add("Encoding","UTF-8")
	request.Header.Add("Content-Type","application")
	request.Header.Add("Authorization","Basic "+Base64Str(RPCUSER+":"+RPCPASSWORD))

	//使用客户端发送请求
	response,err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}
	//通过response获取响应的数据
	code := response.StatusCode


	if code == 200 {
		//fmt.Println("成功")

		resultBytes,err:=ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err.Error())
			return err.Error()
		}
		//RPC调用的显示结果
		//{"result":0,"error":null,"id":1608017843}
		//一个json格式
		//fmt.Println(string(resultBytes))
		//json的反序列化
		var result entity.RPCResult
		err = json.Unmarshal(resultBytes,&result)
		if err != nil {
			fmt.Println(err.Error())
			return err.Error()
		}
		//fmt.Println("功能调用结果：" ,result.Result)
		return result.Result
	}else{
		fmt.Println(code)
		//501内部错误
		fmt.Println("失败")
	}
	return nil
}





