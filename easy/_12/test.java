package easy._12;

import java.util.HashMap;
import java.util.Map;
import java.util.UUID;

public class test {
    /**
     * 字符 数值 
     * I 1 
     * V 5 
     * X 10
     * L 50 
     * C 100 
     * D 500 
     * M 1000
     * 
     * @param args
     */
    public static void main(String[] args) {

        System.out.println(UUID.randomUUID().toString().replace("-", ""));
        System.out.println(intToRoman(153));
    }
    static Map<Integer,String> map = new HashMap<>();

    {
        map.put(1,"I");
        map.put(5,"V");
        map.put(10,"X");
        map.put(50,"L");
        map.put(100,"C");
        map.put(500,"D");
        map.put(1000,"M");
    }
    //1 <= num <= 3999
    public static String intToRoman(int num) {
        // StringBuilder s = new StringBuilder();
        int t = num,idx = 0;
        int[] arr = new int[4];
        while(t>0){
            arr[idx++] = (t % 10 ) - 5;
            t /=10;
        }
        
        

        return "";
    }
}