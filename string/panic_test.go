package string

import "testing"

func TestNormalDefer(t *testing.T) {

	firstName := "Liu"
	lastName := "Ding"
	doIt(&firstName, &lastName, deferedFunc, false)
}

/**
  when panic occurs,
  1. the defer function of all ancessors are called
  2. report panic
  3. print stack trace
*/
func TestPanicDefer(t *testing.T) {

	firstName := "Liu"
	doIt(&firstName, nil, deferedFunc, false)
}
