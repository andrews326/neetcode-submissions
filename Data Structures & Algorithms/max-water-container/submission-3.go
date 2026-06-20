func maxArea(heights []int) int {
   left, right := 0, len(heights)-1
   maxArea := 0
   for left < right{
       h := min(heights[left],heights[right]) 

       width := right - left
       area := h * width
       
       if area > maxArea{
           maxArea = area
       }
       
       if heights[left] < heights[right]{
           left++
       }else{
           right--
       }
   }
   
   return maxArea
   
   
}

func min(i , j int)int{
    if i < j{
        return i
    }else {
       return j 
    } 
}
