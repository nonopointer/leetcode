package easy._643;

import java.text.DecimalFormat;

public class Test {
    /**给定 n 个整数，找出平均数最大且长度为 k 的连续子数组，并输出该最大平均数。 */
    public static void main(String[] args){
        System.out.println(findMaxAverage(new int[]{111,2,2,33}, 2));
    }

    public static double findMaxAverage(int[] nums, int k) {
        int len = nums.length;
        int sum = 0;
        DecimalFormat df=new DecimalFormat("0.00");
        
        if(len == 1 ){
            return nums[0];
        }
        
        for(int i = 0;i < k ; i++){
            sum += nums[i];
        }

        if(len == k){
            return Double.valueOf(df.format((double)sum/k));
        }

        for(int i = 0;i < len-k; i++){
            // sum = nums[i] > nums[i+k]? 0 : nums[i+k]-nums[i]; 
            if(nums[i] > nums[i+k]){
                
            }else{
                
            }
        }
        
        return Double.valueOf(df.format((double)sum/k));
    }
}