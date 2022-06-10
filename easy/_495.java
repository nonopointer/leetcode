package easy;

public class _495 {
    public static void main(String[] args) {
        System.out.println(findPoisonedDuration(new int[] { 1, 3, 5, 7, 9, 11, 13, 15 }, 3));
    }

    public static int findPoisonedDuration(int[] timeSeries, int duration) {
        // int cnt = 0;
        // for (int )

        // int[] arr = new int[timeSeries[timeSeries.length-1]+duration];
        // for (int i : timeSeries){
        // int step = 0;
        // while (step < duration){
        // arr[i+step++] = 1;
        // }
        // }
        // int cnt = 0;
        // for (int i :arr){
        // cnt += i;
        // }
        // [1,3,5,7,9,11,13,15]
        // 3
        int res = 0;
        int len = timeSeries.length;
        for (int i = 1; i < len; i++) {
            if (timeSeries[i] >= timeSeries[i - 1] + duration) {
                res = res + duration;
            } else {
                res = res + (timeSeries[i] - timeSeries[i - 1]);
            }
        }
        return res + duration;

    }
}
