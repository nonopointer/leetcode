package easy;

public class _66 {
    public static void main(String[] args){
        for (int i : plusOne(new int[]{0})){
            System.out.println(i);
        }
    }

    public static int[] plusOne(int[] digits) {
        int idx = digits.length-1;
        while(idx >= 0){
            if (digits[idx]==9){
                digits[idx--]=0;
            }else {
                digits[idx]++;
                break;
            }
        }
        int[] rs = digits;
        if(digits[0]==0){
            rs = new int[digits.length+1];
            rs[0]= 1;
        }
        return rs;
    }
}
