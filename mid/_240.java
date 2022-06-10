package mid;

public class _240 {
    public static void main(String[] args){
        // System.out.println(searchMatrix(new int[][]{{1,4,7,11,15},{2,5,8,12,19},{3,6,9,16,22},{10,13,14,17,24},{18,21,23,26,30}}, 20));
        System.out.println(searchMatrix(new int[][]{{-5}}, -10));
    }
    public static boolean searchMatrix(int[][] matrix, int target) {
        int row = matrix.length;
        int col = matrix[0].length;
        int x = col-1,y = 0;
        while (x >= 0 && y < row){
            if(matrix[y][x]==target) return true;
            if(matrix[y][x]>target){
                --x;
                continue;
            } 
            if(matrix[y][x]<target) {
                ++y;
                continue;
            } 
        }
        return false;
    }
}
