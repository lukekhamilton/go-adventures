package main

import (
	"fmt"
)

// String ...
func String() {

	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	// const sample = "Hello World \xe2\x8c\x98 \u2318 \x20"

	fmt.Println(sample)
	fmt.Println(len(sample))

	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i])
	}

	fmt.Printf("\n% x\n", sample)

	fmt.Printf("%q\n", sample)

	fmt.Printf("%+q\n", sample)

	const placeOfInterest = `⌘`

	fmt.Printf("plain string: ")
	fmt.Printf("%s\n", placeOfInterest)

	fmt.Printf("quoted string: %+q\n", placeOfInterest)

	fmt.Printf("len: %x\n", len(placeOfInterest))
	fmt.Printf("hex bytes: ")
	for i := 0; i < len(placeOfInterest); i++ {
		fmt.Printf("%x ", placeOfInterest[i])
	}
}
