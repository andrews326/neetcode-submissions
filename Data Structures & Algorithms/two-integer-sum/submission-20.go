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

//4,5,6 
// 10

// compliment is 6 and if block will not call  and intMap[4]=0
// compliment is 5 and if block will not call  and intMap[5]=1
// compliment is 4 and if block will call and founf 4 and return the 0 and 2 

