package main

import "fmt"

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

//插入
func (t *Trie) Insert(word string) {
	node := t
	for _, ch := range word {
		ch = ch - 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (t *Trie) SearchPrefix(prefix string) *Trie {
	node := t
	for _, ch := range prefix {
		ch = ch - 'a'
		if node.children[ch] == nil {
			return nil
		}
		node = node.children[ch]
	}
	return node
}

//搜索
func (t *Trie) Search(word string) bool {
	node := t.SearchPrefix(word)
	return node != nil && node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	return t.SearchPrefix(prefix) != nil
}

func main() {
	obj := Constructor()
	obj.Insert("apple")
	fmt.Println(obj.Search("apple"))     //true
	fmt.Println(obj.StartsWith("apple")) //true
}
