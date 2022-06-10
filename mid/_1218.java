package mid;

public class _1218 {
    public static void main(String[] args){
        System.out.println(longestSubsequence(new int[]{1,2,3,4}, 1));
    }

    public static int longestSubsequence(int[] arr, int difference) {
        int[] t = new int[40001];
        int d = 20000,max=0;
        for(int a : arr){
            t[a+d]=t[a+d-difference]+1;
            max=Math.max(t[a+d],max);
        }
        return max;
    }
}
