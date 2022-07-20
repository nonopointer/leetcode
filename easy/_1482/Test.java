package easy._1482;

public class Test {
    public static void main(String[] args){
        int[] bloomDay  = new int[]{};
        int m = 0;
        int k = 1;
        System.out.println(new Test().minDays(bloomDay, m, k));
    }
    public int minDays(int[] bloomDay, int m, int k) {
        int len = bloomDay.length;
        if(len < m * k) return -1;
        int min = bloomDay[0],max = bloomDay[0];
        for(int x : bloomDay){
            if(x < min) min = x;
            if(x > max) max = x; 
        }
        // int mid = (max+min)/2;
        // boolean flag = check(bloomDay,min,max);
        

        return 0;
    }

    // private boolean check(int[] bloomDay, int min, int max) {
        
    //     return false;
    // }
}