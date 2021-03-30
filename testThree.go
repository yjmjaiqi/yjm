package main

import (
	"fmt"
	_ "fmt"
)

//声明一个user结构体
type user struct {
	name     string
	password string
	age      int
}

func main() { //实例化user结构体
	user1 := user{name: "yjm", password: "123", age: 19}
	fmt.Println(user1)
	//声明一个指向user的指针user2
	user2 := &user1
	fmt.Println(user2.name, user2.password, user2.age)
	//利用指针给user1赋值，会改变user1的内容
	user2.name = "hzs"
	user2.password = "321"
	user2.age = 18 //赋值会影响两个指针，地址指向
	fmt.Println(user1.name, user1.password, user1.age)
	fmt.Println(user2.name, user2.password, user2.age)
} //area() perim()面积 周长公式

//&指针*变量
