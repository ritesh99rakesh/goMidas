package midas

import (
	"math"
	"math/rand"
	"time"
)

type Edgehash struct {
	numRows, numBuckets, m int
	hashA, hashB           []int
	count                  [][]float64
}

func (e Edgehash) Edgehash(r, b, m0 int) {
	e.numRows, e.numBuckets, e.m = r, b, m0
	e.hashA, e.hashB = make([]int, r), make([]int, r)
	randomGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < r; i++ {
		e.hashA[i] = randomGenerator.Intn(b-1) + 1
		e.hashB[i] = randomGenerator.Intn(b)
	}
	e.clear()
}

func (e Edgehash) hash(a, b, i int) int {
	resid := ((a+e.m*b)*e.hashA[i] + e.hashB[i]) % e.numBuckets
	if resid < 0 {
		return resid + e.numBuckets
	} else {
		return resid
	}
}

func (e Edgehash) insert(a, b int, weight float64) {
	var bucket int
	for i := 0; i < e.numRows; i++ {
		bucket = e.hash(a, b, i)
		e.count[i][bucket] += weight
	}
}

func (e Edgehash) getCount(a, b int) float64 {
	minCount := math.MaxFloat64
	var bucket int
	for i := 0; i < e.numRows; i++ {
		bucket = e.hash(a, b, i)
		minCount = math.Min(minCount, e.count[i][bucket])
	}
	return minCount
}

func (e Edgehash) clear() {
	e.count = make([][]float64, e.numRows)
	for i := range e.count {
		e.count[i] = make([]float64, e.numBuckets)
	}
}

func (e Edgehash) lower(factor float64) {
	for i := 0; i < e.numRows; i++ {
		for j := 0; j < e.numBuckets; j++ {
			e.count[i][j] = e.count[i][j] * factor
		}
	}
}
