package easy._156;

public class Test {
    public static void main(String[] args){
        Test t  = new Test();
       System.out.println( t.compareVersion("1","1.1"));;
    }
    public int compareVersion(String version1, String version2) {
        String[] ss1 = version1.split("\\.");
        String[] ss2 = version2.split("\\.");
        int[] i1 = new int[ss1.length];
        int[] i2 = new int[ss2.length];

        for(int i = 0;i < ss1.length ;i ++){
            i1[i] = str2int(ss1[i]);
        }

        for(int i = 0;i < ss2.length ;i ++){
            i2[i] = str2int(ss2[i]);
        }

        int i = 0 ;
        int minlen = Math.min(ss1.length,ss2.length);
        for(;i < minlen ; i ++ ){
            if(i1[i]<i2[i]){
                return -1;
            }
            if(i1[i]>i2[i]){
                return 1;
            }
        }
        if(ss1.length>ss2.length){
            while(i < ss1.length){
                if(i1[i++]>0){
                    return 1;
                }
            }
        }else if(ss1.length<ss2.length){
            while(i < ss2.length){
                if(i2[i++]>0){
                    return -1;
                }
            }
        }
        return 0;
    }
    public int str2int(String str){
        int idx = str.length()-1;
        int n = 1;
        int rs = 0;
        while (idx >=0 ){
            rs += Integer.parseInt(str.charAt(idx--)+"") * n;
            n *= 10;
        }
        return rs;
    }
}
