func checkInclusion(s1, s2 string) bool{
 s1Sorted := strSort(s1)
    
    for i := 0; i < len(s2); i++{
        for j := i; j < len(s2); j++{
            s2Sorted := strSort(s2[i:j+1])
            if s1Sorted == s2Sorted{
                return true
            }
        }
    }
    
    return false
}

// s1 = "abc"
// s2 = "lecabee"

func strSort(str string)string{
    arr := []rune(str)
    
    sort.Slice(arr, func(i,j int)bool{
        return arr[i] < arr[j]
    })
    
    return string(arr)
}