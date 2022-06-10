package hard;

public class _339 {
    public static void main(String[] args){
        
    }
    public boolean isSelfCrossing(int[] nums) {
        int size = nums.length;
        if(size<4) return false;
        for(int  i=3;i<size;i++){
            if(nums[i-1]<=nums[i-3] && nums[i]>=nums[i-2]) return true;
            if(i>=4 && nums[i]+nums[i-4] == nums[i-2]  &&  nums[i-1]==nums[i-3]) return true;
            if(i>=5 && nums[i-1] <= nums[i-3] && nums[i-2] > nums[i-4] &&
                nums[i]+nums[i-4] >= nums[i-2]  &&  nums[i-5] +nums[i-1]>=nums[i-3]) return true;
        }
        return false;
    }
}
