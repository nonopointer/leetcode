package easy.offer10_1;

public class T {
    public static void main(String[] args){
        T t = new T();
        System.out.println(t.fib(43));
    }
    public int fib(int n) {
        if (n < 2) return n;
        int rs = 0,a = 0,b = 1, idx = 2;
        while (idx <=n){
            rs = a + b;
            if (rs >= 1000000007) rs -= 1000000007;
            a = b;
            b = rs;
            idx ++;
        }
        // int x = fib(n-1) + fib(n-2);
        // if (x >= 1000000007) x -= 1000000007;
        return rs;
    }
}
