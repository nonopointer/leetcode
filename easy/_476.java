package easy;

public class _476 {
    public static void main(String[] args){
        System.out.println(findComplement(5));
    }
    public static int findComplement(int num) {
        int i = 1;
        while (i <= num ){
            if((i&num)==num) break;
            i = (i<<1)+1;
        }
        return i ^ num;
    }
}
