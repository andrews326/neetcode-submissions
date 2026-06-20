func carFleet(target int, position []int, speed []int) int {
     n := len(position)

	 pairs := make([][2]int, n)

	 for i := 0; i < n; i++{
		pairs[i] = [2]int{position[i],speed[i]}
	 }

	 sort.Slice(pairs, func(i, j int)bool{
		return pairs[i][0] > pairs[j][0]
	 })

	 stack := []float64{}

	 for _, p := range pairs{
		speed := float64(target-p[0])/ float64(p[1])
		stack = append(stack, speed)
		if len(stack) >= 2 && stack[len(stack)-1] <= stack[len(stack)-2]{
			stack = stack[:len(stack)-1]
		}
		fmt.Println(stack)
	 }

	 return len(stack)

}
