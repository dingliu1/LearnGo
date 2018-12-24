package main

import (
	"fmt"
	"reflect"
	"testing"
)

type A interface {
	method1() int
	method2() string
}

type implA struct {
}

func (a *implA) method1() int {
	return 1
}

func (c *implA) method2() string {
	return "coolA"
}

type B interface {
	A
	method3()
}

type implB struct {
}


func (implB) Method8(){

}
func (a *implB) method1() int {
	return 2
}

func (c *implB) method2() string {
	return "coolB"
}

func (c *implB) method3() {

}

func TestMandatoryCovert(t *testing.T) {
	var i = new(implB)
	fmt.Println(i.method2())

	var j = new(implA)
	ret := j.method2()
	fmt.Println(ret)
}

func TestReflect(t *testing.T) {
	var i *A
	elemType := reflect.TypeOf(i).Elem()
	n := elemType.NumMethod()
	for i := 0; i < n; i++ {
		fmt.Println(elemType.Method(i).Name)
	}
}
