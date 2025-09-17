package main

import "fmt"

type complexNumber interface {
	create()
	add(other *Complex) *Complex
}

type Complex struct {
	real      float64
	imaginary float64
}

func newComplex(real float64, imaginary float64) *Complex {
	c := Complex{
		real:      real,
		imaginary: imaginary,
	}

	return &c
}

func (c *Complex) create() {
	fmt.Println("Complex number created:", c.real, "+", c.imaginary, "i")
}

func (c *Complex) add(other *Complex) *Complex {
	sum := Complex{
		real:      c.real + other.real,
		imaginary: c.imaginary + other.imaginary,
	}
	fmt.Println("Complex number added")
	return &sum

}

func main() {
	fmt.Println("Lets implement a complex number class in golang")

	c1 := newComplex(1, 2)
	c2 := newComplex(3, 4)

	c1.create()
	c2.create()

	sum := c1.add(c2)
	sum.create()

}
