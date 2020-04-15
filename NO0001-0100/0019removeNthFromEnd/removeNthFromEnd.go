package removeNthFromEnd

//19. 删除链表的倒数第N个节点
//给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
//示例：
//	给定一个链表: 1->2->3->4->5, 和 n = 2.
//	当删除了倒数第二个节点后，链表变为 1->2->3->5.
//说明：
//	给定的 n 保证是有效的。
//进阶：
//	你能尝试使用一趟扫描实现吗？
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

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil && n == 1 {
		return nil
	}
	if head.Next.Next == nil && n == 1 {
		head.Next = nil
		return head
	}
	lptr, rptr := head, head
	for i := 0; i < n && rptr != nil; i++ {
		rptr = rptr.Next
	}
	if rptr == nil {
		return head.Next
	}
	for rptr.Next != nil {
		rptr = rptr.Next
		lptr = lptr.Next
	}
	lptr.Next = lptr.Next.Next
	return head
}

//space: O(max(m,n)
//time : O(m+n)
