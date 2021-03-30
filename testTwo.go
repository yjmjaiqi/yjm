package main

import (
	"fmt"
	_ "fmt"
)

//值传递
func byValue(a int) {
	a = 0
}
func byPoint(b *int) {
	*b = 0
}
func main() {
	i := 10 //值传递给byValue
	fmt.Println("初始值", i)
	byValue(i)
	fmt.Println("byValue:", i)
	//地址传递给byPoint
	byPoint(&i)
	fmt.Println(i) //0
	fmt.Println("byPoint", &i)
}
