package raindrops

import "strconv"

func Convert(number int) string {
	o := ""
	if number%3 == 0 {
		o = "Pling"
	}
	if number%5 == 0 {
		o += "Plang"
	}
	if number%7 == 0 {
		o += "Plong"
	}
	if o != "" {
		return o
	}
	return strconv.Itoa(number)
}
