package types

import (
	"encoding/json"
	"github.com/1uvu/bitlog/pkg/errorx"
)

type (
	LogTree struct {
		Root     *RawLog                       `json:"root"`     // *ChangeLog raw
		Children map[RawLogType][]*LogTreeNode `json:"children"` // map[StatusType][]*StatusLog raw
	}
	LogTreeNode struct {
		Node         *RawLog   `json:"node"`          // StatusLog raw
		NodeChildren []*RawLog `json:"node_children"` // []*EventLog raw
	}
)

func NewLogTree() *LogTree {
	tree := new(LogTree)
	tree.Children = make(map[RawLogType][]*LogTreeNode, 0)
	statusLogTypes := []RawLogType{StatueTypeTx, StatueTypeBlock, StatueTypeChain, StatueTypeNetwork, StatusTypeUnknown}
	for _, logType := range statusLogTypes {
		tree.Children[logType] = make([]*LogTreeNode, 0)
	}
	return tree
}

func (tree *LogTree) AddChild(logType RawLogType, status *RawLog) error {
	if children, ok := tree.Children[logType]; !ok || children == nil {
		return errorx.ErrLogTreeNotFoundChild
	}
	treeNode := NewLogTreeNode(status)
	tree.Children[logType] = append(tree.Children[logType], treeNode)
	return nil
}

func (tree *LogTree) GetChild(logType RawLogType) (*LogTreeNode, error) {
	n := len(tree.Children[logType])
	if n == 0 {
		return nil, errorx.ErrLogTreeNotFoundChild
	}
	return tree.Children[logType][n-1], nil
}

func (tree *LogTree) Serialize(change *RawLog) ([]byte, error) {
	if change == nil {
		return nil, errorx.ErrLogTreeRootIncomplete
	}
	tree.Root = change
	data, err := json.Marshal(*tree)
	return data, err
}

func NewLogTreeNode(status *RawLog) *LogTreeNode {
	// 注意：make node children
	treeNode := new(LogTreeNode)
	treeNode.Node = status
	treeNode.NodeChildren = make([]*RawLog, 0)
	return treeNode
}

func (node *LogTreeNode) AddNodeChild(event *RawLog) {
	node.NodeChildren = append(node.NodeChildren, event)
}
