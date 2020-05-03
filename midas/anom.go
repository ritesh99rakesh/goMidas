package midas

import (
	"math"
)

func MaxInts(arr []int) (int, int) {
	arrLen := len(arr)
	if arrLen == 0 {
		panic("cannot search in array of size 0")
	}
	maxValue, maxPos := arr[0], int(0)
	for currPos, currVal := range arr {
		if maxValue < currVal {
			maxValue, maxPos = currVal, currPos
		}
	}
	return maxValue, maxPos
}

func count_to_anom(total, current float64, curT int) float64 {
	curMean := total / float64(curT)
	sqErr := math.Pow(math.Max(0, current-curMean), 2)
	return sqErr/curMean + sqErr/(curMean*math.Max(1, float64(curT-1)))
}

func Midas(src, dst, times *[]int, numRows int, numBuckets int) *[]float64 {
	m, _ := MaxInts(*src)
	curCount, totalCount := Edgehash{}, Edgehash{}
	curCount.Edgehash(numRows, numBuckets, m)
	totalCount.Edgehash(numRows, numBuckets, m)
	anomalyScore := make([]float64, len(*src))
	curT, size := int(1), int(len(*src))
	var curSrc, curDst int
	var curMean, sqErr, curScore float64
	for i := int(0); i < size; i++ {
		if i == 0 || (*times)[i] > curT {
			curCount.clear()
			curT = (*times)[i]
		}
		curSrc = (*src)[i]
		curDst = (*dst)[i]
		curCount.insert(curSrc, curDst, 1)
		totalCount.insert(curSrc, curDst, 1)
		curMean = totalCount.getCount(curSrc, curDst) / float64(curT)
		sqErr = math.Pow(curCount.getCount(curSrc, curDst)-curMean, 2)
		if curT == 1 {
			curScore = 0
		} else {
			curScore = sqErr/curMean + sqErr/(curMean*float64(curT-1))
		}
		anomalyScore[i] = curScore
	}
	return &anomalyScore
}

func MidasR(src, dst, times *[]int, numRows, numBuckets int, factor float64) *[]float64 {
	m, _ := MaxInts(*src)
	curCount, totalCount := Edgehash{}, Edgehash{}
	curCount.Edgehash(numRows, numBuckets, m)
	totalCount.Edgehash(numRows, numBuckets, m)
	srcScore, dstScore, srcTotal, dstTotal := Nodehash{}, Nodehash{}, Nodehash{}, Nodehash{}
	srcScore.Nodehash(numRows, numBuckets)
	dstScore.Nodehash(numRows, numBuckets)
	srcTotal.Nodehash(numRows, numBuckets)
	dstTotal.Nodehash(numRows, numBuckets)
	anomalyScore := make([]float64, len(*src))
	curT, size := int(1), int(len(*src))
	var curSrc, curDst int
	var curScore, curScoreSrc, curScoreDst, combinedScore float64

	for i := int(0); i < size; i++ {

		if i == 0 || (*times)[i] > curT {
			curCount.lower(factor)
			srcScore.lower(factor)
			dstScore.lower(factor)
			curT = (*times)[i]
		}

		curSrc = (*src)[i]
		curDst = (*dst)[i]
		curCount.insert(curSrc, curDst, 1)
		totalCount.insert(curSrc, curDst, 1)
		srcScore.insert(curSrc, 1)
		dstScore.insert(curDst, 1)
		srcTotal.insert(curSrc, 1)
		dstTotal.insert(curDst, 1)
		curScore = count_to_anom(totalCount.getCount(curSrc, curDst), curCount.getCount(curSrc, curDst), curT)
		curScoreSrc = count_to_anom(srcTotal.getCount(curSrc), srcScore.getCount(curSrc), curT)
		curScoreDst = count_to_anom(dstTotal.getCount(curDst), dstScore.getCount(curDst), curT)

		combinedScore = math.Max(math.Max(curScoreSrc, curScoreDst), curScore)
		anomalyScore[i] = math.Log(1 + combinedScore)
	}
	return &anomalyScore
}
