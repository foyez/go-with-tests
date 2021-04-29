package structs

import (
	"fmt"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})

}

func ExamplePerimeter() {
	rectangle := Rectangle{10.0, 10.0}
	perimeter := Perimeter(rectangle)
	fmt.Println(perimeter)
	// Output: 40
}

func ExampleRectangle_Area() {
	rectangle := Rectangle{10.0, 10.0}
	area := rectangle.Area()
	fmt.Println(area)
	// Output: 100
}

func ExampleCircle_Area() {
	circle := Circle{10.0}
	area := circle.Area()
	fmt.Println(area)
	// Output: 314.1592653589793
}
