package go_router

type rNode struct {
}

func (t *rNode) match() *rNode {
	return nil
}

func (t *rNode) matchChildren() []*rNode {
	return nil
}

type radixTree struct {
	root *rNode
}

func (t *radixTree) AddRoute() {

}

func (t *radixTree) MatchRoute() Handler {
	return nil
}
