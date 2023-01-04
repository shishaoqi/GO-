package main

import "fmt"

func main() {
	var v1 int
	var v2 string
	var v3 [10]int // 数组
	var v4 []int   // 切片
	var v5 struct {
		f int
	}
	var v6 *int           // 指针
	var v7 map[string]int // map, key 为string类型，value为int类型
	var v8 func(a int) int

	fmt.Println("v1 = ", v1, "v2 = ", v2, "v3 = ", v3, "v4 = ", v4,
		"v5 = ", v5, "v6 = ", v6, "v7 = ", v7, "v8 = ", v8)

	// var 关键字的另一种用法是可以将若干个需要声明的变量放置在一起，免得程序员需要重复
	// 写 var 关键字，如下：
	var (
		val1 int
		val2 string
	)
	fmt.Println("val1 = ", val1, "val2 = ", val2)
}
