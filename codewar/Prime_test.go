package main

import (
	"testing"
	"fmt"
	"time"
)

func TestFindPrime(t *testing.T){

	ret:=findPrime(10)
	fmt.Println(ret)

}

var channel chan int

func findPrime(number int) []int {

	array:=[]int{}
	for i:=1;i<=number;i++ {
		array=append(array, i)
	}

	fmt.Println(array)
	channel = make(chan int, number)

	resultArray:=[]int{}

	for j:=1;j<=number;j++ {
		go isPrime(j, channel)
	}

	go func() {

		for{
			result := <-channel
			resultArray = append(resultArray, result)
		}
		fmt.Println("i done")

	}()

	time.Sleep(3*time.Second)
	fmt.Println("end")
	return resultArray
}


func isPrime(number int, channel chan int) {

	var isPrime = true
	for i:=2;i<number;i++ {
		if number % i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime{
		channel<- number
	}
}



//public List<Integer> finePrime(Integer n) throws Exception {
//	List<Integer> primes = new ArrayList<Integer>();
//	if (n < 1) return primes;
//
//	for (int i = 1; i <= n; i++) {
//	primes.add(i);
//	}
//
//	List<Integer> ret = new ArrayList<Integer>();
//
//	for (Integer prime : primes) {
//	boolean isPrime = true;
//	for (int i = 2; i < prime; i++) {
//	if (prime % i == 0) {
//	isPrime = false;
//	break;
//	}
//	}
//	if(isPrime) ret.add(prime);
//	}
//	return ret;
//	}
