package easy.offer15;

public class Test {
    public static void main(String[] args){
        // Test t = new Test();
        // int a = 0b10100011;
        // System.out.println(Integer.toBinaryString(0x55555555));
        // System.out.println(Integer.toBinaryString(0x33333333));
        // System.out.println(Integer.toBinaryString(0x0f0f0f0f));
        // System.out.println(a + "\t"+ t.hammingWeight(a));
        for(int i= 0;i < 20;i++){
            String emoji3 = String.valueOf(Character.toChars(0x1F448+i));
            System.out.println(emoji3);
        }

    }
    public int hammingWeight(int i) {
        // HD, Figure 5-2
        i = i - ((i >>> 1) & 0x55555555);
        System.out.println(Integer.toBinaryString(i));
        i = (i & 0x33333333) + ((i >>> 2) & 0x33333333);
        System.out.println(Integer.toBinaryString(i));
        i = (i + (i >>> 4)) & 0x0f0f0f0f;
        System.out.println(Integer.toBinaryString(i));
        i = i + (i >>> 8);
        System.out.println(Integer.toBinaryString(i));
        i = i + (i >>> 16);
        System.out.println(Integer.toBinaryString(i));
        System.out.println(Integer.toBinaryString(i & 0x3f));
        return i & 0x3f;

    }
}
