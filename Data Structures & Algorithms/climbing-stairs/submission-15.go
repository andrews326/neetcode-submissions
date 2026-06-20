func climbStairs(n int) int {

  one, two := 1,1

  for i := n-1; i > 0; i--{
    temp := one
    one = one + two
    two = temp

  }

  return one

}
