func dailyTemperatures(temperatures []int) []int {
    res := make([]int, len(temperatures))
    stack := []int{}

    for i := 0; i < len(temperatures); i++{
        for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]]{
            stackIn := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            res[stackIn] = i - stackIn
        }

        stack = append(stack, i)
    }

    return res
}
