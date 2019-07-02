package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"sort"
	"strconv"
)

const (
	ITEMS = 10000000
	NODES = 100
)

type byInt []int

func (s byInt) Len() int {
	return len(s)
}
func (s byInt) Swap(i, j int) {
	s[j], s[i] = s[i], s[j]
}
func (s byInt) Less(i, j int) bool {
	return s[i] < s[j]
}

func MD5(str string) []byte {
	hasher := md5.New()
	io.WriteString(hasher, str)

	return hasher.Sum(nil) /*hex.EncodeToString(hasher.Sum(nil))*/
}

func main() {
	// 创建节点数据
	nodeStat := make(byInt, NODES)

	for i := 0; i < ITEMS; i++ {
		md5Byte := MD5(strconv.Itoa(i))
		value := genValue(md5Byte[6:10])
		n := value % NODES
		nodeStat[n] += 1
	}

	sort.Sort(nodeStat)

	ave := ITEMS / NODES
	Max := nodeStat[len(nodeStat)-1]
	Min := nodeStat[0]

	fmt.Printf("Ave: %d \n", ave)

	fmt.Printf("Max: %d\t(%0.2f%%) \n", Max,
		float64(Max-ave)*100.0/float64(ave))

	fmt.Printf("Min: %d\t(%0.2f%%) \n", Min,
		float64(ave-Min)*100.0/float64(ave))

	//ave := ITEMS / NODES
	//min := float64(nodeStat[0]) / float64(ITEMS)
	//max := float64(nodeStat[len(nodeStat)-1]) / float64(ITEMS)
	//
	//fmt.Println(ave, max, min)
}

func genValue(bs []byte) uint32 {
	if len(bs) < 4 {
		return 0
	}
	v := (uint32(bs[3]) << 24) | (uint32(bs[2]) << 16) | (uint32(bs[1]) << 8) | (uint32(bs[0]))
	return v
}
