func maxProfit(prices []int) int {

      maxPrice := 0
  
  
  for i := 0; i < len(prices); i++{
      for j := i+1; j < len(prices); j++{
          if prices[i] > prices[j]{
              j++
          }else if prices[j] - prices[i] > maxPrice{
              maxPrice = prices[j] - prices[i]
          }
      }
  }

  return maxPrice

}
