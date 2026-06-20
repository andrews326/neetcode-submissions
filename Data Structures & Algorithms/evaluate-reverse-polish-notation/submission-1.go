func evalRPN(tokens []string) int {
    stack := []int{}
   for _, token := range tokens{
      var res int
      switch token{
        case "+", "-", "*", "/":
        b := stack[len(stack)-1]
        a := stack[len(stack)-2]
        stack = stack[:len(stack)-2]
        if token == "+"{
         res =  a + b
         stack = append(stack, res)  
        }else if token == "-"{
         res =  a - b
         stack = append(stack, res)  
        }else if token == "*"{
         res =  a * b
         stack = append(stack, res)  
        }else if token == "/"{
          res =  a / b
         stack = append(stack, res)  
        }

        default:
         num, _ := strconv.Atoi(token)
         stack = append(stack, num)
      }

   }

   return stack[0]
}
