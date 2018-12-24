package string

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestSprintf(t *testing.T) {

	fmt.Printf("Hello %d\n", 23)
	fmt.Fprint(os.Stdout, "Hello ", 23, "\n")
	fmt.Println("Hello", 23)
	fmt.Println(fmt.Sprint("Hello ", 23))

	fmt.Println( 0 % 3)
	fmt.Println( 1 % 3)
	fmt.Println( 2 % 3)
	fmt.Println( 3 % 3)
	fmt.Println( 4 % 3)

}

func TestIndex(t *testing.T) {
	const s, sep, want = "chicken", "ken", 4
	got := strings.Index(s, sep)
	if got != want {
		t.Errorf("Index(%q,%q) = %v; want %v", s, sep, got, want)
	}
}

func TestIndex1(t *testing.T) {
	var tests = []struct {
		s   string
		sep string
		out int
	}{
		{"", "", 0},
		{"", "a", -1},
		{"fo", "foo", -1},
		{"foo", "foo", 0},
		{"oofofoofooo", "f", 2},
		// etc
	}

	for _, test := range tests {
		actual := strings.Index(test.s, test.sep)
		if actual != test.out {
			t.Errorf("Index(%q,%q) = %v; want %v",
				test.s, test.sep, actual, test.out)
		}
	}

}
