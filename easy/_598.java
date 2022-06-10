package easy;

public class _598 {
    public static void main(String[] args){
        System.out.println(maxCount(3, 3 , new int[][]{{2,2},{3,3}}));
    }
    public static int maxCount(int m, int n, int[][] ops) {
        if(ops.length == 0) return m * n;
        int x=m,y = n;
        for (int idx = 0; idx < ops.length ; idx ++){
            x = Math.min(x,ops[idx][0]);
            y = Math.min(y,ops[idx][1]);
        }
        return x * y;
    }
}
