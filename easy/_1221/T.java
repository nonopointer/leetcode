package easy._1221;

import java.util.Stack;

public class T {
    public static void main(String[] args){
        System.out.println(balancedStringSplit("R"));
    }
    public static int balancedStringSplit(String s) {
        
        Stack<Character> stack = new Stack<>();
        for (char c : s.toCharArray()){
            stack.push(c);
        }
        int i = 0,count = 1;
        char first = stack.pop();
        while(!stack.isEmpty()){
            char c = stack.pop();
            if (c == first){
                count++;
            }else{
                count --;
            }
            if (count == 0) {
                i ++ ;
                if (!stack.isEmpty()) {
                    first = stack.pop();
                    count = 1;
                }
            }
        }
        if (count != 0){
            i=1;
        }
        return i;
    }
}
