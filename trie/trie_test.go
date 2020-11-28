package trie

import (
	"testing"
)

func TestNode_Sort(t *testing.T) {
	var tests = []*node{
		&node{part: "user", priority: 1},
		&node{part: "system", priority: 2},
		&node{part: ":id", isWild: true, paramKey: "id", priority: -1},
		&node{part: "*file", isWild: true, paramKey: "file", priority: -2},
		&node{part: "order", priority: 3},
		&node{part: "project", priority: 4},
	}
	node := createNode("/")
	for _, item := range tests {
		children = append(children, item)
	}
	sort()
	for i, v := range []int{4, 3, 2, 1, -1, -2} {
		if priority != v {
			t.Errorf("node priority sort error: %d", v)
		}
	}
}
func TestNode_Match(t *testing.T) {
	var tests = []*node{
		&node{part: "user"},
		&node{part: "system"},
		&node{part: "order"},
		&node{part: "project"},
		&node{part: "account"},
		&node{part: "address"},
	}
	var dynamicTests = []*node{
		&node{part: "user", priority: 1},
		&node{part: "system", priority: 2},
		&node{part: ":id", isWild: true, paramKey: "id", priority: -1},
		&node{part: "*file", isWild: true, paramKey: "file", priority: -2},
		&node{part: "order", priority: 3},
		&node{part: "project", priority: 4},
	}
	node := createNode("/")
	for _, item := range tests {
		children = append(children, item)
	}

	for _, v := range tests {
		if actual := match(part); actual == nil || part != part {
			t.Errorf("node match error: %s", part)
		}
	}
	children = children[:0]
	for _, item := range dynamicTests {
		children = append(children, item)
	}
	sort()
	// 动态路由拥有更低的优先级
	for _, part := range []string{"user", "system", "project", "order"} {
		if actual := match(part); actual == nil || part != part {
			t.Errorf("node match error: %s", part)
		}
	}
	// : * 同时存在 则永远不会进入到 *
	if actual := match("12345"); actual == nil || !isWild || paramKey != "id" {
		t.Errorf("node match error dynamic params [:id]")
	}
}

func TestTrie_Match(t *testing.T) {
	var tests = []struct {
		pattern string
		url     string
		payload int
		params  map[string]string
	}{
		{"/v1/user", "/v1/user", 1, nil},
		{"/v1/user/:id", "/v1/user/123", 2, map[string]string{"id": "123"}},
		{"/v2/user/info", "/v2/user/info", 3, nil},
		{"/v1/order/info", "/v1/order/info", 4, nil},
		{"/v1/order/list", "v1/order/list", 5, nil},
		{"/v1/static/*js", "v1/static/js/index.mjs", 6, map[string]string{"js": "js/index.mjs"}},
		{"/v1/article/list", "v1/article/list", 7, nil},
		{"/v1/article/:id", "v1/article/1", 8, map[string]string{"id": "1"}},
		{"/v1/article/:id/comment", "v1/article/1/comment", 9, map[string]string{"id": "1"}},
		{"/v1/article/:id/comment/:cid", "v1/article/1/comment/2", 10, map[string]string{"id": "1", "cid": "2"}},
	}
	trie := NewTrie()
	for _, val := range tests {
		Insert(val.pattern, val.payload)
	}

	for _, val := range tests {
		payload, params := Match(val.url)
		if payload == nil || payload != val.payload {
			t.Errorf("match error: %s", val.url)
		}
		if val.params != nil && params == nil {
			t.Errorf("not match params: %s", val.url)
		}
		if val.params != nil {
			for k, v := range val.params {
				if params[k] != v {
					t.Errorf("match params error: %s", val.url)
				}
			}
		}
	}
}
