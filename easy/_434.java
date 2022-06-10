package easy;

public class _434 {
    public static void main(String[] args){
        
    }
    public static int countSegments(String s) {
        int n = s.length();
        if(n == 0)
            return 0;
        char[] c = s.toCharArray();
        int ans = 0;
        for(int i = 0; i < n; i++){
            if(c[i] != ' '){
                ans++;
                while(i < n && c[i] != ' '){
                    i++;
                }
            }
        }
        return ans;
    }
}
