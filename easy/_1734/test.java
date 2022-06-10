package easy._1734;


public class test {
    public static void main(String[] args){
        int[] arr = {3,1};
        // for (int x : decode(arr)){
        //     System.out.println(x);
        // }
        // Stream.of(decode(arr)).forEach(System.out::print);
        // int b =0;
        // for (int i = 1; i <= 3 ; i++) {
        //     b ^= i;
        //     System.out.println(b);
        // }
        
        print0(decode(arr));

    }
    public static int[] decode(int[] encoded) {
        int all = 0;
        int n = encoded.length; 
        int[] temp = new int[n+1];
        for (int i = 1; i <= n+1 ; i++) all ^= i;
        for (int i = 1;i < n;i = i+2) all ^= encoded[i];
        temp[0] = all;
        for (int i = 0; i < n ; i++ ) temp[i + 1] = temp[i] ^ encoded[i];
        return temp;//[1,2,3]
    }
    static void print0(int[] arr){
        for (int x : decode(arr)){
            System.out.print(x+"\t");
        }
        System.out.println();
    }
}