package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/toolkits/consistent"
)

type UInt32Slice []uint32

func (s UInt32Slice) Len() int {
	return len(s)
}

func (s UInt32Slice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s UInt32Slice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

const (
	ITEMS     = 10000000
	NODES     = 100
	NODES_NEW = 110
)

func main() {

	c := consistent.New()

	nodes := make(UInt32Slice, NODES)

	for i := 0; i < NODES; i++ {
		c.Add(strconv.Itoa(i))
	}

	for i := 0; i < ITEMS; i++ {
		key, _ := c.Get(strconv.Itoa(i))
		inter, _ := strconv.Atoi(key)
		nodes[inter] += 1
	}
	sort.Sort(nodes)

	ave := uint32(ITEMS / NODES)
	Max := nodes[len(nodes)-1]
	Min := nodes[0]

	fmt.Printf("Ave: %d \n", ave)

	fmt.Printf("Max: %d\t(%0.2f%%) \n", Max,
		float64(Max-ave)*100.0/float64(ave))

	fmt.Printf("Min: %d\t(%0.2f%%) \n", Min,
		float64(ave-Min)*100.0/float64(ave))
}
