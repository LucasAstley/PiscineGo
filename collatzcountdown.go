package piscine

func CollatzCountdown(start int) int {
	if start <= 0 {
		return -1
	}
	steps := 0
	for i := 0; start != 1; i++ {
		if start%2 == 0 {
			start /= 2
		} else {
			start = start*3 + 1
		}
		steps += 1
	}
	return steps
}
