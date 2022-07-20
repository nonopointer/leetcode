package mid;


public class _869 {
    public static void main(String[] args){
        int i = 1;
        while (i<1000){
            System.out.println(reorderedPowerOf2(i));
            i <<=1;
        }
        // System.out.println(reorderedPowerOf2(218));
    }
    public static boolean reorderedPowerOf2(int n) {
        int[] map = new int[10];
        int w = 0;
        int max = 0x80000000;
        while(n>0){
            int a = n%10;
            max = Math.max(max, a);
            // map.put(a, map[a]+1);
            map[a]++;
            n /= 10;
            ++w;
        }
        int max_value = max;
        --w;
        while ( w > 0){
            max_value = max_value*10 + 9;
            --w;
        }
        int target = 1;
        while (target <= max_value){
            if ( checkValue(map,target)) return true;
            target <<= 1;
        }
        return false;
    }
/*
    public static boolean reorderedPowerOf(int n) {
        Map<Integer,Integer> map = new HashMap<>();
        int w = 0;
        int max = 0x80000000;
        while(n>0){
            int a = n%10;
            max = Math.max(max, a);
            map.put(a, map.getOrDefault(a, 0)+1);
            n /= 10;
            ++w;
        }
        int max_value = max;
        --w;
        while ( w > 0){
            max_value = max_value*10 + 9;
            --w;
        }
        int target = 1;
        while (target <= max_value){
            if ( checkValue2(map,target)) return true;
            target <<= 1;
        }
        
        return false;
    }
*/
    private static boolean checkValue(int[] map,int target ) {
        int n = target;
        int[] targetMap = new int[10];
        int w = 0;
        while (n>0){
            int a = n%10;
            // targetMap.put(a, targetMap.getOrDefault(a, 0)+1);
            targetMap[a]++;
            n /= 10;
            ++w;
        }
        int b = 0;
        for (int  count : map){
            b+=count;
        }

        if (b!=w) return false;
        for (int i = 0;i< 10;i++){
            if(map[i]!=targetMap[i]) return false;
        }
    
        return true;
    }
/*
    private static boolean checkValue2(Map<Integer, Integer> map,int target ) {
        int n = target;
        Map<Integer,Integer> targetMap = new HashMap<>();
        int w = 0;
        while (n>0){
            int a = n%10;
            targetMap.put(a, targetMap.getOrDefault(a, 0)+1);
            n /= 10;
            ++w;
        }
        int b = 0;
        for (Map.Entry<Integer,Integer> entry : map.entrySet()){
            b+=entry.getValue();
        }

        if (b!=w) return false;
        for (Map.Entry<Integer,Integer> entry : targetMap.entrySet()){
            if(!entry.getValue().equals(map.get(entry.getKey()))){
                return false;
            }
        }

        return true;
    }*/
}
