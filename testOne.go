package main

import (
	"fmt"
	_ "fmt"
)

//递归函数自调用
func sum(num int) int {
	if num == 1 {
		return num
	} //初始值为输入的参数值如果大于一先进行if外的语句num自增
	return sum(num-1) + num
} //举一反一递归相乘、
func ride(num int) int {
	if num == 1 {
		return num
	}
	return ride(num-1) * num
}
func main() {
	fmt.Println("递归相加值为", sum(10))
	fmt.Println("递归相乘值为", ride(5))
	//定义一个key:value的哈希表
	m := make(map[string]int)
	m["k1"] = 1
	m["k2"] = 2
	fmt.Println("m:", m)
	delete(m, "k1") //删除
	//len(m)长度
	fmt.Println(m)
	//定义加初始化
	n := map[string]int{"foo1": 1, "foot2": 2}
	fmt.Println(n)
	//定义一个数组
	nums := []int{1, 2, 3}
	sum := 0
	//range循环数组
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	for i, num := range nums {
		fmt.Println("index:", i, "num", num)
	} //迭代循环一个map(典)
	kvs := map[string]string{"a": "apple", "b": "banana", "o": "orange"}
	for k, v := range kvs {
		fmt.Printf("%s->%s\n", k, v)
	} //迭代循环一个字符串
	for a, b := range "Iloveyou" {
		fmt.Println(a, string(b))
	}
	//----------------------------------
	//----------------------------------
	//函数定义
	/*func plus2(a int,b int)int{
	return a+b
	}
	*/
}
