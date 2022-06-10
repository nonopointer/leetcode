package easy._203;

public class Test {
    public static void main(String[] args) {
        
    }

    class ListNode {
        int val;
        ListNode next;

        ListNode() {
        }

        ListNode(int val) {
            this.val = val;
        }

        ListNode(int val, ListNode next) {
            this.val = val;
            this.next = next;
        }
    }

    public ListNode removeElements(ListNode head, int val) {

        ListNode l = new ListNode(0);
        ListNode result = l.next;
        while (head != null) {
            if (head.val == val) {
                l.next = head.next;
                l = l.next;
            }
            head = head.next;
        }

        return result;
    }
}
