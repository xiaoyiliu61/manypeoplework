package main

import (
	"BcRPCCode04/entity"
	"BcRPCCode04/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func hanshu1(Method string,params1 ...interface{})([]byte) {
	rpcReq :=entity.RPCRequest{}
	rpcReq.Id = time.Now().Unix()
	rpcReq.Jsonrpc = "2.0"
	rpcReq.Method = Method      // utils.GETBLOCKCOUNT//获取当前节点的数量
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
	request,err  :=http.NewRequest("POST",utils.RPCURL,bytes.NewBuffer(rpcBytes))//NewBuffer,建立一个缓存区，进行读操作和写操作
	if err != nil {                                                                          //!!!!!!!!!!!!!!!!!!!!!!
		fmt.Println(err.Error())
		return err.Error()
	}
	//给post请求添加请求头
	//key --> value
	request.Header.Add("Encoding","UTF-8")
	request.Header.Add("Content-Type","application")
	request.Header.Add("Authorization","Basic "+utils.Base64Str(utils.RPCUSER+":"+utils.RPCPASSWORD))

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
