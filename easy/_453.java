package easy;

public class _453 {
    public static void main(String[] args){
        
    }
    public static int minMoves(int[] nums) {
        int min = 0x7fffffff;
        for (int n : nums){
            if(n<min) min = n; 
        }
        int count = 0;
        for (int n : nums){
            count += (n-min);
        }
        return count;
    }
}
