package easy._554;

import java.util.ArrayList;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

public class test {
    public static void main(String[] args){
        //[[1,2,2,1],[3,1,2],[1,3,2],[2,4],[3,1,2],[1,3,1,1]]
        List<List<Integer>> wall = new ArrayList<>();
        // for(int i = 0 ; i < 3 ;i++ ){
        //     List<Integer> row = new ArrayList<>();
        //     row.add(100000000);
        //     wall.add(row);
        // }

        int[] r1 = {1,2,2,1};
        int[] r2 = {3,1,2};
        int[] r3 = {1,3,2};
        int[] r4 = {2,4};
        int[] r5 = {3,1,2};
        int[] r6 = {1,3,1,1};
        System.out.println(leastBricks(wall));
    }
    static List<List<Integer>>  rows = null;
    public static int leastBricks(List<List<Integer>> wall) {
        if (wall == null || wall.get(0)==null) return -1;
        rows = wall;
        int len = 0,result = 0x7fffffff;
        for (Integer l : wall.get(0)){
            len += l;
        }
        Set<Float> set = new HashSet<>();
        set.add(0.5f);
        for(List<Integer> row : rows){
            int rowidx = 0;
            for(Integer col : row){
                rowidx +=col;
                set.add(rowidx+0.5f);
            }
        }
        
        for(float s : set){
            int n = across(s);
            if(n == 0) continue;
            result = result < n ? result : n;
        }
        if(result == 0x7fffffff) return wall.size();
        return result;
    }

    private static int across(float s) {
        int n = 0;
        for(List<Integer> row : rows){
            int len = 0;
            for(Integer col : row){
                len += col;
                if(len == s) break;
                if(len > s ){
                    n++;
                    break;
                }
            }
        }
        return n;
    }
}