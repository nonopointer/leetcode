package mid;

import java.util.*;

public class _260 {
    public static void main(String[] args){
        for (int i : singleNumber1(new int[]{1,2,1,3,2,5})){
            System.out.println(i);
        };
    }
    public static int[] singleNumber(int[] nums) {
        int idx = 0;
        int[]rs = new int[2];
        Map<Integer,Integer> map = new HashMap<>();
        for (int i : nums){
            map.put(i, map.getOrDefault(i, 0)+1);
        }
        for (Map.Entry<Integer,Integer> entry : map.entrySet()){
            if(entry.getValue() ==1){
                rs[idx++] = entry.getKey();
            }
        }
        return rs;
    }

    public static int[] singleNumber1(int[] nums) {
        if(nums.length == 2 ) return nums;
        List<Integer> a = new ArrayList<>();
        List<Integer> b = new ArrayList<>();
        int[] rs = new int[2];

        int tem = 0;
        for (int i : nums){
            tem ^= i;
        }
        // int idx = 1;
        // while(idx < tem){
        //     if((tem & idx )== idx) break;
        //     idx <<= 1;
        // }
        tem = (tem ^ (tem-1)) & tem;
        for(int i : nums){
            if((i & tem)==tem){
                // a.add(i);
                rs[0] ^= i;
            }else{
                // b.add(i);
                rs[1] ^= i;
            }
        }

        // for(int i : a){
        //     rs[0] ^= i;
        // }

        // for(int i : b){
        //     rs[1] ^= i;
        // }
        return rs;
    }
}
