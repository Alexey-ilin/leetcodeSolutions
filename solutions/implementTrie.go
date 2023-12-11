package solutions

type Trie struct {
	isWord   bool
	children map[rune]*Trie
}

func ConstrucTrie() Trie {
	return Trie{isWord: false, children: make(map[rune]*Trie)}
}

func (t *Trie) Insert(word string) {
	node := t
	for _, c := range word {
		if _, ok := node.children[c]; !ok {
			node.children[c] = &Trie{children: make(map[rune]*Trie)}
		}
		node = node.children[c]
	}
	node.isWord = true

}

func (t *Trie) Search(word string) bool {
	node := t
	for _, c := range word {
		if _, ok := node.children[c]; !ok {
			return false
		}
		node = node.children[c]
	}
	return node.isWord
}

func (t *Trie) StartsWith(prefix string) bool {
	node := t
	for _, c := range prefix {
		if _, ok := node.children[c]; !ok {
			return false
		}
		node = node.children[c]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
