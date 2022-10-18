package types

import (
	"encoding/json"
	"github.com/1uvu/bitlog/pkg/errorx"
)

type (
	LogTree struct {
		Root     *RawLog                       // *ChangeLog raw
		Children map[RawLogType][]*LogTreeNode // map[StatusType][]*StatusLog raw
	}
	LogTreeNode struct {
		Node         *RawLog   // StatusLog raw
		NodeChildren []*RawLog // []*EventLog raw
	}
)

func NewLogTree() *LogTree {
	// 注意：for 循环建立 children

	return nil
}

func NewLogTreeNode(raw *RawLog) *LogTreeNode {
	// 注意：for 循环建立 node children

	return nil
}

func (tree *LogTree) AddChild(logType RawLogType, raw *RawLog) error {
	if children, ok := tree.Children[logType]; !ok || children == nil {
		return errorx.ErrLogTreeNotFoundChild
	}
	treeNode := NewLogTreeNode(raw)
	tree.Children[logType] = append(tree.Children[logType], treeNode)
	return nil
}

// TODO 添加各种方法

func (tree *LogTree) CommitChange(raw *RawLog) {
	tree.Root = raw
}

func (tree *LogTree) Marshal() ([]byte, error) {
	if tree.Root == nil {
		return nil, errorx.ErrLogTreeRootIncomplete
	}
	data, err := json.Marshal(*tree)
	return data, err
}

func (tree *LogTree) StatusTransfer(raw *RawLog) {
	// 状态迁移 并进行后续工作

	tree.Root = raw

	// 后续工作
}

func (tree *LogTree) String() string {
	return ""
}
