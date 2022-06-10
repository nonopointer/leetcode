package easy;

import java.util.*;

public class _412 {
    public static void main(String[] args) {
        System.out.println(fizzBuzz(15));
    }

    public static List<String> fizzBuzz(int n) {
        List<String> rs = new ArrayList<>();
        String[] arr = new String[n + 1];
        for (int i = 0; i < arr.length; i++) {
            arr[i] = i + "";
        }
        int m = 1;
        while (m * 3 <= n) {
            if ("Buzz".equals(arr[m * 3])) {
                arr[m * 3] = "FizzBuzz";
            } else {
                arr[m * 3] = "Fizz";
            }
            if (m * 5 > n) {
                m++;
                continue;
            }
            arr[m * 5] = "Buzz";

            m++;
        }
        for (int i = 1; i < arr.length; i++) {
            rs.add(arr[i]);
        }
        return rs;
    }
}
