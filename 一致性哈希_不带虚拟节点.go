package main

import (
	"fmt"
	"hash/crc32"
	"sort"
	"strconv"
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

type Hash func(data []byte) uint32

type Map struct {
	hash    Hash
	keys    UInt32Slice       // 已排序的节点哈希切片
	hashMap map[uint32]string // 节点哈希和KEY的map，键是哈希值，值是节点Key
}

func New(fn Hash) *Map {
	m := &Map{
		hash:    fn,
		hashMap: make(map[uint32]string),
	}
	// 默认使用CRC32算法
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

func (m *Map) IsEmpty() bool {
	return len(m.keys) == 0
}

// Add 方法用来添加缓存节点，参数为节点key，比如使用IP
func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		// 结合复制因子计算所有虚拟节点的hash值，并存入m.keys中，同时在m.hashMap中保存哈希值和key的映射
		hash := m.hash([]byte(key))
		m.keys = append(m.keys, hash)
		m.hashMap[hash] = key

	}
	// 对所有虚拟节点的哈希值进行排序，方便之后进行二分查找
	sort.Sort(m.keys)
}

// Get 方法根据给定的对象获取最靠近它的那个节点key
func (m *Map) Get(key string) string {
	if m.IsEmpty() {
		return ""
	}

	hash := m.hash([]byte(key))

	// 通过二分查找获取最优节点，第一个节点hash值大于对象hash值的就是最优节点
	idx := sort.Search(len(m.keys), func(i int) bool { return m.keys[i] >= hash })

	// 如果查找结果大于节点哈希数组的最大索引，表示此时该对象哈希值位于最后一个节点之后，那么放入第一个节点中
	if idx == len(m.keys) {
		idx = 0
	}

	return m.hashMap[m.keys[idx]]
}

const (
	ITEMS = 10000000
	NODES = 100
)

func main() {

	hashMap := New(nil)

	nodes := make(UInt32Slice, NODES)
	for i := 0; i < NODES; i++ {
		hashMap.Add(strconv.Itoa(i))
	}

	for i := 0; i < ITEMS; i++ {
		key := hashMap.Get(strconv.Itoa(i))
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
