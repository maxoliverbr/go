package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("true")
	} else if n == 1 {
		return 0, nil
	}
	i := 1
	for {
		if n%2 == 0 {
			n = n / 2
		} else if n != 1 {
			n = 3*n + 1
		} else if n == 1 {
			break
		}
		i++
	}
	return i - 1, nil
}
