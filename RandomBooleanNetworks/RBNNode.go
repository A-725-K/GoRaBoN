package rbn

import (
	"math"
	"math/rand"
)

// RBNNode : data type that represents a node in the network
type RBNNode struct {
	id     int
	links  []int
	value  bool
	layers int
}

func (node *RBNNode) flip() {
	node.value = !node.value
}

func (node *RBNNode) getIntFromLinks(rbn RandomBooleanNetwork) int {
	result := 0

	for i, neighbor := range node.links {
		result += int(math.Pow(2, float64(rbn.layers-i-1))) * bool2int(rbn.network[neighbor].value)
	}

	return result
}

// NewRBNNode : constructor
func NewRBNNode(id, maxN, layers int, randGen *rand.Rand) RBNNode {
	links := make([]int, layers)
	for i := range links {
		links[i] = -1
	}

	for layer := 0; layer < layers; layer++ {
		neighbor := randGen.Intn(maxN)
		for neighbor == id || in(neighbor, links) {
			neighbor = randGen.Intn(maxN)
		}
		links[layer] = neighbor
	}

	return RBNNode{id, links, randGen.Intn(100) > 50, layers}
}
