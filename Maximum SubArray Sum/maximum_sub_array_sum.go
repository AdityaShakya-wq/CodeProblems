package main
import "fmt"

func main() {
  input1 := []int{2, 3, -8, 7, -1, 2, 3}
  fmt.Printf("Input1: %d\n", input1)
  output1 := getMaximumSubArraySum(input1)
  fmt.Printf("Output1: %d\n", output1)
  
  input2 := []int{-2, -4}
  fmt.Printf("Input2: %d\n", input2)
  output2 := getMaximumSubArraySum(input2)
  fmt.Printf("Output2: %d\n", output2)
  
  input3 := []int{5, 4, 1, 7, 8}
  fmt.Printf("Input3: %d\n", input3)
  output3 := getMaximumSubArraySum(input3)
  fmt.Printf("Output3: %d\n", output3)
}

func getMaximumSubArraySum(input []int) int {
    if len(input) > 0 {
        sum := input[0]
        finalSum := sum
        sumArray := make([]int, 0)
        sumArray = append(sumArray, input[0])
        for i := 1; i < len(input); i++ {
            temp := sum + input[i]
            
            if temp < 0 && temp < sum {
                sum = 0
                sumArray = make([]int, 0)
                continue
            }
            
            sum = temp
            sumArray = append(sumArray, input[i])
            if sum > finalSum {
                finalSum = sum
            }
        }
        
        fmt.Printf("Sum Array: %d\n", sumArray)
        return finalSum
    }
    
    return 0
}
