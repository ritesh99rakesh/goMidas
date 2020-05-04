package midas

func MaxInts(arr []int) (int, int) {
	arrLen := len(arr)
	if arrLen == 0 {
		panic("cannot search in array of size 0")
	}
	maxValue, maxPos := arr[0], 0
	for currPos, currVal := range arr {
		if maxValue < currVal {
			maxValue, maxPos = currVal, currPos
		}
	}
	return maxValue, maxPos
}
