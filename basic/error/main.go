package main

import "fmt"

type Divide struct {
	n1 int
	n2 int
}

func (d *Divide) print() {
	fmt.Printf("%+p\n", d)
}

func (d *Divide) Error() {
	fmt.Printf("error,%+p\n", d)
}

func main() {
	d := &Divide{
		1, 2,
	}
	fmt.Printf("%p\n", d)
	d.print()
	d.Error()
}
