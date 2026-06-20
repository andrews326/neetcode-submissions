func productExceptSelf(nums []int) []int {
    final := []int{}
    for i := 0; i < len(nums); i++ {
        prod := 1
        for j := 0; j < len(nums); j++ {
            if i == j {
                continue
            }
            prod *= nums[j]   // ✅ multiply other elements
        }
        final = append(final, prod)
    }
    return final

}
