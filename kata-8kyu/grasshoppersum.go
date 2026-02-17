package kata

func Summation(n int) int {
  var sum int
    for i := 1; i < n+1; i++{
      sum += i
    }
  return sum
}