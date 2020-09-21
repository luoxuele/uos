package main

import "fmt"

func test()(a, b, c int){
	return 1, 2, 3
}

func main()  {
	a, b := 10,20

	//交换两个变量值
	a,b = b,a

	fmt.Println("a = ",a, " b = ",b)

	//匿名变量，丢弃数据不处理，_匿名变量配合函数返回值使用才有优势
	//tmp, _ = i, j
	var i, j int
	_, i, j = test() //return 1 2 3
	fmt.Println("i = ",i, " j = ",j)



	
}