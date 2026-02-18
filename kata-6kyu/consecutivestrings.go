package kata

func LongestConsec(strarr []string, k int) string {
  newstrarr := []string{}
  var max string
  var word string
    for i := 0; i < len(strarr)-k+1; i++{
      word = ""
      for j := 0; j< k; j++{
        if j+i < len(strarr){
        word += strarr[j+i]
          }
      }
      newstrarr = append(newstrarr,word)
    }
  for _,word := range newstrarr{
    if len(word) > len(max){
      max = word
    }
  }
  return max
}