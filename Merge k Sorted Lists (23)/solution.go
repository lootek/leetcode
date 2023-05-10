/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    var resHead, resIter *ListNode

    for {
        min := 100_000
        min_i := -1
        for i := range lists {
            if lists[i] != nil && lists[i].Val < min {
                min_i = i
                min = lists[i].Val
            }
        }

        // nothing found so we're done
        if min == 100_000 {
            break
        }

        lists[min_i] = lists[min_i].Next
        newNode := &ListNode{min, nil}
        if resHead == nil {
            resHead = newNode
            resIter = resHead
        } else {
            resIter.Next = newNode
            resIter = newNode
        }
    }

    return resHead
}