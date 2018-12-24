package string

import (
	"fmt"
	"unicode/utf8"
)

func sample1() {

	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

	fmt.Println(sample)
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x", sample[i])
	}
	fmt.Println()
	// dumps out the sequential bytes of the string as hexadecimal digits, two per byte
	fmt.Printf("%x\n", sample)

	fmt.Printf("% x\n", sample)

	//escape any non-printable byte sequences in a string s
	fmt.Printf("%q\n", sample)

	//escape not only non-printable sequences, but also any non-ASCII bytes
	fmt.Printf("%+q\n", sample)
}

func sample2() {

	const placeOfInterest = `⌘`

	fmt.Printf("plain string: ")
	fmt.Printf("%s", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("quoted string: ")
	//escape not only non-printable sequences, but also any non-ASCII bytes
	fmt.Printf("%+q", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("hex bytes: ")
	for i := 0; i < len(placeOfInterest); i++ {
		fmt.Printf("%x ", placeOfInterest[i])
	}
	fmt.Printf("\n")
}

func sample3() {

	const nihongo = "日本語"
	for index, runeValue := range nihongo {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}

	for i, w := 0, 0; i < len(nihongo); i += w {
		runeValue, width := utf8.DecodeRuneInString(nihongo[i:])
		fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
		w = width
	}

}
