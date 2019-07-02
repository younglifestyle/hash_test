package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
)

const (
	ITEMS     = 10000000
	NODES     = 100
	NEW_NODES = 98
)

type byInt []int

func (s byInt) Len() int {
	return len(s)
}
func (s byInt) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
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
	change := 0
	for i := 0; i < ITEMS; i++ {
		md5Byte := MD5(strconv.Itoa(i))
		value := genValue(md5Byte[6:10])
		// 原映射结果
		n := value % NODES

		//现映射结果
		nNew := value % NEW_NODES
		if n != nNew {
			change += 1
		}
	}

	fmt.Println(float64(change) / float64(ITEMS))
}

func genValue(bs []byte) uint32 {
	if len(bs) < 4 {
		return 0
	}
	v := (uint32(bs[3]) << 24) | (uint32(bs[2]) << 16) | (uint32(bs[1]) << 8) | (uint32(bs[0]))
	return v
}
