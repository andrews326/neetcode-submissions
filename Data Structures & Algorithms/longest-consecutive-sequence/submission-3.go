func longestConsecutive(nums []int) int {
   numSet := make(map[int]bool)
    for _, num := range nums{
        numSet[num] = true
    }

    longest := 0

    for _, num := range nums{
        if !numSet[num-1]{
            length := 1
            current := num

            for numSet[current+1]{
                current++
                length++
            }

            if length > longest{
                longest = length
            }
        }
    }

   return longest

}
