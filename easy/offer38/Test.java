package easy.offer38;

import java.util.HashSet;
import java.util.Set;
import java.util.stream.Stream;

public class Test {
    public static void main(String[] args){
        Test t  = new Test();
        String s = "abc";
        String[] ss = t.permutation(s);
        Stream.of(ss).forEach(System.out::println);
    }

    Set<String> set = new HashSet<>();
    boolean[] flag = new boolean[8];
    public String[] permutation(String s) {
        char[] cs = s.toCharArray();
        dfs(cs,0,"");
        String[] res = new String[set.size()];
        System.out.println(set.size());
        int idx = 0;
        for(String str : set) res[idx++] = str;
        return res;
    }
    public void dfs(char[] cs,int idx,String s){
        int len = cs.length;
        if(len == idx){
            set.add(s);
            return;
        }
        for(int i = 0;i<len ;i++){
            if(!flag[i]){
                flag[i] = true;
                dfs(cs,idx+1,s+cs[i]);
                flag[i] = false;
            }
        }
    }
    // void dfs(char[] cs, int u, String cur) {
    //     int n = cs.length;
    //     if (u == n) {
    //         set.add(cur);
    //         return;
    //     }
    //     for (int i = 0; i < n; i++) {
    //         if (!flag[i]) {
    //             flag[i] = true;
    //             dfs(cs, u + 1, cur + String.valueOf(cs[i]));
    //             flag[i] = false;
    //         }
    //     }
    // }

}
