package hard._149hard;

import java.math.BigDecimal;
import java.util.HashMap;
import java.util.Map;

public class Test {
    public static void main(String[] args){
        Test t = new Test();
        int[][] arr = new int[][]{{1,3},{2,1},{3,4}};
        System.out.println(t.maxPoints(arr));
    }
    public int maxPoints(int[][] points) {
        if(points.length==1) return 1;
        
        for (int i = 0; i < points.length; i++) {
            for(int j = i+1;j<points.length;j++){
                Map<String,Double> map = get(points[i], points[j]);
                System.out.println(map.get("a")+"\t"+map.get("b"));
            }
        }

        return -1;
    }

    public Map<String,Double> get(int[] p1,int[] p2){
        //{1,1},
        //{2,2},
        //{3,4}
        Map<String,Double> map = new HashMap<>();
        
        Double a = new BigDecimal(p1[1]+p2[1]+"")
        .divide(new BigDecimal(p1[0]+p2[0]+""))
        .doubleValue();
        Double b = new BigDecimal((p1[0] * p2[1] - p1[0] * p2[1])+"").divide(new BigDecimal(p2[0]-p1[1]+"")).add(new BigDecimal(p1[1]+"")).doubleValue();
        map.put("a", a);
        map.put("b", b);
        return map;
    }
}
