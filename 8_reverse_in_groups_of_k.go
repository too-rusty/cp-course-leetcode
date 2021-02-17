/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func f1(prev, head *ListNode, k int) *ListNode {
    node_cnt := 0
    var tmp_start *ListNode = head
    for tmp_start != nil && node_cnt < k {
        tmp_start = tmp_start.Next
        node_cnt++
    }
    if node_cnt < k { return head }

    cnt, curr := k, head
    
    for cnt > 0 {
        tmp := curr.Next
        curr.Next = prev
        prev = curr
        curr = tmp
        cnt--
    }
    head.Next = f1(nil, curr, k)
    head = prev
    return head
    
}


func reverseKGroup(head *ListNode, k int) *ListNode {
    return f1(nil, head, k)
}

/*

here we reverse the nodes in groups of k
this is a combination of loops and recursion
recursively modify the linked list in groups of k
and how do we actually modify them, using loops of course


*/
