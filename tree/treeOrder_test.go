package tree

import (
	"fmt"
	"testing"
)

func TestNewBinaryTree(t *testing.T)  {
	bt := new(BinaryTree)
	root := NewHeroNode(1, "宋江")
	node2 := NewHeroNode(2, "吴用")
	node3 := NewHeroNode(3, "卢俊义")
	node4 := NewHeroNode(4, "林冲")
	node5 := NewHeroNode(5, "关胜")

	root.left = node2
	root.right = node3
	node3.right = node4
	node3.left = node5
	bt.root = root
	fmt.Println("前序遍历:")
	bt.preOrder()

	fmt.Println("中序遍历:")
	bt.midOrder()

	fmt.Println("后序遍历:")
	bt.postOrder()

	fmt.Println("前序遍历查找:")
	var resNode *HeroNode
	resNode= bt.preOrderSearch(5)
	if resNode != nil {
		fmt.Println("前序找到node:", resNode.no, resNode.name)
	}else {
		fmt.Println("前序没有查找")
	}


	fmt.Println("中序遍历查找:")
	resNode= bt.midOrderSearch(5)
	if resNode != nil {
		fmt.Println("中序找到node:", resNode.no, resNode.name)
	}else {
		fmt.Println("中序没有查找")
	}


	fmt.Println("后序遍历查找:")
	resNode= bt.postOrderSearch(15)
	if resNode != nil {
		fmt.Println("后序找到node:", resNode.no, resNode.name)
	}else {
		fmt.Println("后序没有查找")
	}

}
