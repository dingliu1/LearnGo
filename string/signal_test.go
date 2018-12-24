package string

import (
	"testing"
	"time"
)

func TestSysSignalHandleDemo(t *testing.T) {

	go sysSignalHandleDemo()
	time.Sleep(time.Hour) // make the main goroutine wait!
}
