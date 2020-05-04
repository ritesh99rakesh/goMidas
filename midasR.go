package midas

import "math"

func MidasR(src, dst, times []int, numRows, numBuckets int, factor float64) []float64 {
	m, _ := MaxInts(src)
	curCount, totalCount := Edgehash{}, Edgehash{}
	curCount.Edgehash(numRows, numBuckets, m)
	totalCount.Edgehash(numRows, numBuckets, m)
	srcScore, dstScore, srcTotal, dstTotal := Nodehash{}, Nodehash{}, Nodehash{}, Nodehash{}
	srcScore.Nodehash(numRows, numBuckets)
	dstScore.Nodehash(numRows, numBuckets)
	srcTotal.Nodehash(numRows, numBuckets)
	dstTotal.Nodehash(numRows, numBuckets)
	anomalyScore := make([]float64, len(src))
	curT, size := 1, len(src)
	var curSrc, curDst int
	var curScore, curScoreSrc, curScoreDst, combinedScore float64

	for i := 0; i < size; i++ {

		if i == 0 || times[i] > curT {
			curCount.lower(factor)
			srcScore.lower(factor)
			dstScore.lower(factor)
			curT = times[i]
		}

		curSrc = src[i]
		curDst = dst[i]
		curCount.insert(curSrc, curDst, 1)
		totalCount.insert(curSrc, curDst, 1)
		srcScore.insert(curSrc, 1)
		dstScore.insert(curDst, 1)
		srcTotal.insert(curSrc, 1)
		dstTotal.insert(curDst, 1)
		curScore = countsToAnom(totalCount.getCount(curSrc, curDst), curCount.getCount(curSrc, curDst), curT)
		curScoreSrc = countsToAnom(srcTotal.getCount(curSrc), srcScore.getCount(curSrc), curT)
		curScoreDst = countsToAnom(dstTotal.getCount(curDst), dstScore.getCount(curDst), curT)

		combinedScore = math.Max(math.Max(curScoreSrc, curScoreDst), curScore)
		anomalyScore[i] = math.Log(1 + combinedScore)
	}
	return anomalyScore
}
