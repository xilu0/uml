package main

import (
	"fmt"
	"testing"
)

// TrieNode 定义前缀树的节点结构
type TrieNode struct {
	children    map[rune]*TrieNode
	isEnd       bool
	permissions []string // 存储权限信息
}

// Trie 定义前缀树的结构
type Trie struct {
	root *TrieNode
}

// NewTrie 创建一个新的前缀树
func NewTrie() *Trie {
	return &Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

// Insert 向前缀树中插入路径及其对应的权限信息
func (t *Trie) Insert(path string, permissions []string) {
	node := t.root
	for _, char := range path {
		if _, found := node.children[char]; !found {
			node.children[char] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		node = node.children[char]
	}
	node.isEnd = true
	node.permissions = permissions
}

// Search 查找路径是否存在于前缀树中
func (t *Trie) Search(path string) (*TrieNode, bool) {
	node := t.root
	for _, char := range path {
		if _, found := node.children[char]; !found {
			return nil, false
		}
		node = node.children[char]
	}
	return node, node.isEnd
}

// CheckPermissions 检查是否具有所需的权限
func (t *Trie) CheckPermissions(path string, requiredPermissions []string) bool {
	node := t.root
	var maxMatchNode *TrieNode

	for _, char := range path {
		if nextNode, found := node.children[char]; found {
			node = nextNode
			if node.isEnd {
				maxMatchNode = node
			}
		} else {
			break
		}
	}

	if maxMatchNode == nil {
		return false
	}

	return hasRequiredPermissions(maxMatchNode.permissions, requiredPermissions)
}

// hasRequiredPermissions 检查是否具有所有必需的权限
func hasRequiredPermissions(nodePermissions, requiredPermissions []string) bool {
	permSet := make(map[string]bool)
	for _, perm := range nodePermissions {
		permSet[perm] = true
	}
	for _, reqPerm := range requiredPermissions {
		if !permSet[reqPerm] {
			return false
		}
	}
	return true
}

// main 主函数展示如何使用前缀树进行权限校验
func main() {
	trie := NewTrie()
	trie.Insert("/api/user/create", []string{"admin", "write"})
	trie.Insert("/api/user/delete", []string{"admin", "delete"})
	trie.Insert("/api/user/view", []string{"user", "view"})

	testCases := []struct {
		path                string
		requiredPermissions []string
		expectedResult      bool
	}{
		{"/api/user/create", []string{"admin", "write"}, true},
		{"/api/user/delete", []string{"admin", "delete"}, true},
		{"/api/user/view", []string{"user", "view"}, true},
		{"/api/user/view", []string{"admin", "view"}, false},
		{"/api/user/create", []string{"write"}, false},
	}

	for _, tc := range testCases {
		result := trie.CheckPermissions(tc.path, tc.requiredPermissions)
		fmt.Printf("Path: %s, Required: %v, Result: %v (Expected: %v)\n",
			tc.path, tc.requiredPermissions, result, tc.expectedResult)
	}
}

// TestCheckPermissions 测试函数
func TestCheckPermissions(t *testing.T) {
	trie := NewTrie()
	trie.Insert("/api/user/create", []string{"admin", "write"})
	trie.Insert("/api/user/delete", []string{"admin", "delete"})
	trie.Insert("/api/user/view", []string{"user", "view"})

	testCases := []struct {
		path                string
		requiredPermissions []string
		expectedResult      bool
	}{
		{"/api/user/create", []string{"admin", "write"}, true},
		{"/api/user/delete", []string{"admin", "delete"}, true},
		{"/api/user/view", []string{"user", "view"}, true},
		{"/api/user/view", []string{"admin", "view"}, false},
		{"/api/user/create", []string{"write"}, false},
	}

	for _, tc := range testCases {
		result := trie.CheckPermissions(tc.path, tc.requiredPermissions)
		if result != tc.expectedResult {
			t.Errorf("Path: %s, Required: %v, Result: %v (Expected: %v)\n",
				tc.path, tc.requiredPermissions, result, tc.expectedResult)
		}
	}
}
