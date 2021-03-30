package main

import (
	"fmt"
	"math"
	"os"
	"time"
)

func main() {
	os.Clearenv()
	a := 5
	fmt.Println(a)
	var pa *int
	pa = &a          //&a代表a的地址
	fmt.Println(pa)  //打印p指向的地址
	fmt.Println(*pa) //打印pa指向的值
	fmt.Print(math.Pi)
	var b = "hello go的世界即将来临，区块链需要你！"
	const yjm string = "\nyjm:你还是太弱了" //const定义常量
	fmt.Println(yjm)
	fmt.Println(b)
	fmt.Printf("hello word\nyjm,welcome you coming!\n")
	switch time.Now().Weekday() { //判断今天是星期六或者是星期日
	case time.Saturday, time.Sunday:
		fmt.Println("周末我爱你！")
	}
	t := time.Now()
	switch {
	case t.Hour() > 20:
		fmt.Println("睡觉觉！")
	case t.Hour() < 20:
		fmt.Println("睡不着！")
	}
	q := make([]string, 3)
	q[0] = "0"
	q[1] = "1"
	q[2] = "2"
	fmt.Println("切片", q)
	fmt.Println("长度", len(q))
	q = append(q, "4")
	fmt.Println("切片追加内容", q)
	c := make([]string, len(q))
	copy(c, q)
	fmt.Println("新切片内容", c)
	l := q[2:4]
	fmt.Println("长度", l)
	d := q[:4]
	fmt.Println("数据0124", d)
	d = q[2:]
	fmt.Println("数据24", d)
	shuZu := []string{"a", "b", "c", "d", "e"}
	fmt.Println("数组数据", shuZu)
	fmt.Println("切片数据shuZu[2:4]", shuZu[2:4])
}
