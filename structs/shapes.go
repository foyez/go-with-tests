package structs

import "math"

// Circle has Radius property of a circle
type Circle struct {
	Radius float64
}

// Area returns the area of the circle
// method - func (receiverName ReceiverType) MethodName(args)
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Rectangle has the dimensions of a rectangle
type Rectangle struct {
	Width float64
	Height float64
}

// Area returns the area of the rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter takes two float64 and returns perimeter in float64
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Area takes height, width and returns area
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
