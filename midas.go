package midas

import "math"

func countsToAnom(total, current float64, curT int) float64 {
	curMean := total / float64(curT)
	sqErr := math.Pow(math.Max(0, current-curMean), 2)
	return sqErr/curMean + sqErr/(curMean*math.Max(1, float64(curT-1)))
}

func Midas(src, dst, times []int, numRows int, numBuckets int) []float64 {
	m, _ := MaxInts(src)
	curCount, totalCount := Edgehash{}, Edgehash{}
	curCount.Edgehash(numRows, numBuckets, m)
	totalCount.Edgehash(numRows, numBuckets, m)
	anomalyScore := make([]float64, len(src))
	curT, size := 1, len(src)
	var curSrc, curDst int
	var curMean, sqErr, curScore float64
	for i := 0; i < size; i++ {
		if i == 0 || times[i] > curT {
			curCount.clear()
			curT = times[i]
		}
		curSrc = src[i]
		curDst = dst[i]
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
	return anomalyScore
}
