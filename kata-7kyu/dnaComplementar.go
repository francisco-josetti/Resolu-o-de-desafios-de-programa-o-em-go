package kata
import(
"strings")
func DNAStrand(dna string) string {
  letrasOpostas := []string{}
  for _,letra := range dna{
    switch letra{
      case 'A':
      letrasOpostas = append(letrasOpostas,"T")
      case 'T':
      letrasOpostas = append(letrasOpostas,"A")
      case 'G':
      letrasOpostas = append(letrasOpostas,"C")
      case 'C':
      letrasOpostas = append(letrasOpostas,"G")
    }
  }
  complementar := strings.Join(letrasOpostas,"")
  return complementar
}