func threeSum(nums []int)[][]int{
    	n := len(nums)
	var result [][]int
	
	sort.Ints(nums)
   
	// To help avoid duplicates
	seen := make(map[[3]int]bool)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
				    key := [3]int{nums[i],nums[j],nums[k]}
				    if _, ok := seen[key]; !ok{
				        seen[key] = true
				        result = append(result,[]int{nums[i],nums[j],nums[k]})
				    }
				}
			}
		}
	}
    
    return result
}
