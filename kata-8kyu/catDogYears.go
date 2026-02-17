package kata

func CalculateYears(years int) (result [3]int) {
  var catYears int
  var dogYears int
  for i := 0; i < years; i++{
    if i == 0{
      catYears += 15
      dogYears += 15
    }else if i == 1{
      catYears += 9
      dogYears += 9
    }else{
      catYears += 4
      dogYears += 5
    }
  }
  result = [3]int{years,catYears,dogYears}
  return result
}