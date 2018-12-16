package html

import (
	"fmt"
	"io"
)

func (c *Contents) marshalContent(out io.Writer) {
	for _, v := range *c {
		switch v := v.(type) {
		case *Node:
			out.Write([]byte("<"))
			out.Write([]byte(v.name))

			if len(*v.attr) > 0 {
				v.attr.marshalAttr(out)
			}

			if len(*v.content) == 0 {
				out.Write([]byte(" />"))
			} else {
				out.Write([]byte(">"))

				v.content.marshalContent(out)

				out.Write([]byte("</"))
				out.Write([]byte(v.name))
				out.Write([]byte(">"))
			}
		case *Contents:
			v.marshalContent(out)
		default:
			out.Write([]byte(fmt.Sprint(v)))
		}
	}
}

func (n *Node) marshalNode(out io.Writer) {
	out.Write([]byte("<"))
	out.Write([]byte(n.name))

	if len(*n.attr) > 0 {
		n.attr.marshalAttr(out)
	}

	if len(*n.content) == 0 {
		out.Write([]byte(" />"))
	} else {
		out.Write([]byte(">"))

		n.content.marshalContent(out)

		out.Write([]byte("</"))
		out.Write([]byte(n.name))
		out.Write([]byte(">"))
	}
}

func (a *Attr) marshalAttr(out io.Writer) {
	for k, v := range *a {
		out.Write([]byte(" "))
		out.Write([]byte(k))
		if len(v) > 0 {
			out.Write([]byte("=\""))
			for k, v := range v {
				if k != 0 {
					out.Write([]byte(" "))
				}
				out.Write([]byte(v))
			}
			out.Write([]byte("\""))
		}
	}
}
