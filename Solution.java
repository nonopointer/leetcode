import java.util.concurrent.ConcurrentLinkedDeque;

/*
 * @lc app=leetcode.cn id=141 lang=java
 *
 * [141] 环形链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) {
 *         val = x;
 *         next = null;
 *     }
 * }
 */
  class ListNode {
      int val;
      ListNode next;
      ListNode(int x) {
          val = x;
          next = null;
      }
  }
public class Solution {
    public static void main(String[] args){

        ConcurrentLinkedDeque deque = new ConcurrentLinkedDeque<>();
        deque.add(1);
    }
    public boolean hasCycle(ListNode head) {
        if (head == null) return false;
        ListNode n = head;
        int x = n.val;
        while (n.next !=null){
            n = n.next;
            x = n.val;
        }
        n = head;
        while (n.next !=null){
            if(n.val==x) return true;
            n = n.next;
        }
        // if (n.next == null && n.val== x)
        return false;
    }
}
// @lc code=end

