package avl

import (
	"errors"

	"github.com/go-qiu/rrs-web-server/pkg/ds/stack"
)

/*
	avl tree struct
*/
type AVL struct {
	root *BinaryNode
}

var (
	ErrEmptyTree             error = errors.New("[AVL]: tree is empty")
	ErrNodeNotFound          error = errors.New("[AVL]: node not found")
	ErrDuplicatedNode        error = errors.New("[AVL]: found duplicated parcel job record")
	ErrEmptyNodeItemStatus   error = errors.New("[AVL]: item status code cannot be empty")
	ErrInvalidNodeItemStatus error = errors.New("[AVL]: item status code is invalid")
)

/*
	Wrapper function to instantiate an AVL tree (in-memory)
*/
func New() *AVL {
	return &AVL{root: nil}
}

/*
	Wrapper function to insert a new node into the AVL tree.
*/
func (tree *AVL) InsertNode(item interface{}, id string) error {

	root, err := insertNode(tree.root, item, id)
	if err != nil {
		return err
	}

	tree.root = root
	return nil
}

func (tree *AVL) ListAllNodes(s *stack.Stack) error {

	// ok. tree is not empty

	// traverse all the nodes on the tree (in-order sequence)
	// and cache the outcome into the stack (passed in).
	traverse(tree.root, s)
	return nil
}

/*
	Wrapper function to find a specific node by id
*/
func (tree *AVL) Find(id string) *BinaryNode {

	return findNode(tree.root, id)
}

func (tree *AVL) Remove(id string) error {
	root, err := removeNode(tree.root, id)
	if err != nil {
		return err
	}
	tree.root = root
	return nil
}

func (tree *AVL) Update(id string, updated interface{}) (*BinaryNode, error) {

	// find the node
	found := findNode(tree.root, id)
	if found == nil {
		// not found
		return found, ErrNodeNotFound
	}

	// ok. found node.
	// update the node.
	found.item = updated

	return found, nil
}
