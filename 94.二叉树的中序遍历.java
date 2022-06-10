import java.util.ArrayList;
import java.util.List;

/*
 * @lc app=leetcode.cn id=94 lang=java
 *
 * [94] 二叉树的中序遍历
 */

// @lc code=start

 // Definition for a binary tree node.
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

class Solution1 {
    
    public List<Integer> inorderTraversal(TreeNode root) {
        List<Integer> list = new ArrayList<>();
        if(root == null) return list;
        fx(root,list);
        return list;
    }

    private void fx(TreeNode root, List<Integer> list) {
        if(root.left != null) fx(root.left,list);
        list.add(root.val);
        if(root.right != null) fx(root.right,list);
    }

}
// @lc code=end

