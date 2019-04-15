package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//通过时间创建随机种子
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	//创建数组
	var numbers [10]int
	//初始化一个切片(slice)
	numarr := []float32{10.1, 12.3, 5.2, 5.6, 12.3, 45.124, 3.2, 0.546}

	//类似传统的for循环
	for i := 0; i < 10; i++ {
		numbers[i] = r.Intn(100)
	}

	//类似传统的while循环
	i := 0
	for i < 5 {
		fmt.Println("类似while循环")
		i++
	}

	//循环遍历所有元素
	for i, value := range numbers {
		fmt.Printf("第%d位的数据是%d\n", i, value)
	}
	for _, value := range numarr {
		fmt.Printf("数据是%f\n", value)
	}
}
