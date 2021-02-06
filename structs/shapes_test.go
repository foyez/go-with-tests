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

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Area()
		want := 100.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793

		if got != want {
				t.Errorf("got %g want %g", got, want)
		}
	})

}

func ExamplePerimeter() {
	rectangle := Rectangle{10.0, 10.0}
	perimeter := Perimeter(rectangle)
	fmt.Println(perimeter)
	// Output: 40
}

func ExampleArea() {
	rectangle := Rectangle{10.0, 10.0}
	area := rectangle.Area()
	fmt.Println(area)
	// Output: 100
}

func ExampleArea_second() {
	circle := Circle{10.0}
	area := circle.Area()
	fmt.Println(area)
	// Output: 314.1592653589793
}
