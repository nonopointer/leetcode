package easy;

import java.util.stream.Collectors;
import java.util.*;

public class _414 {
    public static void main(String[] args){
        thirdMax(new int[]{1,2});
    }
    public static int thirdMax(int[] nums) {
        List<Integer> l = new ArrayList<>();
        for(int num: nums){
            l.add(num);
        }
        List<Integer> list = l.stream().distinct().sorted().collect(Collectors.toList());
        // Arrays.sort(nums)
        if(list.size()<3) return list.get(list.size()-1);
        return list.get(list.size()-3);
    }
}
