package main

import (
	"fmt"
	"strings"
)

func main() {
	/* // 案例1
	fmt.Println("请输入您的生日, 按格式: 年-月-日")
	var birth string
	fmt.Scan(&birth)
	birthSlice := strings.Split(birth, "-")
	fmt.Println("birthSlice:", birthSlice)

	fmt.Printf("您的生日是%s年-%s月-%s日", birthSlice[0], birthSlice[1], birthSlice[2])
	*/
	// 案例2
	var name string
	fmt.Println("请输入一个英文名: ")
	fmt.Scan(&name)
	//var b = strings.HasPrefix(name,"a") || strings.HasPrefix(name,"a")
	UpperName := strings.ToUpper(name)
	var b = strings.HasPrefix(UpperName, "A")
	fmt.Println(b)
}
