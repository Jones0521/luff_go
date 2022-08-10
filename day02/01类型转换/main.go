package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	// (1) 整型之间的转换 int8 int16 int32 int64
	var x int8 = 10
	var y int16 = 20
	// 注意转换后的范围
	fmt.Println(x + int8(y))

	// (2) 字符串与整数之间的转换: strconv
	var ageStr = "32"
	var ageInt, err = strconv.Atoi(ageStr)
	if err != nil {
		panic(err)
	}
	//fmt.Println(err, reflect.TypeOf(err))
	fmt.Println(ageInt, reflect.TypeOf(ageInt))
	price := 100
	priceStr := strconv.Itoa(price)
	fmt.Println(priceStr, reflect.TypeOf(priceStr))

	// (3) strconv Parse 系列函数
	// 将字符串转换为整数 ParseInt: s 字符串 base 进制 bitSize 字节大小限制
	ret, _ := strconv.ParseInt("28", 10, 8)
	fmt.Println(ret, reflect.TypeOf(ret))

	// 将字符串转换为浮点型
	ret2, _ := strconv.ParseFloat("3.1415926", 64)
	fmt.Println(ret2, reflect.TypeOf(ret2))

	// 将字符串转换为 bool值
	// strconv.ParseBool 它接受1、0、t、f、T、F、true、false、True、False、TRUE、FALSE; 否则返回错误.
	b1, _ := strconv.ParseBool("true")
	fmt.Println(b1, reflect.TypeOf(b1))
	b2, _ := strconv.ParseBool("T")
	fmt.Println(b2, reflect.TypeOf(b2))
	b3, _ := strconv.ParseBool("1")
	fmt.Println(b3, reflect.TypeOf(b3))
	b4, _ := strconv.ParseBool("100")
	fmt.Println(b4, reflect.TypeOf(b4))
}
