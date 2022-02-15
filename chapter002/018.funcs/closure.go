package main

import "fmt"

var counter int

/*
go语言支持函数式编程:
	支持将一个函数作为另一个函数的参数,也支持将一个函数作为另一个函数的返回值.

闭包closure:
	一个外层函数中,有内层函数,该内层函数中,会操作外层函数的局部变量(外层函数中的参数,或者外层函数中直接定义的变量
),并且该外层函数的返回值就是这个内层函数
这个内层函数和外层函数的局部变量,统称为闭包结构.
*/

func main() { //要复习,不是很理解.
	//a := calSum()
	showUsedTimes()
	fmt.Println(counter)
	b := genImprovementFunc()
	b(1.0)
	genImprovementFunc()
	fmt.Println(b)

	closureMain()

}

func calSum(nums ...int) (sum int) {
	for _, item := range nums {
		sum += item
	}
	counter++
	return
}

func showUsedTimes() {
	fmt.Println("used: ", counter)

}
func genImprovementFunc() func(percentage float64) { //外层函数 //匿名函数是一个"内联"语句或表达式.
	//匿名函数的优越性在于可以直接使用函数内的变量,不必申明.
	//错误,这个不是函数方法
	//方法是一个包含了接受者的函数
	//func (c Circle) getArea() float64{}   //该 method 属于 Circle 类型对象中的方法
	base := 1000.0                    //定义一个局部变量
	return func(percentage float64) { //返回该匿名函数
		base = base * (1 + percentage)
		fmt.Println("new: ", base)
	}
}

func closureMain() {
	imp := genImprovementFunc()
	imp(0.1)
	imp(0.1)
	imp(0.1)
	imp(0.1)

}
