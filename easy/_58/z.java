package easy._58;

public class z {
    public static void main(String[] args){
        System.out.println(lengthOfLastWord( "   fly me   to   the moon  "));
    }
    public static int lengthOfLastWord(String s) {
        boolean f = false;
        int count = 0;
        for (int i = s.length()-1; i >= 0; i--) {
            if(s.charAt(i) != ' '){
                f = true;
                count ++;
            }
            if(s.charAt(i) == ' ' && f){
                break;
            }
        }
        return count;
    }
}
