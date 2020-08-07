package main

// SpecialArray - test
type SpecialArray []interface{}

// Tip: Each element of `array` can either be cast
// to `SpecialArray` or `int`.
func productSum(array []interface{}) int {
	return helper(array, 1)
}

func helper(array SpecialArray, mul int) int {
	sum := 0
	for _, e := range array {
		if el, ok := e.(SpecialArray); ok {
			sum += helper(el, mul+1)
		} else if el, ok := e.(int); ok {
			sum += el
		}
	}
	return sum * mul
}
