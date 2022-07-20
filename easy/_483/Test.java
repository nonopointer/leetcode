package easy._483;

import java.util.ArrayList;
import java.util.List;

public class Test {
    public static void main(String[] args){
        System.out.println(smallestGoodBase("1000000000000000000"));
        // test();
    }

    public static String smallestGoodBase(String s) {
        Long n = Long.parseLong(s);
        for (long i = 2L; i < n ;i++){
            Long  x = 1L;
            List<Long> list = new ArrayList<>();
            while (x < n){
                list.add(x);
                x *= i;
            }
            Long result = 0L;
            for (Long a: list){
                result += a;
            }
            if(result .equals(n)  ) 
                return i+"";
        }
        return "";
    }

    // private static void test(){

    //     int a = 1;
    //     List<Integer> l  = new ArrayList<>();
    //     while (a <5000){
    //         l.add(a);
    //         a *= 8;
    //     }
    //     int r = 0;
    //     for ( int x : l){
    //         System.out.print(x+"\t");        
    //     }
    //     System.out.println();
    //     for ( int x : l){
    //         r += x;
    //         System.out.print(r+"\t");        
    //     }
        
    // }
}
