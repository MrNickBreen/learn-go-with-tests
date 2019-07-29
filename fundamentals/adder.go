package fundamentals

// When you have more than one argument of the same type
// (in our case two integers) rather than having (x int, y int)
// you can shorten it to (x, y int).
// See this documentation http://localhost:6060/pkg/github.com/mrnickbreen/fundamentals/
// when running `godoc -http=:6060`
func Add(x, y int) int {
	return x + y
}
