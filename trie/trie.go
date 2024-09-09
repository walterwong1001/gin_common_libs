package trie

import (
	"regexp"
	"strings"
)

type node[T any] struct {
	meta     T
	children map[string]*node[T]
	end      bool
}

type Trie[T any] struct {
	root *node[T]
}

func (n node[T]) newEmptyNode() *node[T] {
	return &node[T]{
		children: make(map[string]*node[T]),
	}
}

// New initializes and returns a new Trie.
func New[T any]() *Trie[T] {
	return &Trie[T]{root: &node[T]{children: make(map[string]*node[T])}}
}

// Add adds a key (segments) to the Trie with associated metadata.
func (t *Trie[T]) Add(segments []string, meta T) {
	node := t.root
	for _, v := range segments {
		if _, exists := node.children[v]; !exists {
			node.children[v] = node.newEmptyNode()
		}
		node = node.children[v]
	}
	node.end = true
	node.meta = meta
}

// Find searches for a key (segments) in the Trie and returns the metadata.
func (t *Trie[T]) Find(segments []string) *T {
	n := t.root
	for _, segment := range segments {
		// 如果不存在，则判断是否为正则
		if _, exists := n.children[segment]; !exists {
			var b bool
			for k, child := range n.children {
				if strings.HasPrefix(k, "^") && strings.HasSuffix(k, "$") {
					if matched, err := regexp.MatchString(k, segment); err == nil && matched {
						n = child
						b = true
						break
					}
				}
			}
			if !b {
				return nil
			}
		} else {
			n = n.children[segment]
		}
	}
	return &n.meta
}
