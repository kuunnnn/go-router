package go_router

type tNode struct {
	part     string
	pattern  string
	children []*tNode
	isWild   bool
}

type Trie struct {
	root *tNode
}

func (n *tNode) match() *tNode {
	return nil
}

func (n *tNode) matchChildren() []*tNode {
	return nil
}

func (t *Trie) AddRoute() {

}

func (t *Trie) MatchRoute() Handler {
	return nil
}
