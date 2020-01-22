    package main

    import (
    	"fmt"
    )

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

    // 	BTreeApplyByLevel(root, fmt.Println)
    // }

    // func main() {
    // 	root := &TreeNode{Data: "01"}
    // 	BTreeInsertData(root, "07")
    // 	BTreeInsertData(root, "05")
    // 	BTreeInsertData(root, "12")
    // 	BTreeInsertData(root, "02")
    // 	BTreeInsertData(root, "10")
    // 	BTreeInsertData(root, "03")

    // 	BTreeApplyByLevel(root, fmt.Print)
    // }

    func NextL(root *TreeNode, f func(...interface{}) (int, error), lv int) {
    	if root == nil {
    		return
    	}
    	if lv == 0 {
    		f(root.Data)
    	} else if lv > 0 {
    		NextL(root.Left, f, lv-1)
    		NextL(root.Right, f, lv-1)
    	}
    }

    func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
    	l := BTreeLevelCount(root)
    	for i := 0; i < l; i++ {
    		NextL(root, f, i)
    	}
    }
    func BTreeLevelCount(root *TreeNode) int {
    	if root == nil {
    		return 0
    	}
    	cl := 0
    	cr := 0
    	if root.Left != nil {
    		cr = BTreeLevelCount(root.Left)
    	}
    	if root.Right != nil {
    		cl = BTreeLevelCount(root.Right)
    	}
    	if cl > cr {
    		cr = cl
    	}
    	return cr + 1
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