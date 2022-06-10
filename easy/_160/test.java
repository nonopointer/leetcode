package easy._160;

import java.util.HashSet;
import java.util.Set;

public class test {
    public static void main(String[] args){
        int i = "".hashCode();
    }
    class ListNode {
             int val;
            ListNode next;
             ListNode(int x) {
                 val = x;
                 next = null;
             }
         }
    public ListNode getIntersectionNode(ListNode headA, ListNode headB) {
        if (headA ==null || headB == null) return null;
        Set<Integer> set = new  HashSet<>();
        while(headA!=null){
            set.add(headA.hashCode());
            headA = headA.next;
        }
        while(headB!=null){
            if(set.contains(headB.hashCode())) return headB;
            headB = headB.next;
        }
        return null;
    }
}
