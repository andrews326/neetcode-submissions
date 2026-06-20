func checkInclusion(s1, s2 string) bool{
    s1sorted := []rune(s1)
    sort.Slice(s1sorted,func(i,j int)bool{
      return   s1sorted[i] < s1sorted[j]
    })
    s1str := string(s1sorted)
    for i := 0; i < len(s2); i++{
        for j := i; j < len(s2); j++{
            subStr :=  s2[i : j+1]
            subSorted := []rune(subStr)
            sort.Slice(subSorted,func(i , j int)bool{
                return subSorted[i] < subSorted[j]
            })
            
            if string(subSorted) == s1str{
                return true
            }
        }
    } 
    return false
}

// s1 = "abc"
// s2 = "lecabee"