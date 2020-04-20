package swapPairs

//24. 两两交换链表中的节点
//给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
//你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
//示例:
//	给定 1->2->3->4, 你应该返回 2->1->4->3.
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

func SwapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	cur := &ListNode{}
	tmp, res := head, head.Next

	for head != nil && head.Next != nil {
		tmp = head.Next.Next
		cur.Next = head.Next
		cur.Next.Next = head
		cur = cur.Next.Next
		head = tmp
	}

	if head != nil {
		cur.Next = head
		cur.Next.Next = nil
		cur = cur.Next
	}
	cur.Next = nil //必须守收尾，否则会最后两个会成环
	return res

}

//space: O(?)
//time : O(?)
