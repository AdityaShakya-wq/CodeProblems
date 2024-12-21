import java.util.*;

class Main {
    public static void main(String[] args) {
        var input = new int[]{1, 2, 3, 6, 3, 6, 1};
        Set<Integer> tempList = new HashSet<>();
        Set<Integer> output = new HashSet<>();
        for(var in : input) {
            if (tempList.contains(in)) {
                output.add(in);
            }else {
                tempList.add(in);
            }
        }
        
        System.out.println(output);
    }
}
