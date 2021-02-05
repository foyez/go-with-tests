package iteration

// const repeatedCount = 5

// Repeat takes two params and returns repeated string
func Repeat(character string, repeatedCount int) string {
	var repeated string

	for i := 0; i < repeatedCount; i++ {
		repeated += character
	}

	return repeated
}