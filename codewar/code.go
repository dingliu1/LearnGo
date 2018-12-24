package main

import (
	"fmt"
	"time"
)

type field struct {
	name string
}

func (p *field) print() {
	fmt.Println(p.name)
}

func main() {
	// 加号必须放到末尾
	var s = "中国人中国人中国人中国人中国人中国人中国人中国人中国人中国人中国人中国人中国人中国人" +
		"中国人中国人中国人中国人中国人中国人中国人中国人中国人中国人中国人中国人中国人中国人中国人中国人中国人中国人中国人"
	var a = [...]int{5, 6, 7,
	 	8}

	fmt.Println(s,a)
	data1 := []*field{{"one"}, {"two"}, {"three"}}
	for _, v := range data1 {
		go v.print()
	}

	data2 := []field{{"four"}, {"five"}, {"six"}}
	for _, v := range data2 {
		go v.print()
	}

	time.Sleep(3 * time.Second)
}
