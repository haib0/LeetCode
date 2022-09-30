/*
 * @lc app=leetcode.cn id=21 lang=rust
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
// Definition for singly-linked list.
// #[derive(PartialEq, Eq, Clone, Debug)]
// pub struct ListNode {
//   pub val: i32,
//   pub next: Option<Box<ListNode>>
// }
//
// impl ListNode {
//   #[inline]
//   fn new(val: i32) -> Self {
//     ListNode {
//       next: None,
//       val
//     }
//   }
// }
impl Solution {
    pub fn merge_two_lists(
        l1: Option<Box<ListNode>>,
        l2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        let mut r: Option<Box<ListNode>> = None;
        let mut r_p = &mut r;
        let mut l1_p = &l1;
        let mut l2_p = &l2;
        while let (Some(l1_b),Some(l2_b)) = (l1_p, l2_p) {
            if l1_b.val <= l2_b.val{
                if let Some(r_p_b) = r_p {
                    r_p_b.next =Some(Box::new(ListNode::new(l1_b.val)));
                    r_p = &mut r_p_b.next;
                }else {
                    // first time
                    r =  Some(Box::new(ListNode::new(l1_b.val)));
                    r_p = &mut r;
                }
                l1_p = &l1_b.next;
            } else {
                if let Some(r_p_b) = r_p {
                    r_p_b.next =Some(Box::new(ListNode::new(l2_b.val)));
                    r_p = &mut r_p_b.next;
                }else {
                    // first time
                    r =  Some(Box::new(ListNode::new(l2_b.val)));
                    r_p = &mut r;
                }
                l2_p = &l2_b.next;
            }
        }
        while let Some(l1_b) = l1_p {
            if let Some(r_p_b) = r_p {
                r_p_b.next =Some(Box::new(ListNode::new(l1_b.val)));
                r_p = &mut r_p_b.next;
            }else {
                // first time
                r =  Some(Box::new(ListNode::new(l1_b.val)));
                r_p = &mut r;
            }
            l1_p = &l1_b.next;
        }
    
        while let Some(l2_b) = l2_p {
            if let Some(r_p_b) = r_p {
                r_p_b.next =Some(Box::new(ListNode::new(l2_b.val)));
                r_p = &mut r_p_b.next;
            }else {
                // first time
                r =  Some(Box::new(ListNode::new(l2_b.val)));
                r_p = &mut r;
            }
            l2_p = &l2_b.next;
        }
    
        r
    }
}
// @lc code=end

