// 变量声明与初始化
package main

import "fmt"

func main() {
	var v1 int = 10 // 正确的使用方式1
	var v2 = 10     // 正确的使用方式2，编译器可以自动推导出v2的类型
	v3 := 10        // 正确的使用方式3，编译器可以自动推导出v3的类型 （:= 表示同时进行变量声明与初始化的工作）

	fmt.Println("v1 = ", v1, "v2 = ", v2, "v3 = ", v3)
}
