package easy;

import java.util.*;

public class _500 {
    public static void main(String[] args){
        // System.out.println((int)'a' + "\t"+ (int)'A');
        for (String s : findWords(new String[]{"Hello","Alaska","Dad","Peace"})){
            System.out.println(s);
        }
    }

    public static String[] findWords(String[] words) {
        int[] q = new int[26];
        int[] a = new int[26];
        int[] z = new int[26];

        List<String> list = new ArrayList<>();

        q['Q'-'A'] = 1;
        q['W'-'A'] = 1;
        q['E'-'A'] = 1;
        q['R'-'A'] = 1;
        q['T'-'A'] = 1;
        q['Y'-'A'] = 1;
        q['U'-'A'] = 1;
        q['I'-'A'] = 1;
        q['O'-'A'] = 1;
        q['P'-'A'] = 1;

        a['A'-'A'] = 1;
        a['S'-'A'] = 1;
        a['D'-'A'] = 1;
        a['F'-'A'] = 1;
        a['G'-'A'] = 1;
        a['H'-'A'] = 1;
        a['J'-'A'] = 1;
        a['K'-'A'] = 1;
        a['L'-'A'] = 1;

        z['Z'-'A'] = 1;
        z['X'-'A'] = 1;
        z['C'-'A'] = 1;
        z['V'-'A'] = 1;
        z['B'-'A'] = 1;
        z['N'-'A'] = 1;
        z['M'-'A'] = 1;
        // q['q'-'A'] = 1;

        for(String s :words){
            int idx = s.charAt(0) > 96 ? s.charAt(0) - 32 : s.charAt(0);
            if(q[idx-'A']==1){
                check(s,q,list);
            }else if(a[idx-'A'] == 1){
                check(s, a, list);
            }else {
                check(s, z, list);
            }
        }
        String[] rs = new String[list.size()];
        int idx = 0;
        for(String s : list){
            rs[idx++] = s;
        }
        return rs;
    }

    private static void  check(String s, int[] arr, List<String> list) {
        for(char c : s.toCharArray()){
            int idx = c > 96 ? c - 32 : c;
            if (arr[idx-'A'] !=1) return;
        }
        list.add(s);
    }
}
