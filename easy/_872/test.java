package easy._872;

import java.util.LinkedList;
import java.util.List;

public class test {
    class TreeNode {
             int val;
             TreeNode left;
             TreeNode right;
             TreeNode() {}
             TreeNode(int val) { this.val = val; }
             TreeNode(int val, TreeNode left, TreeNode right) {
                 this.val = val;
                 this.left = left;
                 this.right = right;
             }
         }
    public static void main(String[] args){
        List<Integer> l = new LinkedList<>();
        l.add(1);
        l.add(2);
        l.add(1);
        l.add(2);
        l.add(1);
        l.add(2);
        l.remove(0);
        System.out.println(l);
        System.out.println(l.size());
    }
    static List<Integer> l1 = new LinkedList<>();
    static List<Integer> l2 = new LinkedList<>();
    static List<Integer> l = new LinkedList<>();
    public boolean leafSimilar(TreeNode root1, TreeNode root2) {
        searchAdd(l1,root1);
        searchAdd(l2,root2);
        if(l1.size()!=l2.size()) return false;
        for(int i = 0;i < l1.size();i++){
            if(l1.remove(0)!=l2.remove(0)) return false;
        }
        System.out.println(l1);
        return true;
    }
	private void searchAdd(List<Integer> l, TreeNode node) {
        if (node.left != null) searchAdd(l, node.left);
        if (node.right != null ) searchAdd(l,node.right);
        if(node.left == null && node.right == null)l.add(node.val);
	}

    // public boolean leafSimilar(TreeNode root1, TreeNode root2) {
    //     searchAdd(root1);
    //     searchAdd(root2);
    //     if(l1.size()!=l2.size()) return false;
    //     for(int i = 0;i < l1.size();i++){
    //         if(l1.remove(0)!=l2.remove(0)) return false;
    //     }
    //     return true;
    // }
	// private void searchAdd(TreeNode node) {
    //     l.add(node.val);
    //     if (node.left != null){
    //         l.add(node.left.val);
    //     } 
    //     if (node.right != null ) {
    //         l.add(node.right.val);
    //     }
            
    //     l1.add(node.val);
	// }
}