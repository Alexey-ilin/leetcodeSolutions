package solutions

// https://leetcode.com/problems/reverse-odd-levels-of-binary-tree/

import (
	"errors"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Simple queue implementation
type queue[T any] struct {
	elements []T
}

func (q *queue[T]) Push(newNode T) {
	q.elements = append(q.elements, newNode)
}

func (q *queue[T]) Get() (T, error) {
	var t T
	if len(q.elements) == 0 {
		return t, errors.New("queue is empty")
	}
	t = q.elements[0]
	q.elements = q.elements[1:]
	return t, nil
}

func (q queue[T]) Peek() (T, error) {
	var t T
	if len(q.elements) == 0 {
		return t, errors.New("queue is empty")
	}
	return q.elements[0], nil
}

// Queue constructor
func NewQueue[T any](length int) queue[T] {
	return queue[T]{elements: make([]T, 0, length)}
}

// Main functions

func ReverseOddLevels(root *TreeNode) *TreeNode {
	reverseBfs(root)
	return root
}

// DFS solution
func reverseDfs(root *TreeNode) {
	dfsHelper(root.Left, root.Right, 0)
}

func dfsHelper(left, right *TreeNode, level int) {
	if left == nil && right == nil {
		return
	}
	if level%2 == 0 {
		left.Val, right.Val = right.Val, left.Val
	}
	dfsHelper(left.Left, right.Right, level+1)
	dfsHelper(left.Right, right.Left, level+1)
}

// BFS solution
func reverseBfs(root *TreeNode) {
	bfsHelper(root)
}

func bfsHelper(root *TreeNode) {
	level := 0
	fmt.Printf("Level %v:\n", level)
	queue := NewQueue[*TreeNode](10)
	queue.Push(root)
	queue.Push(nil)
	for len(queue.elements) > 1 {
		n, _ := queue.Get()
		// keep tracking of depth of tree
		if n == nil {
			level++
			// reversing odd levels
			if level%2 != 0 {
				for i := 0; i < (len(queue.elements))/2; i++ {
					queue.elements[i].Val, queue.elements[len(queue.elements)-1-i].Val =
						queue.elements[len(queue.elements)-1-i].Val, queue.elements[i].Val
				}
			}
			queue.Push(nil)
			if endCheck, _ := queue.Peek(); endCheck == nil {
				return
			}
			fmt.Printf("\nLevel %v:\n", level)
			continue
		}
		fmt.Printf("%v ", n.Val)
		if n.Left != nil {
			queue.Push(n.Left)
		}
		if n.Right != nil {
			queue.Push(n.Right)
		}
	}
}
