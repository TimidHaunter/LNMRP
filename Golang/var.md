## var1

```go
package main

import "fmt"

/**
 * 定义主函数
 */
func main() {
	/**
	  变量：variable
	  概念：一小块内存，用于存储数据，在程序运行过程中数值可以改变

	  使用：
	  1.声明，也叫定义
	    第一种：var 变量名 数据类型
	           变量名 = 赋值

	           var 变量名 数据类型 = 赋值

	    第二种：类型推断，省略数据类型
	           var 变量名 = 赋值

	    第三种：简短声明，省略var
	           变量名 := 赋值

	  2.访问，赋值和取值
	  直接根据变量名访问
	*/

	//第一种：定义变量并且进行赋值（官方建议：应将变量声明与下一行的赋值合并）
	var num1 int
	num1 = 30
	fmt.Printf("num1的数值是：%d\n", num1)

	//写在一行
	var num2 int = 15
	fmt.Printf("num2的数值是：%d\n", num2)

	//第二种：类型推断
	var name = "王阳明"
	fmt.Printf("类型是：%T，数值是：%s\n", name, name)

	//第三种：简短定义，也叫简短声明
	sum := 100
	fmt.Println(sum)

	/**
	  多个变量同时定义
	*/
	var a, b, c int
	a = 1
	b = 2
	c = 3
	fmt.Println(a, b, c)

	var m, n int = 100, 200
	fmt.Println(m, n)

	var n1, f1, s1 = 100, 3.14, "Go"
	fmt.Println(n1, f1, s1)

	var (
		nickname = "咸蛋超人"
		age      = 18
		sex      = "男"
	)
	fmt.Printf("名字是：%s，年龄是：%d，性别是：%s\n", nickname, age, sex)
}
```



## var2

```go
package main

import "fmt"

func main() {
	var num int
	num = 100
	//&取地址符
	fmt.Printf("num的数值是：%d，地址是：%p\n", num, &num)

	num = 200
	fmt.Printf("num的数值是：%d，地址是：%p\n", num, &num)
	/*
		PS G:\data\Go\src\basic> go run .\var2.go
		num的数值是：100，地址是：0xc000012098
		num的数值是：200，地址是：0xc000012098
		地址没有变化
	*/
}
```

