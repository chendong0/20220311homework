package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)

	i := 1
	j := &i
	d := &j

	fmt.Println(i)
	fmt.Printf("%x\n", &i)
	fmt.Printf("%x\n", &j)
	fmt.Printf("%x\n", j)
	fmt.Printf("%v\n", d)

}
