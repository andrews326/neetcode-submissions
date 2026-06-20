type Solution struct{}

func (s *Solution) Encode(strs []string) string {
    var encoded string
    for _, str := range strs {
        encoded += fmt.Sprintf("%d#%s", len(str), str)
    }
    return encoded
}

func (s *Solution) Decode(str string) []string {
    var decoded []string
    i := 0
    for i < len(str) {
        // Find the '#' separator
        j := i
        for j < len(str) && str[j] != '#' {
            j++
        }

        // Parse the length
        lengthStr := str[i:j]
        length, _ := strconv.Atoi(lengthStr)

        // Move past '#'
        j++

        // Extract the string of that length
        decoded = append(decoded, str[j:j+length])

        // Move pointer
        i = j + length
    }
    return decoded
}
