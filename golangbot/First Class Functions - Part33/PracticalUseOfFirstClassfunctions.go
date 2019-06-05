package main

import "fmt"

type student struct {
	firstName string
	lastName  string
	grade     string
	country   string
}

type students []student

func filter(s []student, f func(student) bool) []student {
	var r []student
	for _, v := range s {
		if f(v) == true {
			r = append(r, v)
		}
	}
	return r
}

func (s *students) filter(f func(student) bool) students {
	var r students
	for _, v := range *s {
		if f(v) == true {
			r = append(r, v)
		}
	}
	return r
}

func callFilter1() {
	s1 := student{
		firstName: "Naveen",
		lastName:  "Ramanathan",
		grade:     "A",
		country:   "India",
	}
	s2 := student{
		firstName: "Samuel",
		lastName:  "Johnson",
		grade:     "B",
		country:   "USA",
	}

	s := []student{s1, s2}

	f := filter(s, func(s student) bool {
		if s.grade == "B" {
			return true
		}
		return false
	})

	fmt.Println("f: ", f)

	c := filter(s, func(s student) bool {
		if s.country == "India" {
			return true
		}
		return false
	})

	fmt.Println("c: ", c)
}

func callFilter2() {
	s1 := student{
		firstName: "Naveen",
		lastName:  "Ramanathan",
		grade:     "A",
		country:   "India",
	}
	s2 := student{
		firstName: "Samuel",
		lastName:  "Johnson",
		grade:     "B",
		country:   "USA",
	}

	s := students{s1, s2}

	As := s.filter(func(s student) bool {
		if s.grade == "A" {
			return true
		}
		return false
	})

	fmt.Println("A's: ", As)

}

func Map(s []int, f func(int) int) []int {
	var r []int
	for _, v := range s {
		r = append(r, f(v))
	}
	return r
}

type numbers []int

func (n numbers) Map(f func(int) int) []int {
	var r []int
	for _, v := range n {
		r = append(r, f(v))
	}
	return r
}

func callMap2() {
	a := numbers{1, 2, 3, 4, 5, 6, 7, 8, 9}

	r := a.Map(func(n int) int {
		return n * 3
	})

	fmt.Println("r: ", r)
}

func callMap1() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := Map(a, func(n int) int {
		return n * 5
	})

	fmt.Println("r: ", r)
}

type addd func(a, b int) int

func display(sum addd) {
	fmt.Println(sum(10, 10))
}

func callDisplay() {
	a := func(a, b int) int {
		return a + b
	}

	display(a)

	var x addd = func(a, b int) int {
		return a + b
	}

	fmt.Printf("x: %v\n", x(1, 2))

	display(x)
}

func main() {
	callDisplay()
}
