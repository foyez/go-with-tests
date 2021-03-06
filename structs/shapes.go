package structs

import "math"

// Shape has Area method
type Shape interface {
	Area() float64
}

// Rectangle has the dimensions of a rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Area returns the area of the rectangle (method - func (receiverName ReceiverType) MethodName(args))
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle has Radius property of a circle
type Circle struct {
	Radius float64
}

// Area returns the area of the circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Triangle has base and height property of a triangle
type Triangle struct {
	Base   float64
	Height float64
}

// Area return the area of the triangle
func (t Triangle) Area() float64 {
	return .5 * t.Base * t.Height
}

// Perimeter takes rectangle type and returns perimeter
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}
