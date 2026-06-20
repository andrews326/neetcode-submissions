func lengthOfLongestSubstring(s string) int {
  left, res := 0, 0
  charSet := make(map[byte]bool)

  for right := 0; right < len(s); right++{
    for charSet[s[right]]{
      delete(charSet,s[left])
      left++
    }

    charSet[s[right]] = true
    if right - left + 1 > res{
      res = right - left + 1
    }
  }

  return res
}

//pwwkew
