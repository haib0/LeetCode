/*
 * @lc app=leetcode.cn id=2 lang=rust
 *
 * [2] 两数相加
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
    pub fn add_two_numbers(
        l1: Option<Box<ListNode>>,
        l2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        let mut result = None;
        let mut tail = &mut result;
        // (l1, l2, carry, sum)
        let mut t = (l1, l2, 0, 0);
        loop {
            t = match t {
                (Some(list1), Some(list2), carry, _) if list1.val + list2.val + carry < 10 => (
                    list1.next,
                    list2.next,
                    0,
                    list1.val + list2.val + carry,
                ),

                (Some(list1), Some(list2), carry, _) => (
                    list1.next,
                    list2.next,
                    1,
                    list1.val + list2.val + carry - 10,
                ),

                (Some(list), None, carry, _) | (None, Some(list), carry, _)
                    if list.val + carry < 10 =>
                {
                    (list.next, None, 0, list.val + carry)
                }

                (Some(list), None, carry, _) | (None, Some(list), carry, _) => {
                    (list.next, None, 1, list.val + carry - 10)
                }

                (None, None, 0, _) => {
                    break;
                }

                (None, None, carry, _) => (None, None, 0, carry),
            };

            *tail = Some(Box::new(ListNode::new(t.3)));
            tail = &mut tail.as_mut().unwrap().next;
        }
        result
    }
}
// @lc code=end
