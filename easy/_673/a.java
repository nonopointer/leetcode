package easy._673;


public class a {
    // public static void main(String[] args){
    //     System.out.println(findNumberOfLIS(new int[]{1,3,5,4,7}));
    // }
    // public static int findNumberOfLIS(int[] nums) {
    //     Map<Integer,LinkedList<Integer>>  map = new HashMap<>();
    //     int end =  nums.length-1;
    //     map.put(nums[end], new LinkedList(){{
    //         add(nums[end]);
    //     }});
        
    //     for (int i = end-1; i > 0; --i) {
    //         for (Map.Entry<Integer,LinkedList<Integer>> entry: map.entrySet()){
    //             LinkedList<Integer> tem = entry.getValue();
    //             if(tem.getLast() > nums[i]){
    //                 tem.addLast(nums[i]);
    //             }
    //             if(tem.getLast() < nums[i]){
    //                 tem = new LinkedList<>();
    //                 tem.addLast(nums[i]);
    //             }
    //             map.put(entry.getKey(), tem);
    //         }
        
    //     }
    //     int maxlen = 0,count = 0;
    //     for (Map.Entry<Integer,LinkedList<Integer>> entry: map.entrySet()){
    //         if(entry.getValue().size() == maxlen){
    //             ++ count ;
    //         }
    //         if(entry.getValue().size() > maxlen){
    //             count = 1;
    //         }
    //     }
    //     return count;
    // }
}
