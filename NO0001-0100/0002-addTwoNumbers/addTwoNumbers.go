package addtwonumbers

//2. 两数相加
//给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
//如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
//您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//示例：
//	输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
//	输出：7 -> 0 -> 8
//原因：342 + 465 = 807
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/add-two-numbers
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func ListToArray(list *ListNode) []int {
	tmp := make([]int, 0, 10)
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

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ptr1List := l1
	ptr2List := l2
	carry := false
	var head *ListNode
	var tail *ListNode
	for {
		if ptr1List == nil && ptr2List == nil && !carry {
			return head
		}
		node := new(ListNode)
		if head == nil {
			head = node
			tail = node
		} else {
			tail.Next = node
			tail = tail.Next
		}
		if ptr1List != nil {
			tail.Val += ptr1List.Val
			ptr1List = ptr1List.Next
		}
		if ptr2List != nil {
			tail.Val += ptr2List.Val
			ptr2List = ptr2List.Next
		}
		if carry {
			tail.Val++
			carry = false
		}
		if tail.Val >= 10 {
			tail.Val %= 10
			carry = true
		}
	}
	return head
}

//打印列表
func PrintList(list *ListNode) {
	p := list
	for p != nil {
		//fmt.Printf("%d-%v-%p ", p.Val, p, p.Next)
		fmt.Printf("%#v", p.Val)
		p = p.Next
	}
	fmt.Println()
}

//func main() {
//num1head := ArrayToList([]int{2, 4, 3})
//num2head := ArrayToList([]int{5, 6, 4})
//PrintList(num1head)
//PrintList(num2head)
//addresulthead := AddTwoNumbers(num1head, num2head)
//PrintList(addresulthead)
//}

//space: O(max(m,n)
//time : O(m+n)
