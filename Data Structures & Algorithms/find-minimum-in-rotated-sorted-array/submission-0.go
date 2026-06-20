func findMin(nums []int) int {
    min := math.MaxInt64

    for _, num := range nums{
        if num < min{
            min = num
        }
    }

    return min

}
