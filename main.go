package main

import (
	"fmt"
	"math"
	"math/rand"
	rbn "rbn/RandomBooleanNetworks"
	"time"
)

// some useful constants
const (
	EPOCHS = 50   //how many iterations in the generation
	LINKS  = 12   //how many neighbors has each node in the DiGraph
	NODES  = 4096 //how many nodes in the network
)

func isPowerOf2(n int) bool {
	if n == 0 {
		return false
	}

	return math.Ceil(math.Log2(float64(n))) == math.Floor(math.Log2(float64(n)))
}

func checkParameters(LINKS, NODES, EPOCHS int) (bool, string) {
	if EPOCHS < 2 {
		return false, "[Error]: Number of epochs must be greater than 1"
	}

	if NODES < 3 {
		return false, "[Error]: Number of nodes must be greater than 2"
	}

	if LINKS < 1 {
		return false, "[Error]: Number of links must be greater than 0"
	}

	n := int(math.Sqrt(float64(NODES)))
	if n*n != NODES || !isPowerOf2(NODES) {
		return false, "[Error]: Number of nodes must be a perfect square and a power of 2"
	}

	if int(math.Pow(2.0, float64(LINKS))) != NODES {
		return false, "[Error]: 2^(lumber of links) must be equal to the number of nodes"
	}

	return true, ""
}

func main() {
	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))

	if t, err := checkParameters(LINKS, NODES, EPOCHS); !t {
		fmt.Println(err)
		return
	}

	rbn := rbn.NewRBN(NODES, LINKS, randGen)
	//rbn.PrettyPrintNetwork() //print initial status
	rbn.Evolve(EPOCHS)
}
