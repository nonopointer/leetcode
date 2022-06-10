package easy._1588;

public class Test {
    public static void main(String[] args){
        Test t = new Test();
        int[] arr = new int[]{1,4,2,5,3};
        System.out.println(t.sumOddLengthSubarrays(arr));
        System.out.println(t.sumOne(arr, 0, 1));
        System.out.println(t.sum(arr, 1));
    }

    public int sumOddLengthSubarrays(int[] arr) {
        int len = arr.length;
        int rs = 0;
        for (int i = 1 ; i <= len ; i = i+2){
            rs += sum(arr,i);
        }

        return rs;
    }

    public int sum(int[] arr , int len){
        int rs = 0;
        
        for(int i = 0;i+len <=arr.length;i++){
            rs += sumOne(arr,i,len);
        }

        return rs;
    }

    public int sumOne(int[] arr,int start,int len){
        int rs = 0;
        for(int i = start;i < start + len;i++){
            rs += arr[i];
        }
        return rs;
    }

}
