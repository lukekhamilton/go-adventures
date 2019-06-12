package main

import (
	"fmt"

	"./pkg"
)

var xyz = 4

// Pkg ...
func Pkg() {
	var x = 1
	var y int
	var ip *int

	ip = &x
	y = *ip

	fmt.Printf("%+v\n", x)
	fmt.Printf("%+v\n", y)
	fmt.Printf("%+v\n", *ip)
	fmt.Printf("%+v\n", ip)
	fmt.Printf("%+v\n", &x)
	fmt.Printf("%+v\n", &y)

	x = 2

	fmt.Printf("%+v\n", x)
	fmt.Printf("%+v\n", y)
	fmt.Printf("%+v\n", *ip)
	fmt.Printf("%+v\n", ip)
	fmt.Printf("%+v\n", &x)
	fmt.Printf("%+v\n", &y)

	fmt.Println("*****************************************")

	ptr := new(int)
	*ptr = 5

	fmt.Println(*ptr)

	fmt.Println(ptr)

	fmt.Printf("%+v\n", xyz)

	_ = func() {
		fmt.Printf("%+v\n", xyz)
	}

	fmt.Printf("%+v\n", pkg.Hello())
	pkg.Main()
	fmt.Printf("%+v\n", pkg.X)

	fn := pkg.Hello
	fmt.Printf("%+v\n", fn())
}
