package misc

// SomeCalculation does some calculation
func SomeCalculation(x int, y int) int {
	y = 99

	yPtr := &y

	return x + y + *yPtr
}
