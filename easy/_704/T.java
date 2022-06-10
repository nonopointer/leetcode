package easy._704;

public class T {
    public static void main(String[] args){
        // System.out.println(2&-1);
        int[] nums = new int[]{-1,0,5};
        int target = -1;
        System.out.println(search(nums, target));
    }
    public static int search(int[] nums, int target) {
        int l = 0,r = nums.length-1,m = (l + r)/2;
        if(nums.length == 1){
            return nums[0] == target ? 0 : -1;
        }
        while (l <= r && r >= 0 && l <nums.length){
            
            if (nums[m] == target){
                return m;
            }
            if (nums[m] > target){
                r = --m;
                m = (l + r) /2;
                continue;
            }
            if (nums[m] < target){
                l = ++m;
                m = (l + r) /2;
                continue;
            }
            
        }

        return -1;
    }
}
