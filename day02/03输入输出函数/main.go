package main

import "fmt"

func main() {
	// 输出函数
	// (1) Print 和 Println
	var name, age = "yuan", 32
	fmt.Println("hello world")
	fmt.Println(name, age)
	fmt.Println("姓名:", name, "年龄:", age)
	//fmt.Print(name, age)

	// (2) printf: 标准格式化输出
	var isMarried = false
	fmt.Printf("姓名: %s,年龄: %d,婚否: %t \n", name, age, isMarried)
	// (3) Sprintf:
	s := fmt.Sprintf("姓名: %s,年龄: %d,婚否: %t \n", name, age, isMarried)
	fmt.Println(s)

	// 输入函数
	// (1) fmt.Scan
	//var iName string
	//fmt.Scan(&iName) // 等等用户在命令行输入
	//fmt.Println(iName)
	//fmt.Println("end")

	// (2) fmt.Scanln
	//var iName2 string
	//var age2 int8
	//fmt.Print("请输入姓名和年龄, 用空格分隔:")
	//fmt.Scanln(&iName2, &age2) // 等等用户在命令行输入
	//fmt.Println(iName2, age2)
	//fmt.Println("end")

	// (3) fmt.Scanf
	var a, b int
	fmt.Println("请按指定格式输入: (例如1+1) ")
	fmt.Scanf("%d+%d", &a, &b)
	fmt.Println(a + b)
}
