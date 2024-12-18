import java.util.Arrays;

class Main {
    public static void main(String[] args) {
        int[] input = {0, -1, 2, -3, 1};
        int target = -2;
        
        System.out.println("Sum Present :");
        System.out.println(checkTarget(input, target));
    
        target = 5;
        System.out.println("Sum Not Present :");
        System.out.println(checkTarget(input, target));
    }
    
    public static boolean checkTarget(int[] input, int target) {
        Arrays.sort(input);
        
        for (int i = 0; i < input.length; i ++) {
            var difference = target - input[i];
            if (binarySearch(input, i + 1, input.length -1, difference)) {
                return true;
            }
        }
        
        return false;
    }
    
    private static boolean binarySearch(int[] input, int left, int right, int difference) {
        while(left <= right) {
            int mid = left + (right-left)/2;
            if (input[mid] == difference) {
                return true;
            } else if (input[mid] < difference) {
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }
        
        return false;
    }
}
