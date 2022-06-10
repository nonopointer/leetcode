package easy._430;

import java.util.*;
import java.util.Stack;

public class q {
    public static void main(String[] args) {

    }

    
    
    public Node flatten(Node head) {
        if(head==null)return null;
        if(head.child==null){
            flatten(head.next);
            return head;
        }
        Node first=flatten(head.child);
        Node second = flatten(head.next);
        head.next=first;
        first.prev=head;
        while(first.next!=null){
            first=first.next;
        }
        first.next=second;
        if(second!=null)
            second.prev=first;

        head.child=null;    
        return head;
    }
   

    // Definition for a Node.
    class Node {
        public int val;
        public Node prev;
        public Node next;
        public Node child;
    };

}
