import java.util.*;

class Main {
    public static void main(String[] args) {
       var input = new int[] {10, 3, 5, 6, 2};
       var output = new int[input.length];
       var zeroCounter = 0;
       var firstZeroIndex = 0;
       var totalProduct = 1;
       for(var i = 0; i < input.length; i++) {
           if (input[i] != 0) {
               totalProduct *= input[i];
           }else {
               zeroCounter++;
               firstZeroIndex = i;
           }
       }
       
       if(zeroCounter == 1) {
           for (var i = 0; i < input.length; i++) {
               if (i == firstZeroIndex) {
                   output[i] = totalProduct;
               }else {
                   output[i] = 0;
               }
           }
       } else if (zeroCounter > 1) {
           for (var i = 0; i < input.length; i++) {
               output[i] = 0;
           }
       } else {
           for (var i = 0; i < input.length; i++) {
               output[i] = totalProduct / input[i];
           }
       }
       System.out.println(Arrays.toString(output));
    }
}
