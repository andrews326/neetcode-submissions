func isPalindrome(str string) bool {
	start := 0

	d := strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			return r // Keep the character if it's a letter
		}
		return -1 // Drop the character if it's not a letter
	}, str)

	s := strings.ToLower(d)
	fmt.Println(s)

	end := len(s) - 1

	for i := 0; i < len(s); i++ {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}

	return true

}