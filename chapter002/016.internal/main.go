package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	startTime := time.Now()
	defer func() {
		finishTime := time.Now()
		fmt.Println(startTime, finishTime)
		fmt.Println(finishTime.Sub(startTime))
	}()
	defer func() {
		fmt.Println("1")
	}()
	defer func() {
		fmt.Println("2")
	}()
	defer func() {
		fmt.Println("3")
	}()
	arr2 := testPanic()
	arr3 := make([]string, 3, 4)
	copy(arr3, arr2)

	//联系copy()方法
	p1 := copy(arr3, arr2)
	fmt.Println(p1)
	i := 3333
	j := &i
	//fmt.Printf("查看j的值%p",j)
	fmt.Println(reflect.TypeOf(j))
	//reflect包封装了反射相关的方法;
	//获取类型信息:reflect.TypeOf是静态的
	//获取值信息:reflect.ValueOf,是动态的
}
func testPanic() []string {
	arr2 := make([]string, 0, 4)
	fmt.Println("len: ", len(arr2), "cap :", cap(arr2))
	fmt.Println("default: ", arr2[0]) //panic: runtime error: index out of range [0] with length 0
	fmt.Println("default: ", arr2[1]) //panic: runtime error: index out of range [0] with length 0
	fmt.Println("default: ", arr2[2]) //panic: runtime error: index out of range [0] with length 0
	//切片长度为0,无法打印出default,引发panic
	return arr2

}
