func twoSum(numbers []int, target int) []int {
    final := []int{}
    low, high := 0, len(numbers)-1
    for low < high {
         sum := numbers[low] + numbers[high]
         if sum == target{
            final = append(final,low+1,high+1)
            return final
         }
         if sum > target{
            high--
         }else{
            low++
         }
    }

    return nil
}
