package kata
import(
"strings")
func Accum(s string) string {
  var accum string
    for i,letra := range s{
      if i < len(s)-1{
       accum += strings.ToUpper(string(letra)) + strings.Repeat(strings.ToLower(string(letra)),i)+"-"
        }else{
        accum += strings.ToUpper(string(letra)) + strings.Repeat(strings.ToLower(string(letra)),i)
        }
      }
  return accum
    }