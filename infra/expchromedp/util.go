package expchromedp

import "github.com/chromedp/cdproto/cdp"

func NodeValues(nodes []*cdp.Node) []string {
	var vs []string
	for _, n := range nodes {
		vs = append(vs, n.NodeValue)
	}
	return vs
}
