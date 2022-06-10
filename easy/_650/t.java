package easy._650;

public class t {
    public static void main(String[] args) {
        // System.out.println(minSteps(9));
        for (int i = 1; i < 10; i++) {
            System.out.println(minSteps(i));
        }
    }

    public static int minSteps(int n) { // 3
        if (n == 1)
            return 0;
        for (int i = 2; i < n; i++) {
            if (n % i == 0) {
                return minSteps(i) + minSteps(n / i);
            }
        }
        return n;
    }

}
