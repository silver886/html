package html

import (
	"io"
	"reflect"
	"strings"
)

// Node is the basic element in HTML document
type Node struct {
	name    string
	attr    *Attr
	content *Contents
}

// Contents is the contents of a node
type Contents []interface{}

// Attr is the attributes of a node
type Attr map[string]Values

// Values is the values of a attribute
type Values []string

// NewNode create a node
func NewNode(name string) *Node {
	return &Node{
		name:    name,
		attr:    &Attr{},
		content: &Contents{},
	}
}

// Copy a node
func (n *Node) Copy() *Node {
	attr := &Attr{}
	for k, v := range *n.attr {
		(*attr)[k] = v
	}
	content := &Contents{}
	for _, v := range *n.content {
		*content = append(*content, v)
	}
	return &Node{
		name:    n.name,
		attr:    attr,
		content: content,
	}
}

// AddAttr attribute to node
func (n *Node) AddAttr(attr Attr) *Node {
	for k, v := range attr {
		(*n.attr)[k] = v
	}
	return n
}

// AddChild node to node
func (n *Node) AddChild(contents ...interface{}) *Node {
	for _, v := range contents {
		if v == nil || (reflect.TypeOf(v).Kind() == reflect.Ptr && reflect.ValueOf(v).IsNil()) {
			continue
		}
		switch v := v.(type) {
		case *Contents:
			s := make(Contents, len(*v))
			for i, v := range *v {
				s[i] = v
			}
			*n.content = append(*n.content, s...)
		default:
			*n.content = append(*n.content, v)
		}
	}
	return n
}

// AddSibling node to node
func (n *Node) AddSibling(contents ...interface{}) *Contents {
	s := make(Contents, len(contents)+1)
	s[0] = n
	for i, v := range contents {
		s[i+1] = v
	}
	return &s
}

// Mershal node to Writer
func (n *Node) Mershal(out io.Writer) {
	n.marshalNode(out)
}

func (n *Node) String() string {
	var buf strings.Builder

	n.marshalNode(&buf)

	return buf.String()
}
