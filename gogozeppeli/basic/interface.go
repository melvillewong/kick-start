package basic

import "math"

type shape interface {
	area() float64
	perimeter() float64
}

type recta struct {
	width, height float64
}

func (r recta) area() float64 {
	return r.width * r.height
}

func (r recta) perimeter() float64 {
  return 2 * (r.width + r.height)
}

type circle struct {
  radius float64
}

func (c circle) area() float64 {
  return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
  return 2 * c.radius * math.Pi
}

func interface() {

}
