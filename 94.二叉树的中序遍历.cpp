/*
 * @lc app=leetcode.cn id=94 lang=cpp
 *
 * [94] 二叉树的中序遍历
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
#include <vector>
using namespace std;
class Solution
{
public:
    struct TreeNode
    {
        int val;
        TreeNode *left;
        TreeNode *right;
        TreeNode() : val(0), left(nullptr), right(nullptr) {}
        TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
        TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
    };

    void inorderT(TreeNode *node, vector<int> *ans)
    {
        if (node->left)
        {
            inorderT(node->left, ans);
        }

        ans->push_back(node->val);

        if (node->right)
        {
            inorderT(node->right, ans);
        }
    }

    vector<int> inorderTraversal(TreeNode *root)
    {
        vector<int> *ans = new vector<int>;
        if (root)
        {
            inorderT(root, ans);
        }
        return *ans;
    }
};
// @lc code=end
