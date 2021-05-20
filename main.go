package main

import "fmt"

type obj struct {
	Count int
}

func main()  {
	fmt.Println("Go Design Patterns")
	x := new(obj)
	x.Count = 1
	p := &x
	a := *p
	b := *p
	a.Count = 2

	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", &b)
	fmt.Println(a, b)
}
