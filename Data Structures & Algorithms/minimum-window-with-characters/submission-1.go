func minWindow(s string, t string) string {
    charSet := make(map[rune]int)
   
   for _, v := range t {
       charSet[v]++
   }
   
   res := []int{-1,-1}
   
   min := math.MaxInt
   
   
   for i := 0; i < len(s); i++{
       cSet := make(map[rune]int)
       for j := i; j < len(s); j++{
           cSet[rune(s[j])]++
           flag := true
           for key, cnt := range charSet{
               if cnt > cSet[key] {
                   flag = false
                   break
               }
           }
           
           if flag &&  j-i+1 < min{
               res[0] = i
               res[1] = j
               
               min = j-i+1
           }
           
           
       }
   }
   
   if res[0] == -1{
       fmt.Println()
       return ""
   }
   
   return s[res[0]: res[1]+1]
}
