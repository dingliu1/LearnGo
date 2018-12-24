package _select

import (
	"fmt"
	"testing"
	"time"
)

func TestSelectBlockForNoDefaultStatement(t *testing.T) {
	chan1 := make(chan int)
	chan2 := make(chan int)

	writeFlag := false
	go func() {
		for {
			if writeFlag {
				chan1 <- 1
			}
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			if writeFlag {
				chan2 <- 1
			}
			time.Sleep(time.Second)
		}
	}()

	select {
	case <-chan1:
		fmt.Println("chan1 ready.")
	case <-chan2:
		fmt.Println("chan2 ready.")
	}

	fmt.Println("main exit.")
}

func TestSelectWithClose(t *testing.T) {
	chan1 := make(chan int)
	chan2 := make(chan int)

	go func() {
		close(chan1)
	}()

	go func() {
		close(chan2)
	}()

	select {
	case <-chan1:
		fmt.Println("chan1 ready.")
	case <-chan2:
		fmt.Println("chan2 ready.")
	}

	fmt.Println("main exit.")
}


func TestSelectPanic(t *testing.T) {
	select {

	}
}

func takeARecvChannel() chan int {
	fmt.Println("invoke takeARecvChannel")
	c := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		c <- 1
	}()

	return c
}

func getAStorageArr() *[5]int {
	fmt.Println("invoke getAStorageArr")
	var a [5]int
	return &a
}

func takeASendChannel() chan int {
	fmt.Println("invoke takeASendChannel")
	return make(chan int)
}

func getANumToChannel() int {
	fmt.Println("invoke getANumToChannel")
	return 2
}

func TestSelectOrder(t *testing.T) {
	select {
	//recv channels
	case (getAStorageArr())[0] = <-takeARecvChannel():
		fmt.Println("recv something from a recv channel")

		//send channels
	case takeASendChannel() <- getANumToChannel():
		fmt.Println("send something to a send channel")
	}
}