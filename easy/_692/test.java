package easy._692;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class test {
    public static void main(String[] args){
        
    }
    public static List<String> topKFrequent(String[] words, int k) {
        Map<String,Integer> map = new HashMap<>();
        for(String s : words){
            map.put(s, map.getOrDefault(s, 0)+1);
        }
        List<Map.Entry<String,Integer>> list = new ArrayList<>();
        for(Map.Entry<String,Integer> e :  map.entrySet()){
            list.add(e);
        }
        Collections.sort(list,(x,y)-> y.getValue()-x.getValue());
        List<String> result = new ArrayList<>(k);
        for(int i = 0;i<k;i++) result.add(list.get(i).getKey());
        return result;
    }
}