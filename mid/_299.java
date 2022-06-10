package mid;
import java.util.*;

public class _299 {
    public static void main(String[] args){
        System.out.println(getHint("1807"  ,"7810"));
    }
    public static String getHint(String secret, String guess) {
        int[] arr = new int[10];
        int[] arr2 = new int[10];
        int count1 = 0,count2 =0;
        for (int i = 0 ; i< secret.length(); i++ ){
            if(secret.charAt(i)==guess.charAt(i)){
                count1++;
            } else {
                arr[secret.charAt(i)-'0'] ++;
                arr2[guess.charAt(i)-'0'] ++;
            }
        }
        for(int i = 0;i < 10; i++){
            count2 += Math.min(arr[i], arr2[i]);
        }

        return count1+"A"+count2+"B";
    }
}
