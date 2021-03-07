// +build gofuzz

package goquery

import "bytes"

func FuzzGoQuery (data []byte) int {
	doc, err := NewDocumentFromReader(bytes.NewReader(data))

	if err != nil {
		return 0
	}

	if len(doc.Nodes) == 0 {
		panic("no nodes!")
	}


	sel := doc.Find("*")

	if len(sel.Nodes) < 4 {
		return 0
	}

	res, err := OuterHtml(doc.Selection)

	if err != nil {
		panic(err)
	}

	if res == "" {
		panic(err)
	}

	return 1
}