package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Node struct {
	Kind   string // P - pair, N = number
	Num    int
	Parent *Node
	Left   *Node
	Right  *Node
}

type Nodes []*Node

func main() {
	start := time.Now()

	data := LoadInput()
	// fmt.Println(data)

	// result := data[0]
	// for i := 1; i < len(data); i++ {
	// 	result = sumTrees(result, data[i])
	// }

	bestMag := -1
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data); j++ {
			result := sumTrees(data[i], data[j])
			m := magnitude(result)
			if m > bestMag {
				bestMag = m
			}
		}
	}
	fmt.Println(bestMag)

	// fmt.Println("-= Input =-")
	// printTrees(data)
	// fmt.Println("-= Sum =-")
	// fmt.Println(treeAsString(result))
	// fmt.Println(magnitude(result))

	elapsed := time.Since(start)
	fmt.Println("done", elapsed)
}

func magnitude(node *Node) int {
	if node.Kind == "N" {
		return node.Num
	}
	return magnitude(node.Left)*3 + magnitude(node.Right)*2
}

func sumTrees(rootNodeL *Node, rootNodeR *Node) *Node {

	nodeL := &Node{}
	copyTree(rootNodeL, nodeL)

	nodeR := &Node{}
	copyTree(rootNodeR, nodeR)

	result := Node{
		Kind:   "P",
		Num:    -1,
		Left:   nodeL,
		Right:  nodeR,
		Parent: nil,
	}
	nodeL.Parent = &result
	nodeR.Parent = &result

	// fmt.Println("-= sum =-", treeAsString(&result))

	for {
		nodeEx := findLeftExplode(&result, 0)

		if nodeEx != nil {
			// fmt.Println("-= explode =-", treeAsString(nodeEx))
			explode(nodeEx)
		} else {
			nodeSp := findLeftSplit(&result)
			if nodeSp != nil {
				// fmt.Println("-= split =-", treeAsString(nodeSp))
				split(nodeSp)
			} else {
				break
			}
		}

		// fmt.Println(treeAsString(&result))
		// fmt.Println(treeAsString(nodeEx))
		// break
	}

	return &result
}

func copyTree(source *Node, empty *Node) {

	empty.Kind = source.Kind
	empty.Num = source.Num

	if source.Kind == "N" {
		empty.Left = nil
		empty.Right = nil
	} else {
		left := &Node{}
		copyTree(source.Left, left)
		empty.Left = left
		left.Parent = empty

		right := &Node{}
		copyTree(source.Right, right)
		empty.Right = right
		right.Parent = empty
	}
}

func split(node *Node) {
	v := node.Num
	lv := math.Floor(float64(v) / 2)
	rv := math.Ceil(float64(v) / 2)

	left := Node{
		Kind:   "N",
		Num:    int(lv),
		Left:   nil,
		Right:  nil,
		Parent: node,
	}

	right := Node{
		Kind:   "N",
		Num:    int(rv),
		Left:   nil,
		Right:  nil,
		Parent: node,
	}

	node.Kind = "P"
	node.Num = -1
	node.Left = &left
	node.Right = &right
}

func findLeftSplit(current *Node) *Node {
	if current.Kind == "N" && current.Num >= 10 {
		return current
	}

	if current.Left != nil {
		left := findLeftSplit(current.Left)
		if left != nil {
			return left
		}
	}

	if current.Right != nil {
		right := findLeftSplit(current.Right)
		if right != nil {
			return right
		}
	}

	return nil
}

func explode(node *Node) {
	lNum := node.Left.Num
	rNum := node.Right.Num
	addToLeft(node, lNum, "U")
	addToRight(node, rNum, "U")

	node.Kind = "N"
	node.Num = 0
	node.Left.Parent = nil
	node.Right.Parent = nil
	node.Left = nil
	node.Right = nil
}

func addToRight(node *Node, num int, dir string) {
	if node.Parent == nil {
		return
	}

	if node.Kind == "N" {
		node.Num = node.Num + num
		return
	}

	if dir == "U" {
		if node.Parent.Right == node {
			addToRight(node.Parent, num, "U")
		}
		if node.Parent.Left == node {
			addToRight(node.Parent.Right, num, "D")
		}
	} else {
		addToRight(node.Left, num, "D")
	}
}

func addToLeft(node *Node, num int, dir string) {
	// fmt.Println("addToLeft", treeAsString(node), num, dir, node.Parent == nil, node.Kind)

	if node.Parent == nil {
		return
	}

	if node.Kind == "N" {
		node.Num = node.Num + num
		return
	}

	if dir == "U" {
		if node.Parent.Left == node {
			addToLeft(node.Parent, num, "U")
		}
		if node.Parent.Right == node {
			addToLeft(node.Parent.Left, num, "D")
		}
	} else {
		addToLeft(node.Right, num, "D")
	}

}

func findLeftExplode(current *Node, level int) *Node {
	if current.Kind == "N" {
		return nil
	}

	if level == 4 && current.Kind == "P" {
		return current
	}

	if current.Left != nil {
		left := findLeftExplode(current.Left, level+1)
		if left != nil {
			return left
		}
	}

	if current.Right != nil {
		right := findLeftExplode(current.Right, level+1)
		if right != nil {
			return right
		}
	}

	return nil
}

func printTrees(data Nodes) {
	for _, node := range data {
		fmt.Println(treeAsString(node))
	}
}

func treeAsString(node *Node) string {
	if node.Kind == "N" {
		return strconv.Itoa(node.Num)
	}

	s1 := treeAsString(node.Left)
	s2 := treeAsString(node.Right)

	// var p string
	// if node.Parent == nil {
	// 	p = "nil"
	// } else if node.Parent.Left == node || node.Parent.Right == node {
	// 	p = "?"
	// } else {
	// 	p = "+"
	// }

	return fmt.Sprintf("[%s,%s]", s1, s2)
}

func LoadInput() Nodes {
	file, _ := os.Open("input.txt")

	result := make(Nodes, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		temp := deepParseNodes(text)
		result = append(result, temp)
	}

	return result
}

func deepParseNodes(text string) *Node {

	if strings.Index(text, ",") == -1 {
		temp, _ := strconv.ParseInt(text, 10, 64)
		return &Node{
			Kind:   "N",
			Num:    int(temp),
			Left:   nil,
			Right:  nil,
			Parent: nil,
		}
	}

	inside := text[1 : len(text)-1]

	var index int
	if inside[0:1] == "[" {
		openBr := 0
		closeBr := 0
		index = 0
		for openBr != closeBr || index == 0 {
			if inside[index:index+1] == "[" {
				openBr++
			}
			if inside[index:index+1] == "]" {
				closeBr++
			}
			index++
		}
	} else {
		index = strings.Index(inside, ",")
	}

	leftText := inside[0:index]
	rightText := inside[index+1:]
	// fmt.Println(text, "L=", leftText, "  R=", rightText)

	leftNode := deepParseNodes(leftText)
	rightNode := deepParseNodes(rightText)

	parent := &Node{
		Kind:   "P",
		Num:    -1,
		Left:   leftNode,
		Right:  rightNode,
		Parent: nil,
	}

	leftNode.Parent = parent
	rightNode.Parent = parent

	return parent
}
