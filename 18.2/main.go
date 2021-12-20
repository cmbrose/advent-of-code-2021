package main

import (
	"fmt"

	"../util"
)

type Node struct {
	Value int

	Left   *Node
	Right  *Node
	Parent *Node
}

func (n Node) String() string {
	if n.IsLeaf() {
		return fmt.Sprintf("%d", n.Value)
	}

	return fmt.Sprintf("[%s,%s]", n.Left.String(), n.Right.String())
}

func (n Node) IsLeaf() bool {
	return n.Left == nil // Both left and right are also set a the same time
}

func (n *Node) IsLeftChild() bool {
	return n.Parent != nil && n.Parent.Left == n
}

func (n *Node) IsRightChild() bool {
	return n.Parent != nil && n.Parent.Right == n
}

func (n Node) IsRoot() bool {
	return n.Parent == nil
}

func (n Node) Magnitude() int {
	if n.IsLeaf() {
		return n.Value
	}

	return n.Left.Magnitude()*3 + n.Right.Magnitude()*2
}

func (n *Node) FindExplodeCandidate(depth int) *Node {
	if n.IsLeaf() {
		return nil
	}

	if n.Left.IsLeaf() && n.Right.IsLeaf() && depth >= 4 {
		return n
	}

	left := n.Left.FindExplodeCandidate(depth + 1)
	if left != nil {
		return left
	}

	right := n.Right.FindExplodeCandidate(depth + 1)
	if right != nil {
		return right
	}

	return nil
}

func (n *Node) Explode() {
	n.Value = 0 // This should have already been the case??

	curr := n.Left
	for curr.IsLeftChild() {
		curr = curr.Parent
	}
	if !curr.IsRoot() {
		curr = curr.Parent.Left
		for !curr.IsLeaf() {
			curr = curr.Right
		}
		curr.Value += n.Left.Value
	}
	n.Left = nil

	curr = n.Right
	for curr.IsRightChild() {
		curr = curr.Parent
	}
	if !curr.IsRoot() {
		curr = curr.Parent.Right
		for !curr.IsLeaf() {
			curr = curr.Left
		}
		curr.Value += n.Right.Value
	}
	n.Right = nil
}

func (n *Node) FindSplitCandidate() *Node {
	if n.IsLeaf() {
		if n.Value >= 10 {
			return n
		}
		return nil
	}

	left := n.Left.FindSplitCandidate()
	if left != nil {
		return left
	}

	right := n.Right.FindSplitCandidate()
	if right != nil {
		return right
	}

	return nil
}

func (n *Node) Split() {
	leftVal := n.Value / 2
	rightVal := n.Value - leftVal

	n.Left = &Node{Value: leftVal, Parent: n}
	n.Right = &Node{Value: rightVal, Parent: n}

	n.Value = 0
}

func (n *Node) AddAndReduce(other *Node) *Node {
	temp := &Node{Left: n, Right: other}
	n.Parent = temp
	other.Parent = temp

	for {
		explode := temp.FindExplodeCandidate(0)
		if explode != nil {
			explode.Explode()
			continue
		}

		split := temp.FindSplitCandidate()
		if split != nil {
			split.Split()
			continue
		}

		break
	}

	return temp
}

type Scanner struct {
	text string
}

func (s Scanner) Peek() rune {
	return rune(s.text[0])
}

func (s *Scanner) Next() rune {
	r := rune(s.text[0])
	s.text = s.text[1:]
	return r
}

func (s *Scanner) AssertNext(r rune) {
	if s.Peek() != r {
		panic(fmt.Sprintf("Next char should be %c, but was %c", r, s.Peek()))
	}

	s.Next()
}

func (s *Scanner) NextInt() int {
	val := 0

	for s.Peek() >= '0' && s.Peek() <= '9' {
		r := s.Next()
		val = val*10 + int(r-'0')
	}

	return val
}

func main() {
	lines := util.ReadInputLines("./input.txt")

	maxMag := 0

	for i, leftExpr := range lines {
		for j, rightExpr := range lines {
			if i == j {
				continue
			}

			// The addition changes the tree structure
			// so we can't just parse once and reuse.
			left := parseExpression(leftExpr)
			right := parseExpression(rightExpr)

			result := left.AddAndReduce(right)
			mag := result.Magnitude()
			if mag > maxMag {
				maxMag = mag
			}
		}
	}

	fmt.Println(maxMag)
}

func parseExpression(expr string) *Node {
	scanner := &Scanner{expr}
	return parseExpressionWithScanner(scanner)
}

func parseExpressionWithScanner(s *Scanner) *Node {
	n := &Node{}

	if s.Peek() == '[' {
		s.Next()

		n.Left = parseExpressionWithScanner(s)
		n.Left.Parent = n

		s.AssertNext(',')

		n.Right = parseExpressionWithScanner(s)
		n.Right.Parent = n

		s.AssertNext(']')
	} else {
		n.Value = s.NextInt()
	}

	return n
}
