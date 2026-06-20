func merge(intervals [][]int) [][]int {

    sort.Slice(intervals, func(i, j int)bool{
        return intervals[i][0] < intervals[j][0]
    })

    output := [][]int{intervals[0]}
	
	for _, interval := range intervals[1:]{
	    start, end := interval[0], interval[1]
	    lastIndex :=  output[len(output)-1][1]
	    
	    if start <= lastIndex{
	        output[len(output)-1][1] = max(end, lastIndex)
	    }else{
	        output = append(output, interval)
	    }
	}

    return output
    
}
