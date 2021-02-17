/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


import (
	"container/heap"
)


type Node struct {
    Val int
    Pointer *ListNode
}

// A NodeHeap is a min-heap of Nodes.
type NodeHeap []Node

func (h NodeHeap) Len() int           { return len(h) }
func (h NodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h NodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(Node))
}
func (h *NodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mergeKLists(pointers []*ListNode) *ListNode {
    var final_merged_LL *ListNode = nil
    
    h := &NodeHeap{}
    heap.Init(h)
    
    for _, pointer := range pointers {
        if pointer != nil {
            nodeVal := Node { pointer.Val , pointer }
            heap.Push(h, nodeVal)
        }
    }
    
    var curr *ListNode = nil
    
    for h.Len() > 0 {
        topVal := heap.Pop(h).(Node)
        if final_merged_LL == nil {
            final_merged_LL = new(ListNode)
            final_merged_LL.Val = topVal.Val
            // final_merged_LL.Next = nil
            curr = final_merged_LL
        } else {
            curr.Next = new(ListNode)
            curr = curr.Next
            curr.Val = topVal.Val
        }
        topVal.Pointer = topVal.Pointer.Next
        if topVal.Pointer != nil {
            nodeVal := Node { topVal.Pointer.Val , topVal.Pointer }
            heap.Push(h,nodeVal)
        }
    }
    
    
    return final_merged_LL
}

/*

k sorted linked lists

l1 , l2 , l3, l4, lk

l1 + l2 + l3 + l4 .... + lk



1e4 * 500 = 5e6 max total len

here we merge k sorted linked lists
we make a min heap on the basis of the list values
and eliminate the elements one by one from the heap


*/
