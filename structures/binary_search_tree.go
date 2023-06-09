package structures

// Node
type NodeBST struct {
	Key   int
	Left  *NodeBST
	Right *NodeBST
}

func (n *NodeBST) Insert(k int) {
	if n.Key < k {
		if n.Right == nil {
			n.Right = &NodeBST{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		if n.Left == nil {
			n.Left = &NodeBST{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

// Insert
// Search
