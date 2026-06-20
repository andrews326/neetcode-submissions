func maxSlidingWindow(nums []int, k int) []int {
    final := []int{}
    
    for i := 0; i <= len(nums)-k; i++{
        max := nums[i]
        for j := i; j < i+k; j++{
            if nums[j] > max{
                max = nums[j]
            }
        }
        
        final = append(final,max)
    }
    
    return final
} 
