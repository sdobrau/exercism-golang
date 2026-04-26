package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	steps := 0
	for progress := n; progress != 1; steps++ {
		if progress <= 0 {
			return 0, errors.New("0 or smaller is not a valid value")
		}
		if progress % 2 == 0 {
			progress = progress / 2
		} else if progress % 2 != 0 {
			progress *= 3
			progress += 1
		}
	}
	return steps, nil
}
