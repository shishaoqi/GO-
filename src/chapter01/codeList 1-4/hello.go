package main

// 每个Go源代码文件的开头都是一个package声明，表示该Go代码所属的包。包是Go语言里最基本的分发单位，也是工程管理中依赖关系的体现。

// 要生成Go可以执行程序，必须建立一个包名为 main 的包，并且在该包中包含一个叫 main() 的函数（该函数是Go可执行程序的执行起点）

import "fmt"

func main() {
	fmt.Println("Hello world")
}
