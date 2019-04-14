package main

import (
	"fmt"
	"unsafe"
)

//使用常量作为枚举
const (
	Unknown = 0
	Female  = 1
	Male    = 2
)

//常量表达式中，函数必须为内置函数，不能为自定义函数
const (
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(a) //字符串类型在 go 里是个结构, 包含指向底层数组的指针和长度,这两部分每部分都是 8 个字节，所以字符串类型大小为16
)

//iota时内置变量，在const中第一次出现时被重置为0，之后每个值iota都加一，便于进行枚举
const (
	ai = iota
	bi = iota
	ci = iota
)

const (
	di = iota    //iota = 0
	ei = "hello" //iota ++
	fi = 15.2    //iota ++
	gi           //iota ++ , gi = fi
	hi           //iota ++ , hi = fi
	ii = iota    //iota ++ , ii = iota
	ji           //iota ++ , ji = iota
)

const (
	ki = 2 * iota
	li
	mi
	ni = 3 << iota
	oi
	pi
)

func main() {
	//常量定义
	const WIDTH, HEIGTH = 640, 360
	const TEXT string = "面积"
	area := WIDTH * HEIGTH
	fmt.Println(TEXT, area)
	fmt.Println(a, b, c)

	//iota测试
	fmt.Println(ai, bi, ci)
	fmt.Println(di, ei, fi, gi, hi, ii, ji)
	fmt.Println(ki, li, mi, ni, oi, pi)

	//条件语句与格式化输入
	var a int
	//value, err = fmt.Scan(&a)
	_, err := fmt.Scan(&a) //接受输入，并接收错误信息，空白符表示忽略返回
	if err != nil {
		fmt.Println("输入值格式错误", err)
		return
	}
	fmt.Println(a)
	if a > 5 {
		fmt.Println("a > 5")
	} else {
		fmt.Println("a < 5")
	}

	//switch的使用
	var b bool
	_, err = fmt.Scan(&b) //对bool进行输入时，可以输入false与true，若输入数字，则1为true，其他数字均为false
	if err != nil {
		fmt.Println("输入值格式错误", err)
		return
	}
	switch a {
	case 1:
		fmt.Println("1") //go语言的switch不需要break
	case 2, 3:
		fmt.Println("a的值为2或3") //可以同时有多个条件
	default:
		fmt.Println("？？？")
	}

	switch {
	//switch可以在switch后不接条件
	case b == true:
		fmt.Println("b的值为true")
		fallthrough //使用fallthrough，可以强制执行后面的case而不进行条件判断
	case b == false:
		fmt.Println("b的值为false")
	}

	//select的使用
	/**
	1. select要求每一个case都是一个通信
	2. 每一个channel表达式都会被求值
	3. 所有被发送的表达式都会被求职
	4. 如果任意某个通信可以进行，它就执行，其他被忽略。
	5. 如果有多个 case 都可以运行，Select 会随机公平地选出一个执行。其他不会执行。
		否则：
		(1) 如果有 default 子句，则执行该语句。
		(2) 如果没有 default 子句，select 将阻塞，直到某个通信可以运行；Go 不会重新对 channel 或值进行求值。
	*/
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-c3): // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}
}
