func lengthOfLongestSubstring(s string) int {
    max := 0

    for i := 0; i < len(s); i++{
       sets := make(map[byte]bool)
       for j := i; j < len(s); j++{
         if sets[s[j]]{
            break
         }
         sets[s[j]] = true
       }

       if len(sets) > max {
        max = len(sets)
       }
    }

    return max
}

//pwwkew
