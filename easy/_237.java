package easy;

public class _237 {
    public static void main(String[] args) {

    }

    public static void deleteNode(ListNode node) {
        ListNode nextNode = node.next;
        ListNode nextNextNode = nextNode.next;
        node.val = nextNode.val;
        node.next = nextNextNode;
    }

    public class ListNode {
        int val;
        ListNode next;

        ListNode(int x) {
            val = x;
        }
    }
}
