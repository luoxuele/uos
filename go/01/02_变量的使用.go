package main

import "fmt"


func main(){
	// 1. 声明并赋值
	var a int = 10
	fmt.Println("a = ",a)

	//先声明，后赋值
	var b, c int
	b= 20
	c = 30;
	fmt.Println("c = ",c," b = ",b)

	//3. 自动推导类型，必须初始化
	d := 40
	fmt.Println("d = ",d)





}