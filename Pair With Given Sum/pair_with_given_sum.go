package main
import "fmt"

func main() {
  input := []int32{0, -1, 2, -3, 1}
  var target int32 = -2
  fmt.Printf("%s: %t\n", "Assert True:", checkTarget(input, target))
  
  input2 := []int32{1, -2, 1, 0, 5}
  var target2 int32 = 0
  fmt.Printf("%s: %t", "Assert False:", checkTarget(input2, target2))
}

func checkTarget(input []int32, target int32) bool {
    checkMap := make(map[int32]bool)
    for _, in := range input {
        if _, ok := checkMap[in]; ok {
            return true
        } else {
            checkMap[target - in] = true
        }
    }
    
    return false
}
