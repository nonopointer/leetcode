package easy;

public class jianzhi069 {
    public static void main(String[] args){
        System.out.println(peakIndexInMountainArray(new int[]{24,69,100,99,79,78,67,36,26,19}));
    }
    public static int peakIndexInMountainArray(int[] arr) {
        int i = 1;
        while (i < arr.length -1){
            if (arr[i]>arr[i+1]) return i;
            ++i;
        }
        return i;
    }
}
