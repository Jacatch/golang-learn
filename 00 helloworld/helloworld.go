package main

import (
	"fmt"
	"unicode/utf8"
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

	var str = "hello 你好"
	fmt.Println(len(str))                    //go语言底层是使用byte数组，byte数组的中文在unicode编码下占两位，utf-8编码下占三位，故长度为12
	fmt.Println(utf8.RuneCountInString(str)) //读取真实字符长度
	runes := []rune(str)                     //直接转换为rune数组
	fmt.Println(len(runes))

	//初始化一个切片
	numarr := []float32{10.1, 12.3, 5.2, 5.6, 12.3, 45.124, 3.2, 0.546}
	slice := numarr[0:3] //切片指向numarr的第一个元素，长度位3，容量为8
	slice2 := slice[1:2] //slice2也是基于numarr的切片,但指向第二个元素，容量为在那之后的7
	fmt.Printf("len=%d cap=%d slice=%v\n", len(slice), cap(slice), slice)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(slice2), cap(slice2), slice2)
	fmt.Println(numarr)
	slice2[0] = 0.00 //因为切片是对数组的引用，所以对切片数据的修改也会影响到数组
	fmt.Println(numarr)

	numarr2 := make([]float32, len(numarr), (cap(numarr))*2) //初始化一个新切片，容量为原切片的两倍
	fmt.Printf("len=%d cap=%d slice=%v\n", len(numarr2), cap(numarr2), numarr2)
	copy(numarr2, numarr) //将原切片复制给新切片
	fmt.Printf("len=%d cap=%d slice=%v\n", len(numarr2), cap(numarr2), numarr2)
	numarr2 = append(numarr2, 2.2) //切片允许添加元素
	fmt.Printf("len=%d cap=%d slice=%v\n", len(numarr2), cap(numarr2), numarr2)
	numarr2 = append(numarr2, 2.2, 1.0, 2.0, 1.2, 3.2, 4.6, 8.1, 4.6, 5.4) //允许添加多个元素，超出容量时自动进行扩容
	fmt.Printf("len=%d cap=%d slice=%v\n", len(numarr2), cap(numarr2), numarr2)
	/**
	关于切片(slice)的扩容
	1. 当每次扩大的数量不超过原先容量的两倍，并总长度不超过1024时，扩大为原来的两倍
	2. 切片一次新增个数超过原来1倍，但不超过1024个，且增加后总长度小于1024个，这种情况下扩容后比实际具有的总长度还要大一些。
	3. 原切片长度超过1024时，一次增加容量不是2倍而是0.25倍，每次超过预定的都是0.25累乘
	*/
	r := make([]float32, 10)
	fmt.Printf("len=%d cap=%d\n", len(r), cap(r))
	r1 := append(r, make([]float32, 15)...)
	fmt.Printf("len=%d cap=%d\n", len(r1), cap(r1))

}
