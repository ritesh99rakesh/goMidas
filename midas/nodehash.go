package midas

import (
	"math"
	"math/rand"
	"time"
)

type Nodehash struct {
	numRows, numBuckets, m int
	hashA, hashB           []int
	count                  [][]float64
}

func (n Nodehash) Nodehash(r, b int) {
	n.numRows, n.numBuckets = r, b
	n.hashA, n.hashB = make([]int, r), make([]int, r)
	randomGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < r; i++ {
		n.hashA[i] = randomGenerator.Intn(n.numBuckets-1) + 1
		n.hashB[i] = randomGenerator.Intn(n.numBuckets)
	}
	n.clear()
}

func (n Nodehash) hash(a, i int) int {
	resid := (a*n.hashA[i] + n.hashB[i]) % n.numBuckets
	if resid < 0 {
		return resid + n.numBuckets
	} else {
		return resid
	}
}

func (n Nodehash) insert(a int, weight float64) {
	var bucket int
	for i := 0; i < n.numRows; i++ {
		bucket = n.hash(a, i)
		n.count[i][bucket] += weight
	}
}

func (n Nodehash) getCount(a int) float64 {
	minCount := math.MaxFloat64
	var bucket int
	for i := 0; i < n.numRows; i++ {
		bucket = n.hash(a, i)
		minCount = math.Min(minCount, n.count[i][bucket])
	}
	return minCount
}

func (n Nodehash) clear() {
	n.count = make([][]float64, n.numRows)
	for i := range n.count {
		n.count[i] = make([]float64, n.numBuckets)
	}
}

func (n Nodehash) lower(factor float64) {
	for i := 0; i < n.numRows; i++ {
		for j := 0; j < n.numBuckets; j++ {
			n.count[i][j] = n.count[i][j] * factor
		}
	}
}
