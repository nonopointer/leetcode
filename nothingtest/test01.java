package nothingtest;

import java.util.*;

public class test01 {
    public static void main(String[] args) throws InterruptedException {

        Random r = new Random();
        test01 t = new test01();
        for(int i = 0;i < 10; i ++ ){
            System.out.println(t.generate(25));
        }
    }

    String generate(int n) {
        char[] tem = new char[62];
        int idx = 0;
        for (int i = 0; i < 26; i++) {
            tem[idx++] = (char) ('a' + i);
            tem[idx++] = (char) ('A' + i);
        }
        for(int i = 0 ; i < 10 ;i++){
            tem[idx++] =(char)('0'+ i);
        }
        Random r = new Random();
        char[] result = new char[n];
        idx = 0;
        while (idx < n) {
            char c = tem[r.nextInt(62)];
            while (idx != 0 && c == result[idx - 1]) {
                c = tem[r.nextInt(62)];
            }
            result[idx++] = c;
        }
        return new String(result);
    }
}