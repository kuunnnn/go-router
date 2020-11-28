package trie

import (
	"testing"
)

func TestToSlice(t *testing.T) {
	var tests = []struct {
		in       string
		expected string
	}{
		{"/a/b/c", "/a/b/c"},
		{"a/b/c/", "/a/b/c"},
		{"//a///b///c", "/a/b/c"},
		{"//a///b///c///", "/a/b/c"},
		{"//a///:b///c///", "/a/:b/c"},
		{"//a///*b///c///", "/a/*b/c"},
	}
	for _, v := range tests {
		if actual := splitUrl(v.in); join(actual) != v.expected {
			t.Errorf("splitUrl(%s) = %+v; expected %+v", v.in, actual, v.expected)
		}
	}
}
func TestFormat(t *testing.T) {
	var tests = []struct {
		in       string
		expected string
	}{
		{"/a/b/c", "/a/b/c"},
		{"a/b/c/", "/a/b/c"},
		{"//a///b///c", "/a/b/c"},
		{"//a///b///c///", "/a/b/c"},
		{"//a///:b///c///", "/a/:b/c"},
		{"//a///*b///c///", "/a/*b/c"},
	}
	for _, v := range tests {
		if actual := format(v.in); actual != v.expected {
			t.Errorf("format(%s) = %+v; expected %+v", v.in, actual, v.expected)
		}
	}
}

func TestJoin(t *testing.T) {
	var tests = []struct {
		in       []string
		expected string
	}{
		{[]string{"a", "b", "c"}, "/a/b/c"},
		{[]string{}, "/"},
	}
	for _, v := range tests {
		if actual := join(v.in); actual != v.expected {
			t.Errorf("join(%s) = %+v; expected %+v", v.in, actual, v.expected)
		}
	}
}
