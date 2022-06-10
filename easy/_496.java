package easy;

import java.util.*;

public class _496 {
    public static void main(String[] args) {
        for(int i : nextGreaterElement(new int[]{2,4}, new int[]{1,2,3,4})){
            System.out.println(i);
        }
    }

    public static int[] nextGreaterElement(int[] nums1, int[] nums2) {
        Stack<Integer> stack = new Stack<>();
        Map<Integer, Integer> map = new HashMap<>();
        int tem = nums2[0];
        for (int i : nums2) {
            if (i > tem) {
                while (!stack.isEmpty()) {
                    int k;
                    if ( ( k = stack.pop())<i){
                        map.put(k, i);
                    }else{
                        stack.push(k);
                        break;
                    }
                }
            }
            stack.push(i);
            tem = i;
        }
        for (int i = 0; i < nums1.length; i++) {
            nums1[i] = map.getOrDefault(nums1[i],-1);
        }
        return nums1;
    }
}
