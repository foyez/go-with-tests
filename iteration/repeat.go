package iteration

// Repeat takes one string and returns repeated string
func Repeat(character string) string {
	var repeated string

	for i := 0; i < 5; i++ {
		repeated += character
	}

	return repeated
}