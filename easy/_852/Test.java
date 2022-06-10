package easy._852;

public class Test {
    public static void main(String[] args){

        System.out.println(peakIndexInMountainArray(new int[]{0,1,0}));
    }
    public static int peakIndexInMountainArray(int[] arr) {
        // System.out.println(arr[1]);
        for(int i = 0;i < arr.length-1;i++){
            // System.out.println(arr[i]);
            if (arr[i] > arr[i+1]) 
            return i;
        }
        return -1;
    }
}
