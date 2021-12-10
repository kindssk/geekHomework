package main

import (
	"fmt"
	"geekHomework/homework/exception"
	"geekHomework/initialize"
)

func main() {
	initialize.InitMysql()

	val, err := exception.FindUserById(1)
	if err != nil {
		return
	}
	fmt.Println(val.Id, val.Name, val.Age)
	val, err = exception.FindUserById(20)
	if err != nil {
		fmt.Println(fmt.Sprintf("%+v", err))
		return
	}
	fmt.Println(val.Id, val.Name, val.Age)
}
