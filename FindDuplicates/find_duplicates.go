package main
import "fmt"
import "sort"

func main() {
  input := []int{1, 1, 2, 3, 6, 3, 3, 6, 1}
  sort.Slice(input, func(i, j int) bool {
      return input[i] < input[j]
  })
  
  fmt.Printf("Duplicate entries: %v", getDuplicates(input))
}

func getDuplicates(input []int) []int {
    var output []int
    for i, in := range input {
        if i == len(input) - 1 {
            break    
        }
        
        if in == input[i + 1] {
            if i == 0 || in != input[i - 1]{
                output = append(output, in)       
            }
        }
    }
    
    return output
}
