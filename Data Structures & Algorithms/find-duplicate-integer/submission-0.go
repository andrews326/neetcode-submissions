func findDuplicate(nums []int) int {
    hashMap := make(map[int]int)

	for _, num := range nums{
		if _, ok := hashMap[num]; ok{
			 return  num
		}

		hashMap[num] = num
 	}

	return -1
}
