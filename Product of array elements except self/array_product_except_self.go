package main
import "fmt"

func main() {
  input := []int{1, 2, 3, 4, 5, 0}
  var totalProduct int = 1
  zeroCounter := 0
  firstZeroIndex := 0
  for i, in := range input {
      if in != 0 {
          totalProduct *= in
      }else {
          zeroCounter++
          firstZeroIndex = i
      }
  }

  output := make([]int, len(input))
  if zeroCounter == 1 {
      for i, _ := range input {
          if firstZeroIndex == i {
              output[i] = totalProduct
          }else {
              output[i] = 0
          }
      }
  } else if zeroCounter > 1 {
      for i, _ := range input {
          output[i] = 0
      }
  }else {
      for i, in := range input {
      if in == 0 {
          output[i] = 0
      }else {
          output[i] = totalProduct/in
      }
    }
  }   
  
  fmt.Printf("Result: %v", output)
}
