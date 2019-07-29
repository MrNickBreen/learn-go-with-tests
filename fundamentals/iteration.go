package fundamentals

// Note: You can't have optional params!
func Repeat(char string, n int) (repeated string) {
	if n < 0 {
		n = 0
	}
	for i := 0; i < n; i++ {
		repeated += char
	}
	return
}
