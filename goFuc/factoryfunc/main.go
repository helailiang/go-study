package main

import (
	"fmt"
	"study/goFuc/factoryfunc/factory"
)

func main() {
	fmt.Println(factory.BookFactory{}.CreateProduct(factory.BookType).GetInfo())
}
