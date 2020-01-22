package main

import "fmt"

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

// func main() {
// 	root := &TreeNode{Data: "01"}
// 	BTreeInsertData(root, "07")
// 	BTreeInsertData(root, "05")
// 	BTreeInsertData(root, "12")
// 	BTreeInsertData(root, "02")
// 	BTreeInsertData(root, "10")
// 	BTreeInsertData(root, "03")

// 	node := BTreeSearchItem(root, "02")
// 	fmt.Println("Before delete:")
// 	BTreeApplyInorder(root, fmt.Println)
// 	root = BTreeDeleteNode(root, node)
// 	fmt.Println("After delete:")
// 	BTreeApplyInorder(root, fmt.Println)
// }

// func main() {
// 	root := &TreeNode{Data: "33"}
// 	BTreeInsertData(root, "20")
// 	BTreeInsertData(root, "5")
// 	BTreeInsertData(root, "13")
// 	BTreeInsertData(root, "31")
// 	BTreeInsertData(root, "52")
// 	BTreeInsertData(root, "11")

// 	node := BTreeSearchItem(root, "20")
// 	fmt.Println("Before delete:")
// 	BTreeApplyInorder(root, fmt.Println)
// 	root = BTreeDeleteNode(root, node)
// 	fmt.Println("After delete:")
// 	BTreeApplyInorder(root, fmt.Println)
// }

// func main() {
// 	root := &TreeNode{Data: "03"}
// 	BTreeInsertData(root, "39")
// 	BTreeInsertData(root, "11")
// 	BTreeInsertData(root, "99")
// 	BTreeInsertData(root, "14")
// 	BTreeInsertData(root, "44")
// 	BTreeInsertData(root, "11")

// 	node := BTreeSearchItem(root, "03")
// 	fmt.Println("Before delete:")
// 	BTreeApplyInorder(root, fmt.Println)
// 	root = BTreeDeleteNode(root, node)
// 	fmt.Println("After delete:")
// 	BTreeApplyInorder(root, fmt.Println)
// }

// func main() {
// 	root := &TreeNode{Data: "03"}
// 	BTreeInsertData(root, "01")
// 	BTreeInsertData(root, "03")
// 	BTreeInsertData(root, "94")
// 	BTreeInsertData(root, "19")
// 	BTreeInsertData(root, "111")
// 	BTreeInsertData(root, "24")

// 	node := BTreeSearchItem(root, "03")
// 	fmt.Println("Before delete:")
// 	BTreeApplyInorder(root, fmt.Println)
// 	root = BTreeDeleteNode(root, node)
// 	fmt.Println("After delete:")
// 	BTreeApplyInorder(root, fmt.Println)
// }

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	headparent := node.Parent
	flag := false
	if headparent != nil && headparent.Data <= node.Data {

		flag = true
	}
	if node.Right != nil {
		BTreeMin(node.Right).Left = node.Left
		node = node.Right
	} else {

		node = node.Left
	}
	if headparent != nil {
		if flag {
			headparent.Right = node
		} else {

			headparent.Left = node
		}
	}
	if node != nil {
		node.Parent = headparent
	}
	if node != nil && node.Parent == nil {
		root = node
	}
	return root
}

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {

	if root == nil {
		return nil
	}
	if elem < root.Data {
		root = BTreeSearchItem(root.Left, elem)
	} else if elem == root.Data {

		return root

	} else {
		root = BTreeSearchItem(root.Right, elem)
	}

	return root
}
func BTreeMin(root *TreeNode) *TreeNode {

	if root == nil {

		return nil
	}
	if root.Left != nil {

		root = BTreeMin(root.Left)
	}

	return root
}
func BTreeInsertData(root *TreeNode, data string) *TreeNode {

	if root == nil {
		return &TreeNode{Data: data}
	}
	if data < root.Data {
		root.Left = BTreeInsertData(root.Left, data)
		root.Left.Parent = root
	} else {
		root.Right = BTreeInsertData(root.Right, data)
		root.Right.Parent = root
	}
	return root

}
func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root.Left != nil {
		BTreeApplyInorder(root.Left, f)
	}
	f(root.Data)
	if root.Right != nil {
		BTreeApplyInorder(root.Right, f)
	} else {
		return
	}
}