package main

import (
	"calc/simplemath"
	"fmt"
	"os"
	"strconv"
)

/*
// 基于命令行的计算机程序
<calcporj>
|--<src>
	|--<calc>
		|--calc.go
	|--<simplemath>
		|--add.go
		|--add_test.go
		|--sqrt.go
		|--sqrt_test.go
|--<bin>
|--<pkg>#包将被安装到此处
*/

var Usage = func() {
	fmt.Println("USAGE: calc command [arguments] ...")
	fmt.Println("\nThe commands are:\n\tadd\tAddition of two values.\n\tsqrt\tsquare root of a non-negative value")
}

func main() {
	args := os.Args
	fmt.Println(args)
	fmt.Println(len(args))
	if args == nil || len(args) < 2 {
		Usage()
		return
	}

	switch args[1] {
	case "add":
		if len(args) != 4 {
			addErrMsg()
			return
		}

		v1, err1 := strconv.Atoi(args[2])
		v2, err2 := strconv.Atoi(args[3])
		if err1 != nil || err2 != nil {
			addErrMsg()
			return
		}
		ret := simplemath.Add(v1, v2)
		fmt.Println("Result: ", ret)
	case "sqrt":
		if len(args) != 3 {
			sqrtErrmsg()
			return
		}
		v, err := strconv.Atoi(args[2])
		if err != nil {
			sqrtErrmsg()
			return
		}
		ret := simplemath.Sqrt(v)
		fmt.Println("Result: ", ret)
	default:
		Usage()
	}
}

func addErrMsg() {
	fmt.Println("USAGE: calc add <inteter1> <interger2>")
}

func sqrtErrmsg() {
	fmt.Println("USAGE: calc sqrt <interger>")
}
