type Solution struct{}

func (s *Solution) Encode(strs []string) string {
    var encoded string
    for _, str := range strs {
        encoded += fmt.Sprintf("%d#%s", len(str), str)
    }
    return encoded
}

func (s *Solution) Decode(str string) []string {
    var result []string
    for len(str) > 0 {
        // Find position of first '#'
        delimPos := strings.Index(str, "#")
        if delimPos == -1 {
            break // malformed string
        }

        // Extract length prefix and convert to int
        lengthStr := str[:delimPos]
        length, err := strconv.Atoi(lengthStr)
        if err != nil {
            break // invalid length
        }

        // Extract the actual string
        start := delimPos + 1
        end := start + length
        if end > len(str) {
            break // malformed string
        }

        result = append(result, str[start:end])

        // Move to next segment
        str = str[end:]
    }

    return result
}

