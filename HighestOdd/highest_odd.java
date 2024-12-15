// Program to find the highest odd
// input -> "1321125126932222222" Output -> "132112512693"

import java.util.Scanner;
import java.util.Arrays;
class Main {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        System.out.println("Enter a number?");
        String str = sc.nextLine();
        if (checkIfNumber(str)) {
            System.out.println("Highest odd: " + getHighestOdd(str));
        }
    }
    
    public static boolean checkIfNumber(String str) {
        var charArray = str.toCharArray();
        for(var ch: charArray) {
           if (!Character.isDigit(ch)) {
               return false;
           } 
        }
        
        return true;
    }
    
    public static String getHighestOdd(String str) {
        var finalIndex = 0;
        for(int i = str.length() - 1; i >= 0; i--) {
            var ch = str.charAt(i);
            var isEven = false;
            var chL = Long.valueOf(ch);
            if (chL % 2 != 0) {
                finalIndex = i;
                break;
            }else {
                isEven = true;
            }
            
            if (i == 0 && isEven) {
                return "";
            }
        }
        
        return str.substring(0, finalIndex + 1);
    }
}
