package main
import (
	"fmt"
)
func concatArraySlice(fa1 [3]float64, fa2 [3]float64) []float64 {
  return (append(fa1[0:], fa2[0:]...))
}

func concatArrayArray(fa1 [3]float64, fa2 [3]float64) [6]float64 {
  result := [6]float64{}
  for i := 0; i < len(fa1); i++ {
    result[i] = fa1[i]
  }
  for i := len(fa1); i < len(fa1)+len(fa2); i++ {
    result[i] = fa2[i-len(fa1)]
  }
  return (result)
}

func concatSliceArray(fs1 []float64, fs2 []float64) []float64 {
  return (append(fs1, fs2...))
}

func main() {
  a1 := [3]float64{1,2,3}
  a2 := [3]float64{1.1, 2.2, 3.3}
  s1 := []float64{4,5,6}
  s2 := []float64{4.4, 5.5, 6.6}
  as1 := concatArraySlice(a1, a2)
  aa1 := concatArrayArray(a1, a2)
  sa1 := concatSliceArray(s1, s2)
  fmt.Println(a1)
  fmt.Println(a2)
  fmt.Println(as1)
  fmt.Println(aa1)
  fmt.Println(sa1)
}