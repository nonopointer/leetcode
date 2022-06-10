package easy._470;

import java.util.Random;

public class T {
    public static void main(String[] args){
        int  n = 10;
        for (int i = 0; i < n; i++) {
            System.out.println(rand10());
        }
    }


    /**
     * 
     *    1 2 3 4 5 6 7 
     * 
     * 
     * 1 2 3 4 5 6 7 8 9 10
     * 
     */
    public static int rand7(){
        Random r = new Random();
        return r.nextInt(7)+1;
    }
    

    public static int rand10() {
        int a = rand7() ,b = rand7();
        while (a < 4 && b > 4){
            a = rand7();
            b = rand7();
        }
        return (a + b)% 10 + 1;
    }
}
