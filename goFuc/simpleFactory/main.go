package main

import (
	"fmt"
	"study/goFuc/simpleFactory/factory"
	"study/goFuc/simpleFactory/model"
)

//简单工厂模式
func main() {
	user := factory.CreateUser(factory.AdminUser)(1, "Aduser").(*model.AdminUser)
	fmt.Println(user)
}
