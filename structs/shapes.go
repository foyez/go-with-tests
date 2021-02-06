package structs

// Perimeter takes two float64 and returns perimeter in float64
func Perimeter(height, width float64) float64 {
	return 2 * (height + width)
}

// Area takes height, width and returns area
func Area(height, width float64) float64 {
	return height * width
}