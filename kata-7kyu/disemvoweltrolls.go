package kata
import(
"strings")
func Disemvowel(comment string) string {
  var newword string
for _,word := range comment{
  if strings.ToLower(string(word)) == "a" || strings.ToLower(string(word)) == "e" || strings.ToLower(string(word)) == "i" || strings.ToLower(string(word)) == "o" || strings.ToLower(string(word)) == "u"{
    newword += ""
  }else{
    newword += string(word)
  }
}
  return newword
}