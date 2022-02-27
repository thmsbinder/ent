package utils

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Node struct {
	Links map[uint][]Target
}

type Target struct {
	Type uint32
	Hash Hash
}

type Link struct {
	// If base is empty, then it is relative to the same root.
	Base Hash
	Path Path
}

type Path []Selector

type Selector struct {
	FieldID uint
	Index   uint
}

func ParseNode(b []byte) (*Node, error) {
	node := Node{}
	err := json.Unmarshal(b, &node)
	if err != nil {
		return nil, fmt.Errorf("invalid node: %w", err)
	}
	return &node, nil
}

func SerializeNode(node *Node) ([]byte, error) {
	return json.Marshal(node)
}

// Parse a selector of the form "0[1]"
func ParseSelector(s string) (*Selector, error) {
	var fieldID, index uint
	_, err := fmt.Sscanf(s, "%d[%d]", &fieldID, &index)
	if err != nil {
		return nil, fmt.Errorf("invalid selector: %w", err)
	}
	return &Selector{
		FieldID: fieldID,
		Index:   index,
	}, nil
}

func PrintSelector(s Selector) string {
	return fmt.Sprintf("%d[%d]", s.FieldID, s.Index)
}

func ParsePath(s string) (Path, error) {
	selectors := []Selector{}
	for _, s := range strings.Split(s, "/") {
		if s == "" {
			continue
		}
		selector, err := ParseSelector(s)
		if err != nil {
			return nil, fmt.Errorf("invalid selector: %w", err)
		}
		selectors = append(selectors, *selector)
	}
	return selectors, nil
}

func PrintPath(path Path) string {
	out := ""
	for _, s := range path {
		out += "/" + PrintSelector(s)
	}
	return out
}
