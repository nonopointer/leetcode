package easy._17_4;

import java.util.stream.Stream;

public class Test {
    public static void main(String[] args){
        Test t = new Test();
        int[] arr = t.smallestK(new int[]{1,2,3,4,5,6,7,7,8,8}, 3);
        // Stream.of(arr).forEach(System.out::println);
        for (int a : arr ){
            System.out.println(a);
        }
    }

    public int[] smallestK(int[] arr, int k) {
        int[] rs = new int[k];
        int hight = 0;
        for(int i = 0;i < k;i++ ){
            rs[i] = arr[i];
            if (rs[i] > rs[hight]){
                hight = i;
            }
        }

        for (int i = k ; i < arr.length ; i ++ ){
            if (arr[i] < rs[hight]){
                rs[hight] = arr[i];
                hight = foundBiggest(rs);
            }
        }
        return rs;
    }

    public int foundBiggest(int[] arr){ 
        int a = 0;
        for (int i = 0;i <arr.length;i++ ){
            if (arr[a]< arr[i]){
                a = i;
            }
        }
        return a;
    }
}
