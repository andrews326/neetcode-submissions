func topKFrequent(nums []int, k int) []int {

	frequentMap := make(map[int]int)

	for _, num := range nums {
		frequentMap[num]++
	}

	keys := []int{}

	for key := range frequentMap {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return frequentMap[keys[i]] > frequentMap[keys[j]]
	})
	
    
	return keys[:k]
}