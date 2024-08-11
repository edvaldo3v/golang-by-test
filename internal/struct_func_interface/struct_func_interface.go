package struct_func_interface

import "math"

type Forma interface {
	Area() float64
}

type Rectangle struct {
	height float64
	width  float64
}

type Circle struct {
	ray float64
}

func (r Rectangle) Parimeter() float64 {
	return 2 * (r.height + r.width)
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func (c Circle) Parimeter() float64 {
	return 2 * c.ray
}

func (c Circle) Area() float64 {
	return math.Pi * c.ray * c.ray
}
