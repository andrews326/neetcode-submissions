func maxProfit(prices []int) int {

  maxPrice := 0
  left, right := 0, 1
  
  for  right < len(prices) {
     if prices[left] < prices[right]{
        profit := prices[right] - prices[left]
        if profit > maxPrice{
         maxPrice = profit
        }
     }else{
      left = right
     }
     right++
  }

  return maxPrice

}
