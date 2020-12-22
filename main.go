package main

import (
	_ "ManyPeopleWork/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
	fmt.Println("hello world")
}

