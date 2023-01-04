package main

import (
	"fmt"
)

func main() {
	// 在 go 语法中， 变量初始化和变量赋值是两个不同的概念。
	// 下面为：声明一个变量之后的赋值过程
	var val int
	val = 123

	fmt.Println(val)

	// 多重赋值功能
	var i, j int
	i = 0
	j = 3
	fmt.Println(i, j)
	i, j = j, i
	fmt.Println(i, j)
}
