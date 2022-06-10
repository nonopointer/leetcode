package easy._482;

public class a {
    public static void main(String[] args) {

        System.out.println((int)'Z');
        System.out.println((int)'a');
    }

    public String licenseKeyFormatting(String s, int k) {
        int count = 0;
        s = s.replace("-", "").toUpperCase();
        if (s == null || s.length() == 0 || s.equals("")) return "";
        char[] cs = s.toCharArray();
        StringBuilder sb = new StringBuilder();
        for (int idx = cs.length - 1; idx >= 0; idx--) {
            sb.append(cs[idx]);
            ++count;
            if (count == k) {
                count = 0;
                sb.append('-');
            }
        }
        if (cs.length % k == 0)
            sb.deleteCharAt(sb.length() - 1);
        return sb.reverse().toString();
    }
}
