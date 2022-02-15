package main

import "fmt"

func finished() {
	fmt.Println("Finished finding largest")

}

func largest(nums []int) {
	defer finished()
	fmt.Println("Started finding largest")
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v

		}
	}
	fmt.Println("Largest number in", nums, "is", max)
	//注意作用域的范围
}
func main() {

	a := 1
	b := 2
	defer fmt.Println(b)
	fmt.Println(a)

	nums := []int{10, 25, 33, 14, 45}
	largest(nums)
}
