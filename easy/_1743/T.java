package easy._1743;

import java.util.*;

public class T {
    public static void main(String[] args){

        // System.out.println();
    }

    
    public int[] restoreArray(int[][] adjacentPairs) {
        /*
        Map<Integer,int[]> map = new HashMap<>();
        int[] result = new int[]{};
        for(int[] arr : adjacentPairs){
            int[] as = map.get(arr[0]);
            if(as != null){
                as[1] = arr[1];
            }else{
                int[] t = new int[2];
                t[0] = arr[1];
                map.put(arr[0], t);
            }
            as = map.get(arr[1]);
            if(as != null){
                as[1] = arr[0];
            }else{
                int[] t = new int[2];
                t[0] = arr[0];
                map.put(arr[1], t);
            }
        }
        StringBuilder sb = new StringBuilder();
        sb = 
        return result;*/

        if(adjacentPairs.length==0){
            return null;
        }
        if(adjacentPairs.length ==1){
            return adjacentPairs[0];
        }

        List<Integer> list = new LinkedList<>();
        list.add(adjacentPairs[0][0]);
        list.add(adjacentPairs[0][1]);
        return null;
    }



}
