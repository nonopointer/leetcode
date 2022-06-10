package easy._1310;

public class test {
    public static void main(String[] args){
        int[] arr = {1,3,4,8};
        int[][] q = {{0,1},{1,2},{0,3},{3,3}};
        for (int x : xorQueries(arr, q)) {
            System.out.print(x + "\t");
        }
    }
    public static int[] xorQueries(int[] arr, int[][] queries) {
        int len = queries.length;
        int[] result = new int[len];
        for (int i = 0 ;i < len; ++i){
            for(int  l = queries[i][0] ; l <= queries[i][1] ; ++l){
                System.out.println("i:"+i+"\tl:"+l+"");
                result[i] ^= arr[l];
            }
        }
        return result;
    }
}