package sorting

import "fmt"
import "strconv"

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	val := 0
	switch v := fnb.(type) {
	case FancyNumber:
		val, _ = strconv.Atoi(v.Value())
	default:
		val = 0
	}
	return val
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	return fmt.Sprintf("This is a fancy box containing the number %d.0", ExtractFancyNumber(fnb))
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	val := ""
	switch v := i.(type) {
	case int:
		val = DescribeNumber(float64(v))
	case float64:
		val = DescribeNumber(v)
	case NumberBox:
		val = DescribeNumberBox(v)
	case FancyNumberBox:
		val = DescribeFancyNumberBox(v)
	default:
		val = "Return to sender"
	}
	return val
}
