package easy;

public class _492 {
    public static void main(String[] args){
        for (int i : constructRectangle(7)){
            System.out.println(i);
        }
    }
    public static int[] constructRectangle(int area) {
        int x = (int)Math.sqrt(area);
        int y = x;
        while (x > 0){
            if ( area % x == 0){
                break;
            }
            --x;
        }
        y = area / x;
        return new int[]{y,x};
    }
}
