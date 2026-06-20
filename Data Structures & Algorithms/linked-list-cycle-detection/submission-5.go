/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
    // hashMap := make(map[*ListNode]bool)
    // for head != nil{
    //     if hashMap[head]{
    //         return true
    //     }
    //     hashMap[head] = true
    //     head = head.Next
    // }

    slow, fast := head, head

    for fast != nil && fast.Next != nil{
        slow = slow.Next
        fast = fast.Next.Next

        if slow == fast {
            return true
        }
        head = head.Next
    }

    return false
}
