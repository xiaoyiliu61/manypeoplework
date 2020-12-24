package main

import (
	_ "ManyPeopleWork/routers"
	"fmt"
	"github.com/astaxie/beego"
	"manypeoplework/utils"
)

func main() {

	Getdifficulty1:=utils.GetDifficulty();
	fmt.Println("当前区块的难度值是",Getdifficulty1)
	beego.Run()
	fmt.Println("hello world")
}

