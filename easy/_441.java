package easy;

public class _441 {
    public static void main(String[] args) {
        System.out.println(arrangeCoins(2));
        //test font RI
    }

    public static int arrangeCoins(int n) {
        int i = 1;
        int count = 0;
        while (count <= n-i){
            count += i;
            ++i;
        }
        return i-1;
    }
}
