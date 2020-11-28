package trie

import (
	"strings"
)

type Payload interface {
}

type node struct {
	// 当前 route 片段
	part string
	// 父节点
	parent *node
	// 完整的路由模式
	pattern string
	// 节点负载的数据
	payload Payload
	// 优先级  对于子节点数更多的节点放在更前面
	// 如果是动态路由, 则降低优先级
	priority int
	// 是否使用占位符
	isWild bool
	// 子节点
	children []*node
	// 动态路由的参数
	paramKey string
	isEnd    bool
}

// 匹配第一个节点, 优先匹配非动态节点
func (n *node) match(part string) *node {
	for _, val := range n.children {
		if val.part == part || val.isWild {
			return val
		}
	}
	return nil
}

// 按优先级进行排序
func (n *node) sort() {
	t, j, l := &node{}, 0, len(n.children)
	for i := 0; i < l; i++ {
		t = n.children[i]
		j = i - 1
		for ; j >= 0 && n.children[j].priority <= t.priority; j-- {
			n.children[j+1] = n.children[j]
		}
		n.children[j+1] = t
	}
}

// 创建一个节点
func createNode(part string) *node {
	return &node{
		part:     part,
		priority: 0,
		isWild:   false,
		isEnd:    false,
		children: []*node{},
	}
}

type Trie struct {
	root *node
}

func NewTrie() *Trie {
	return &Trie{
		root: &node{
			children: []*node{},
		},
	}
}

func (t *Trie) Insert(url string, payload Payload) {
	parts := splitUrl(url)
	node := t.root
	for _, part := range parts {
		matchNode := node.match(part)
		if matchNode == nil {
			matchNode = createNode(part)
			isWild := part[0] == '*' || part[0] == ':'
			matchNode.isWild = isWild
			node.children = append(node.children, matchNode)
			if matchNode.isWild {
				matchNode.paramKey = string(part[1:])
				if part[0] == '*' {
					matchNode.priority = -2
				} else {
					matchNode.priority = -1
				}
			} else {
				matchNode.priority = 0
			}
			if node.parent != nil {
				node.priority++
				node.parent.sort()
			}
		}
		node = matchNode
	}
	node.pattern = join(parts)
	node.payload = payload
	node.isEnd = true
}

func (t *Trie) Match(url string) (interface{}, map[string]string) {
	node := t.root
	parts := splitUrl(url)
	params := make(map[string]string)
	for idx, part := range parts {
		matchNode := node.match(part)
		if matchNode == nil {
			return nil, nil
		}
		if matchNode.isWild {
			if matchNode.part[0] == '*' {
				params[matchNode.paramKey] = strings.Join(parts[idx:], "/")
				return matchNode.payload, params
			}
			params[matchNode.paramKey] = part
		}
		node = matchNode
	}
	if !node.isEnd || node.payload == nil {
		return nil, nil
	}
	return node.payload, params
}
