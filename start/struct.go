package main

import "fmt"

type location struct {
	name string
	lat  float64
	long float64
}

func main() {
	lats := []float64{-4.1919, -3.3232, 1.321323, 4.3112} // 创建切片
	longs := []float64{0.23412, 3.452341, 4.2341, 6.121324}

	// 通过结构类型创建变量, 并通过字面量赋值
	locations := []location{
		{name: "lalala", lat: lats[0], long: longs[0]},
		{name: "test2", lat: lats[1], long: longs[1]},
		{name: "test3", lat: lats[2], long: longs[2]},
	}

	fmt.Println(locations)
}
