package iteration

const repeatedCount = 5

// Repeat takes one string and returns repeated string
func Repeat(character string) string {
	var repeated string

	for i := 0; i < repeatedCount; i++ {
		repeated += character
	}

	return repeated
}