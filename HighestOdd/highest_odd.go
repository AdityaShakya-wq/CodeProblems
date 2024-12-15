// Program to find the highest odd
// input -> "1321125126932222222" Output -> "132112512693"
package main
import "fmt"
import "unicode"

func main() {
  var str string
  fmt.Println("Enter a number string:")
  fmt.Scan(&str)
  if !checkIfNumber(str) {
      fmt.Println("Not a number")
      return 
  }
  
  fmt.Println("Final output: " + getHighestOdd(str))
}

func checkIfNumber(str string) bool {
    for _, ch := range str {
        if !unicode.IsDigit(ch) {
            return false
        }
    }
    
    return true
}

func getHighestOdd(str string) string {
    var finalIndex int
    for i := len(str) -1; i >= 0; i -- {
        if str[i] % 2 != 0 {
            finalIndex = i
            break
        }
        
        if i == 0 && str[i] % 2 == 0 {
            return ""
        }
    }
    
    
    return str[:finalIndex + 1]
}
