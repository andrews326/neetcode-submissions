func isPalindrome(s string) bool {
  left , right := 0, len(s)-1
  
  for left <= right {
      for left < right && !checkUnicode(rune(s[left])){
          left++
      }
      
      for right > left && !checkUnicode(rune(s[right])){
          right--
      }
      
      if unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])){
          return false
      }

	  left++
	  right--
  }
  
  return true

}

func checkUnicode(s rune) bool{
    return unicode.IsNumber(s) || unicode.IsLetter(s)
}