package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID     int
	Parent int
}
type Node struct {
	ID       int
	Children []*Node
}

const ROOT int = 0

var errRootID = errors.New("Root ID Should be 0 ")
var errContinuous = errors.New("record Id is not continuous ")
var errRecordID = errors.New("Invalid record Id ")
var errCycle = errors.New("cycle directly")

func Build(input []Record) (*Node, error) {
	if len(input) == 0 {
		return nil, nil
	}
	// sort Record to make them In numerical order
	sort.Slice(input, func(i, j int) bool {
		return input[i].ID < input[j].ID
	})
	nodeChildMap := make(map[int]*Node)
	for i, record := range input {
		if i == ROOT && record.ID != ROOT {
			return nil, errRootID
		}
		if record.ID < record.Parent {
			return nil, errRecordID
		}
		if i != ROOT && record.ID != i {
			return nil, errContinuous
		}
		if i != ROOT && record.ID == record.Parent {
			return nil, errCycle
		}
		node := Node{ID: record.ID}
		parent := nodeChildMap[record.Parent]
		if parent != nil {
			parent.Children = append(parent.Children, &node)
		}
		nodeChildMap[i] = &node
	}
	return nodeChildMap[input[ROOT].ID], nil
}
