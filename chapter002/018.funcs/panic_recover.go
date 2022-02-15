package main

import "fmt"

func main() {
	panicAndRecover()

}

func panicAndRecover() {
	defer coverPanicUpgraded() //这样是能抓住严重错误的

	defer func() { //用闭包,抓不住严重错误
		coverPanicUpgraded()
	}()

	defer coverPanic()
	//var nameScore map[string]int = nil
	//var声明map未初始化默认为nil
	//如果不初始化map,name就会创建一个nil map,nil map不能用来存放键值对.
	//panic: assignment to entry in nil map

	nameScore := make(map[string]int)

	nameScore["吴军"] = 100
	for name := range nameScore {
		fmt.Println(name, "姓名是 ", nameScore[name]) //
	}

}
func coverPanic() { //未能抓住panic
	func() {
		if r := recover(); r != nil {
			fmt.Println("系统出现严重故障: ", r)

		}
	}()

}

func coverPanicUpgraded() { //未能抓住panic
	func() {
		if r := recover(); r != nil {
			fmt.Println("系统出严重故障: ", r)

		}
	}()

}
