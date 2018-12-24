package _select

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Print("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	q := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		q <- 8
	}()

	fibonacci(c, q)

}
