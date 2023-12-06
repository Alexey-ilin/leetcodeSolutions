package solutions

func IsSelfCrossing(distance []int) bool {
	if len(distance) < 4 {
		return false
	}
	isOut, endIdx := isOutSpiral(distance)
	if isOut {
		return false
	}
	isIn, _ := isInSpiral(distance, endIdx)
	return !isIn
}

func isOutSpiral(distance []int) (bool, int) {
	for sIdx := 2; sIdx < len(distance); sIdx++ {
		if distance[sIdx] <= distance[sIdx-2] {
			return false, sIdx
		}
	}
	return true, -1

}

func isInSpiral(distance []int, fromIdx int) (bool, int) {
	if fromIdx >= 3 && distance[fromIdx] == distance[fromIdx-2] && distance[fromIdx+1] >= distance[fromIdx-1]-distance[fromIdx-3] {
		return false, fromIdx
	}
	for sIdx := fromIdx + 1; sIdx < len(distance); sIdx++ {
		if distance[sIdx] >= distance[sIdx-2] {
			return false, sIdx
		}
	}
	return true, -1
}
