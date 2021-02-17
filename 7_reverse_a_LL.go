/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverse(curr , prev *ListNode) *ListNode{
    if curr == nil { return prev }
    tmp := curr.Next
    curr.Next = prev
    return reverse(tmp,curr)
}

func reverse2(head, prev *ListNode) *ListNode {
    curr := head
    for curr != nil {
        tmp := curr.Next
        curr.Next = prev
        prev = curr
        curr = tmp
    }
    head = prev

    return head
}


func reverseList(head *ListNode) *ListNode {
    head = reverse2(head, nil);
    return head
    
}

/*


start -> 1 -> 2 -> nil

curr = 1, prev = nil
tmp = curr.next = 2 , curr.next = prev , 1->nil
curr = 2, prev = 1,  2->1->nil


nil <- 1 <- 2 <- start

-> x -> y -> 

 <- x <- y <-  


x,   nx
y,   ny

nx -> y

nx -> prev
tmp = ny
ny -> x


EXPLANATION
-----------

reverse is using recursion
and reverse2 using loops


note that we MUST store the values in a tmp variable because
the node value is modified itself








*/
