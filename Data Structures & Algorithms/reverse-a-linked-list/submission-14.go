/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
    // [0,1,2,3]
    var prev *ListNode

    curr := head
    for curr != nil{
        temp := curr.Next
        curr.Next = prev
        prev = curr
        curr = temp

    }   

    return prev
 
}
