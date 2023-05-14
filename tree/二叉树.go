package tree

import "fmt"

type BinaryTree struct {
	root *HeroNode
}

func (b *BinaryTree) preOrder() {
	if b.root != nil {
		b.root.preOrder()
	}
}

func (b *BinaryTree) midOrder() {
	if b.root != nil {
		b.root.midOrder()
	}
}

func (b *BinaryTree) postOrder() {
	if b.root != nil {
		b.root.postOrder()
	}
}

func (b *BinaryTree) preOrderSearch(no int) *HeroNode {
	if b.root == nil {
		return nil
	}

	return b.root.preOrderSearch(no)
}

func (b *BinaryTree) midOrderSearch(no int) *HeroNode {
	if b.root == nil {
		return nil
	}

	return b.root.midOrderSearch(no)
}

func (b *BinaryTree) postOrderSearch(no int) *HeroNode {
	if b.root == nil {
		return nil
	}

	return b.root.postOrderSearch(no)
}

func (b *BinaryTree) delNode(no int)  {
	if b.root != nil {
		//一定要判断root节点是否要删除的节点
		if b.root.no == no {
			b.root = nil
		}else {
			b.root.delNode(no)
		}
	}else {
		fmt.Println("空树，不能删除")
	}
}

type HeroNode struct {
	no    int
	name  string
	left  *HeroNode
	right *HeroNode
}

func NewHeroNode(no int, name string) *HeroNode {
	return &HeroNode{
		no:   no,
		name: name,
	}
}

func (h *HeroNode) printf() {
	s := fmt.Sprintf("HeroNode no:%d, name=%s", h.no, h.name)
	fmt.Println(s)
}

//前序遍历
func (h *HeroNode) preOrder() {
	h.printf()
	//递归遍历左节点
	if h.left != nil {
		h.left.preOrder()
	}
	//递归遍历右节点
	if h.right != nil {
		h.right.preOrder()
	}
}

//中序遍历
func (h *HeroNode) midOrder() {
	//递归遍历左节点
	if h.left != nil {
		h.left.midOrder()
	}
	h.printf()
	//递归遍历右节点
	if h.right != nil {
		h.right.midOrder()
	}
}

//后序遍历
func (h *HeroNode) postOrder() {
	//递归遍历左节点
	if h.left != nil {
		h.left.postOrder()
	}

	//递归遍历右节点
	if h.right != nil {
		h.right.postOrder()
	}

	h.printf()
}

//前序节点查找
func (h *HeroNode) preOrderSearch(no int) *HeroNode {
	//判断当前节点是不是
	if h.no == no {
		return h
	}

	//判断左子树是否为nil，进行前序递归查找
	//如果左递归前序查找，找到节点，则返回
	var resNode *HeroNode
	if h.left != nil {
		resNode = h.left.preOrderSearch(no)
	}
	if resNode != nil {
		return resNode
	}

	//同样向右递归前序查找
	if h.right != nil {
		resNode = h.right.preOrderSearch(no)
	}

	return resNode
}

//中序节点查找
func (h *HeroNode) midOrderSearch(no int) *HeroNode {
	//判断左子树是否为nil，进行递归查找
	//如果左递归查找，找到节点，则返回
	var resNode *HeroNode
	if h.left != nil {
		resNode = h.left.midOrderSearch(no)
	}
	if resNode != nil {
		return resNode
	}

	//判断当前节点是不是
	if h.no == no {
		return h
	}

	//同样向右递归查找
	if h.right != nil {
		resNode = h.right.midOrderSearch(no)
	}

	return resNode
}

//后序节点查找
func (h *HeroNode) postOrderSearch(no int) *HeroNode {
	//判断左子树是否为nil，进行递归查找
	//如果左递归查找，找到节点，则返回
	var resNode *HeroNode
	if h.left != nil {
		resNode = h.left.midOrderSearch(no)
	}
	if resNode != nil {
		return resNode
	}

	//同样向右递归查找
	if h.right != nil {
		resNode = h.right.midOrderSearch(no)
	}
	if resNode != nil {
		return resNode
	}

	//判断当前节点是不是
	if h.no == no {
		return h
	}
	return resNode
}

//删除节点，二叉树遍历是单向的，所以要判断当前节点的左子节点和右子节点是否为要删除的节点，没有则进行递归
func (h *HeroNode) delNode(no int)     {
	//一定要查找当前节点的左子树，而不是移动到左子树在判断
	if h.left != nil && h.left.no== no{
		h.left = nil
		return
	}
	if h.right != nil && h.right.no== no{
		h.right = nil
		return
	}

	if h.left != nil {
		h.left.delNode(no)
	}

	if h.right != nil {
		h.right.delNode(no)
	}
}