package easy.offer42;

public class Test {

    public static void main(String[] args){
        Test t = new Test();
        int[] nums = new int[]{-2,1,-3,4,-1,2,1,-5,4};
        System.out.println(t.maxSubArray(nums));
    }
    public int maxSubArray(int[] nums) {
        if(nums.length == 0){
            return -1;
        }
        if(nums.length == 1){
            return nums[0];
        }
        int result = nums[0],tem = 0;
        for(int i = 0;i < nums.length; i++){
            // result = Math.max(result, result + nums[i]);
            if(nums[i] < 0){
                result = nums[i];
                tem = 0;
            }else{
                tem += nums[i];
            }
            result = Math.max(result, tem);
        }
        return result;
    }
}
