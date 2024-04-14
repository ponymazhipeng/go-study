// package 定义包名 main 包名
package main

import (
	"fmt"
	"net"
)

var age int
var level = 1
var (
	a int
	b string
	c []float32
)

//不能
//aa :=1

// func 定义函数 main 函数名
func main() {
	// fmt 包名 . 调用 Print 函数,并且输出定义的字符串
	//fmt.Print("Hello Golang")
	//fmt.Println(age)
	//fmt.Printf("%T", level)
	//%d 整数占位符，%s 字符串占位符， %f 浮点数占位符(默认精度为6)
	//fmt.Printf("%d,%s,%f", a, b, c)

	//aa := 1
	//fmt.Println(aa)
	//var conn net.Conn
	//var err error
	//conn, err = net.Dial("tcp", "127.0.0.1:8080")
	//fmt.Println(conn)
	//fmt.Println(err)
	//conn, err := net.Dial("tcp", "127.0.0.1:8080")
	//conn1, err := net.Dial("tcp", "127.0.0.1:8080")
	//fmt.Println(conn)
	//fmt.Println(conn1)
	//fmt.Println(err)
	//change3()
	//conn, err := net.Dial("tcp", "127.0.0.1:8080")
	//如果不想接收err的值，那么可以使用_表示，这就是匿名变量
	//conn, _ := net.Dial("tcp", "127.0.0.1:8080")
	//fmt.Println(conn)
	conn, _ := net.Dial("tcp", "127.0.0.1:8080")
	//匿名变量可以重复声明
	conn1, _ := net.Dial("tcp", "127.0.0.1:8080")
	// 匿名变量不可以直接开头
	// _ :=1
	fmt.Println(conn)
	fmt.Println(conn1)
}

// 交换
func change1() {
	a := 100
	b := 200
	//var b = 200
	var c int
	c = a
	a = b
	b = c
	fmt.Printf("a=%d,b=%d", a, b)
}

func change2() {
	a := 100
	b := 200

	a = a ^ b
	b = b ^ a
	a = a ^ b

	fmt.Printf("a=%d,b=%d", a, b)
}

func change3() {
	a := 100
	b := 200

	b, a = a, b

	fmt.Printf("a=%d,b=%d", a, b)
}
