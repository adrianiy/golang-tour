// Package tree contains tree building exercise.
package tree

import "errors"

type Record struct {
	ID     int
	Parent int
	// feel free to add fields as you see fit
}

type Node struct {
	ID       int
	Children []*Node
	used     bool
	// feel free to add fields as you see fit
}

// AddChild concats child in correct position sorting by id
func (n *Node) AddChild(child *Node) {
	if len(n.Children) == 0 || n.Children[0].ID < child.ID {
		n.Children = append(n.Children, child)
	} else {
		n.Children = append([]*Node{child}, n.Children...)
	}
}

// Build constructs a tree from records slice.
func Build(records []Record) (*Node, error) {
	var root *Node
	nodeMap := make(map[int]*Node)
	n := len(records)
	sum := (n * (n - 1)) / 2

	for _, r := range records {
		sum -= r.ID
		if r.ID < 0 || r.ID > len(records) {
			return nil, errors.New("Invalid ID")
		}

		var node *Node = nodeMap[r.ID]

		if node == nil {
			node = &Node{ID: r.ID, Children: []*Node{}, used: true}
			nodeMap[r.ID] = node
		} else if node.used {
			return nil, errors.New("Duplicated node")
		}

		if r.ID == 0 {
			if r.Parent != 0 {
				return nil, errors.New("Invalid root")
			}
			root = node
		} else {
			if r.ID == r.Parent {
				return nil, errors.New("Cycle directly")
			}
			if r.Parent > r.ID {
				return nil, errors.New("Higher parent ID")
			}
			var parent *Node = nodeMap[r.Parent]

			if parent != nil {
				parent.AddChild(node)
			} else {
				nodeMap[r.Parent] = &Node{
					ID:       r.Parent,
					Children: []*Node{node},
				}
			}
		}
	}
	if sum != 0 {
		return nil, errors.New("Non continouos")
	}

	if len(records) > 0 && root == nil {
		return nil, errors.New("No root node")
	}

	return root, nil
}
