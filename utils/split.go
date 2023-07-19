package utils

func SplitTargetList(targetList []string, blockNum int) [][]string {
	max := len(targetList)

	if max < blockNum {
		blockNum = max
	}

	var segmens = make([][]string, 0)

	quantity := max / blockNum

	end := 0

	for i := 1; i <= blockNum; i++ {
		qu := i * quantity

		if i != blockNum {
			segmens = append(segmens, targetList[i-1+end:qu])
		} else {
			segmens = append(segmens, targetList[i-1+end:qu])
			if len(targetList[qu-1:]) != 1 {
				for i := 0; i < len(targetList[qu:]); i++ {
					segmens[i] = append(segmens[i], targetList[qu+i])
				}
			}
		}

		end = qu - i
	}
	return segmens
}
