func isValid(s string) bool {
    stack := map[rune]rune{
        ')':'(',
        '}':'{',
        ']': '[',
    }

   slice := []rune{}
    for _, char := range s {
        if char == '(' || char == '[' || char == '{'{
            slice = append(slice, char)
        }else if char == ')' || char == ']' || char == '}'{
            if len(slice) == 0 || slice[len(slice)-1] != stack[char]{
                return false
            }

           slice = slice[:len(slice)-1]
        }
    }

    return len(slice) == 0
    
}
