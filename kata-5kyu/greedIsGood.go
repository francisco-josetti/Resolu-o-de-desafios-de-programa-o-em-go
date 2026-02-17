package kata

func Score(dice [5]int) int {
  var d1,d2,d3,d4,d5,d6,points int
  for _,singleDice := range dice{
    switch singleDice{
      case 1:
      d1 += 1
      case 2:
      d2 += 1
      case 3:
      d3 += 1
      case 4:
      d4 += 1
      case 5:
      d5 += 1
      case 6:
      d6 += 1
    }
  }
  if d1 > 0 && d1 < 3{
    points += d1 * 100
  }else if d1 == 3{
    points += 1000
  }else if d1 > 3{
    points += 1000 + ((d1-3)*100)
  }
  
  if d5 > 0 && d5 < 3{
    points += d5 * 50
  }else if d5 == 3{
    points += 500
  }else if d5 > 3{
    points += 500 + ((d5-3)*50)
  }
  
  if d2 >= 3{
    points += 200
  }
  
  if d3 >= 3{
    points += 300
  }
  
  if d4 >= 3{
    points += 400
  }
  if d6 >= 3{
    points += 600
  }
  return points
}
