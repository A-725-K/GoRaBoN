package rbn

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
)

// RandomBooleanNetwork : representation of Random Boolean Network
type RandomBooleanNetwork struct {
	network    []RBNNode
	truthTable map[int]bool
	size       int
	layers     int
}

// Evolve : evolution the network in time, flipping bytes
func (rbn *RandomBooleanNetwork) Evolve(EPOCHS int) {
	for epoch := 0; epoch < EPOCHS; epoch++ {
		fmt.Println("--- Epoch:", epoch, "---")
		for _, node := range rbn.network {
			idx := node.getIntFromLinks(*rbn)
			rbn.network[idx].flip()
			//node.value = rbn.truthTable[idx]
		}
		rbn.PrintStateCsv(epoch)
	}
}

// NewRBN : constructor
func NewRBN(numNodes, numLinks int, randGen *rand.Rand) RandomBooleanNetwork {
	truthTable := generateTruthTable(numNodes, randGen)
	network := make([]RBNNode, numNodes)

	for i := 0; i < numNodes; i++ {
		network[i] = NewRBNNode(i, numNodes, numLinks, randGen)
	}

	return RandomBooleanNetwork{network, truthTable, numNodes, numLinks}
}

// PrintStateCsv : snapshot the network status on a file
func (rbn *RandomBooleanNetwork) PrintStateCsv(epoch int) {
	filename := "res/rbn_" + strconv.Itoa(epoch) + ".csv"

	f, err := os.Create(filename)
	checkError(err)

	outAsCsv := ""
	perLine := int(math.Sqrt(float64(rbn.size)))
	for i, node := range rbn.network {
		outAsCsv += strconv.Itoa(bool2int(node.value)) + ","
		if (i+1)%perLine == 0 {
			outAsCsv += "\n"
		}
	}

	f.WriteString(outAsCsv)
}

// PrettyPrintNetwork : print the network status on stdout formatted
func (rbn *RandomBooleanNetwork) PrettyPrintNetwork() {
	fmt.Println("Network: [Node] -> n_1{v1}, n_2{v2}, ..., n_n{vn} | as_int")
	for _, node := range rbn.network {
		fmt.Printf("\t[%4d] -> ", node.id)
		for i, neigh := range node.links {
			if i != rbn.layers-1 {
				fmt.Printf("%4d{%d}, ", neigh, bool2int(rbn.network[neigh].value))
			} else {
				fmt.Printf("%4d{%d}\t", neigh, bool2int(rbn.network[neigh].value))
			}

		}
		fmt.Printf("| %d\n", node.getIntFromLinks(*rbn))
	}
}
