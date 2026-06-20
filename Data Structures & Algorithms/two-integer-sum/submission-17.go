func twoSum(nums []int, target int) []int {
    intMap := make(map[int]int) // num -> index
    for i, num := range nums {
        complement := target - num
        if idx, ok := intMap[complement]; ok {
            return []int{idx, i}
        }
        intMap[num] = i // ✅ store number as key, index as value
    }
    return nil
}
