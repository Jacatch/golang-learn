package main

import (
	"fmt"
)

//定义全局变量
var (
	globalvar01 int
	globalvar02 float64
)
var globalvar03, globalvar04 int = 1, 2
var globalvar05, globalvar06 = 123, "hello"
//globalvar07, globalvar08 := 123,0.1 不能用来定义全局变量
func main() {
	fmt.Print("hello world\n")
	fmt.Printf("%s %s\n", "hello", "world")
	fmt.Printf("%d + %d = %d\n", 1, 2, 1+2)

	var a int //单独进行定义时必须指定数据类型
	var c int
	a = 1
	b := 2 //	:=	符号表示定义并赋值
	c = a + b
	fmt.Println(c)

	var front = "hello" //定义时类型可忽略
	var back string = "world"
	fmt.Println(front + " " + back)

	/**
	Go语言基本数据类型
	数据类型	无符号	位数
	int8	Yes		8
	int16	Yes		16
	int32	Yes		32
	int64	Yes		64
	uint8	No		8
	uint16	No		16
	uint32	No		32
	uint64	No		64
	int		Yes		等于cpu位数
	uint	No		等于cpu位数
	rune	Yes		与 int32 等价
	byte	No		与 uint8 等价
	uintptr	No		-
	 */
	var i int
	var f float64
	var bo bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, bo, s)

	var floatNum = 1.0 //基本数据类型没有double，浮点小数省略数据类型默认为float64
	fmt.Println(floatNum)

	//基本数据类型变量都直接指向内存的地址，属于值类型，值类型的数据存储在栈中
	var floatPos = &floatNum //指针获取内存地址
	fmt.Println(floatPos)
}
