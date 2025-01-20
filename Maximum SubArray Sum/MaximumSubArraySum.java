import java.util.*;
import java.lang.Math;

class Main {
    public static void main(String[] args) {
        var input = List.of(2, 3, -8, 7, -1, 2, 3);
        System.out.println(String.format("Input: %s", input));
        
        var res = input.get(0);
        var maxEnding = input.get(0);
        for (int i = 1; i < input.size(); i++) {
            
            maxEnding = Math.max(maxEnding + input.get(i), input.get(i));
            res = Math.max(res, maxEnding);
        }
        
        System.out.println(res);
    }
}
