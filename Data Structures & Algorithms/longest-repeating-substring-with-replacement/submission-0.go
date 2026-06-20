func characterReplacement(s string, k int)int{
    res := 0 
    for i := 0; i < len(s); i++{
        charMap := make(map[byte]int)
        maxF := 0
        for j := i; j < len(s); j++{
            charMap[s[j]]++
            maxF = max(maxF, charMap[s[j]])
            if (j-i+1) - maxF <= k{
                res = max(res, j-i+1)
            }
        } 
    } 
    
   return res 
    
    
}