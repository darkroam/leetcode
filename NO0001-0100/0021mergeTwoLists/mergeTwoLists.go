package mergeTwoLists

//21. 合并两个有序链表
//将两个升序链表合并为一个新的升序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//示例：
//	输入：1->2->4, 1->3->4
//	输出：1->1->2->3->4->4
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/add-two-numbers
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

import "fmt"

func PrintList(list *ListNode) {
	p := list
	for p != nil {
		//fmt.Printf("%d-%v-%p ", p.Val, p, p.Next)
		fmt.Printf("%#v", p.Val)
		p = p.Next
	}
	fmt.Println()
}

func ListToArray(list *ListNode) []int {
	tmp := make([]int, 0, 10)
	if list == nil {
		return tmp
	}
	for list.Next != nil {
		tmp = append(tmp, list.Val)
		list = list.Next
	}
	tmp = append(tmp, list.Val)
	return tmp
}

func ArrayToList(nums []int) *ListNode {
	if len(nums) <= 0 {
		return nil
	}
	head := &ListNode{
		Val:  nums[0],
		Next: nil,
	}
	tail := head
	for _, v := range nums[1:] {
		node := new(ListNode)
		tail.Next = node
		tail = tail.Next
		tail.Val = v
	}
	return head
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			cur, l1 = cur.Next, l1.Next
		} else {
			cur.Next = l2
			cur, l2 = cur.Next, l2.Next
		}
	}
	if l1 == nil {
		cur.Next = l2
	} else {
		cur.Next = l1
	}
	return head.Next
}

//func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
//if l1 == nil {
//return l2
//}
//if l2 == nil {
//return l1
//}
//if l1.Val < l2.Val {
//l1.Next = MergeTwoLists(l1.Next, l2)
//return l1
//} else {
//l2.Next = MergeTwoLists(l1, l2.Next)
//return l2
//}
//}
//space: O(max(m,n)
//time : O(m+n)
