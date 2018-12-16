package html

import (
	"io"
	"reflect"
	"strings"
)

// HTML presents a HTML document
type HTML struct {
	docType string
	content *Contents
}

// NewHTML create a HTML document
func NewHTML(docType string) *HTML {
	return &HTML{
		docType: docType,
		content: &Contents{},
	}
}

// Copy a HTML document
func (h *HTML) Copy() *HTML {
	content := &Contents{}
	for k, v := range *h.content {
		(*content)[k] = v
	}
	return &HTML{
		docType: h.docType,
		content: content,
	}
}

// AddChild node to HTML document
func (h *HTML) AddChild(contents ...interface{}) *HTML {
	for _, v := range contents {
		if v == nil || (reflect.TypeOf(v).Kind() == reflect.Ptr && reflect.ValueOf(v).IsNil()) {
			continue
		}
		switch v := v.(type) {
		case *[]*Node:
			s := make([]interface{}, len(*v))
			for i, v := range *v {
				s[i] = v
			}
			*h.content = append(*h.content, s...)
		default:
			*h.content = append(*h.content, v)
		}
	}
	return h
}

// Mershal HTML to Writer
func (h *HTML) Mershal(out io.Writer) {
	out.Write([]byte("<!DOCTYPE "))
	out.Write([]byte(h.docType))
	out.Write([]byte("><html"))

	if len(*h.content) == 0 {
		out.Write([]byte(" />"))
	} else {
		out.Write([]byte(">"))
		h.content.marshalContent(out)
		out.Write([]byte("</html>"))
	}
}

func (h *HTML) String() string {
	var buf strings.Builder

	h.Mershal(&buf)

	return buf.String()
}
