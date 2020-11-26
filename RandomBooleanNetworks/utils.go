package rbn

import (
	"math/rand"
)

func bool2int(b bool) int {
	if b {
		return 1
	}

	return 0
}

func in(x int, list []int) bool {
	for _, el := range list {
		if el == x {
			return true
		}
	}

	return false
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func generateTruthTable(N int, randGen *rand.Rand) map[int]bool {
	tt := make(map[int]bool)

	for i := 0; i < N; i++ {
		tt[i] = randGen.Intn(100) > 50
	}

	return tt
}
