package string

import (
	"fmt"
	"testing"
)

/**
Go source code is always UTF-8.
A string holds arbitrary bytes.
A string literal, absent byte-level escapes, always holds valid UTF-8 sequences.
Those sequences represent Unicode code points, called runes.
No guarantee is made in Go that characters in strings are normalized.
*/
func TestSample(t *testing.T) {
	sample1()
	fmt.Println("---------sample2-----------------")
	sample2()
	fmt.Println("---------sample3-----------------")
	sample3()
}
