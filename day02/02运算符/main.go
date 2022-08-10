package main

import "fmt"

func main() {
	// (1) 科学运算
	var x, y = 10, 20
	fmt.Println(x + y) // 30
	fmt.Println(x - y) // -10
	fmt.Println(x * y) // 200
	fmt.Println(x / y) // 0
	fmt.Println(x % y) // 10
	// 计算一个数字是奇数还是偶数
	fmt.Println(x%2 == 0)

	// (2) 比较运算符 返回的布尔值。
	fmt.Println(x > y)
	fmt.Println(x == y)
	fmt.Println(x != y)

	// (3) 逻辑运算
	// 与运算  真与真为 真、真与假为 假、假与假为 假
	// 或运算  真与真为 真、真与假为 真、假与假为 假
	// 非运算  非真为假 、非假为真。
	fmt.Println(true && false)  //false
	fmt.Println(true || false)  // true
	fmt.Println(!true || false) // false

}
