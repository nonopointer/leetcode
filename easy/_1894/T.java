package easy._1894;

public class T {
    public static void main(String[] args){
        System.out.println(chalkReplacer(new int[]{3,4,1,2}, 25));
    }
    public static int chalkReplacer(int[] chalk, int k) {
        int tem = 0;
        int total = 0;

        for (int  i= 0 ; i < chalk.length;i++){
            total += chalk[i];
            if(k < total){
                return i;
            }
        }

        for(int i = 0;i < chalk.length;i++){
            if((k%total)-tem <chalk[i]){
                return i;
            }
            tem += chalk[i];
        }
        return 0;
    }
}
