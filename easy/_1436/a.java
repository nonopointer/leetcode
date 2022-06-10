package easy._1436;

import java.util.*;

public class a {
    public static void main(String[] args){
        List<String> l1 = new ArrayList<>(){{
            add("London");
            add("New York");
        }};
        List<String> l2 = new ArrayList<>(){{
            add("New York");
            add("Lima");
        }};
        List<String> l3 = new ArrayList<>(){{
            add("Lima");
            add("Sao Paulo");
        }};
        List<List<String>> list = new ArrayList<>(){{
            add(l1);
            add(l2);
            add(l3);
        }};

        System.out.println(destCity(list));
    }
    public static String destCity(List<List<String>> paths) {
        Map<String,String> map = new HashMap<>();
        for(List<String> list : paths){
            map.put(list.get(0),list.get(1));
        }
        for (Map.Entry<String,String> entry : map.entrySet()){
            if (!map.containsKey(entry.getValue())){
                return entry.getValue();
            }
        }
        return null;
    }
}
