package feature

func min(x, y, z int) int {
	if x <= y && x <= z {
		return x
	}

	if y <= x && y <= z {
		return y
	}

	return z
}

func BooleanData(except, input string) bool {
	lengthExcept := len(except)
	lengthInput := len(input)

	distance := make([][]int, lengthExcept+1)

	for i := 0; i <= lengthExcept; i++ {
		distance[i] = make([]int, lengthInput+1)
	}

	for i := 0; i <= lengthExcept; i++ {
		distance[i][0] = i
	}

	for j := 0; j <= lengthInput; j++ {
		distance[0][j] = j
	}

	for i := 1; i <= lengthExcept; i++ {
		for j := 1; j <= lengthInput; j++ {
			if except[i-1] == input[j-1] {
				distance[i][j] = distance[i-1][j-1]
			} else {
				distance[i][j] = min(distance[i-1][j], distance[i][j-1], distance[i-1][j-1]) + 1
			}
		}
	}

	return distance[lengthExcept][lengthInput] == 1
}
