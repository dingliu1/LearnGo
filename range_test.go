package main

import (
	"fmt"
	"testing"
	"time"
)

func TestRange(t *testing.T) {
	var m = [...]int{1, 2, 3, 4, 5}
	for i, v := range m {
		go func() {
			time.Sleep(time.Second * 3)
			fmt.Println(i, v)
		}()
	}
	time.Sleep(time.Second * 10)
}

func TestRangeRight(t *testing.T) {
	var m = [...]int{1, 2, 3, 4, 5}
	for i, v := range m {
		go func(i int,v int) {
			time.Sleep(time.Second * 3)
			fmt.Println(i, v)
		}(i,v)
	}
	time.Sleep(time.Second * 10)
}

func TestInRangeRemoveElements(t *testing.T) {
	var m = map[string]int{
		"tony": 21,
		"tom":  22,
		"jim":  23,
	}
	counter:=0
	for k, v := range m {
		if counter == 0{
			delete(m, "tony")
			m["lucky"] = 24
		}
		counter ++
		fmt.Println(k, v)
	}
	fmt.Println("counter is:",counter)
}

