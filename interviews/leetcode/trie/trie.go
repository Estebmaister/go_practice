// A trie is a tree-like data structure that stores strings, such as words in a dictionary.
// Tries are also known as prefix trees or digital trees.

// How it works
// Each node in a trie represents a character or prefix of a string.
// Edges connect each node to its children.
// Nodes with no children are called leaf nodes and mark the end of a string.
// To find a word in a trie, start at the root node and follow the branches that correspond to each character in the word.

// Why it's useful
// Tries are efficient for storing words with common prefixes.
// They can be used for string-searching algorithms like spell checking, predictive text, and approximate string matching.
// Tries are used in search engines, text editors, and network routers.
package trie

// Trie node
type Trie struct {
	children [26]*Trie // Ex [AlphabetSize]*RelatedWords
	isEnd    bool      // Rpr ends of a word
}

// NewTrie initializes an empty Trie.
func NewTrie() Trie {
	return Trie{}
}

// Insert adds a word into the Trie.
func (t *Trie) Insert(word string) {
	for i := range word {
		c := word[i] - 'a'
		if child := t.children[c]; child == nil {
			t.children[c] = &Trie{}
		}
		t = t.children[c]
	}
	t.isEnd = true
}

// Search returns if the word is in the Trie.
func (t *Trie) Search(word string) bool {
	for i := range word {
		c := word[i] - 'a'
		if child := t.children[c]; child == nil {
			return false
		}
		t = t.children[c]
	}
	return t.isEnd
}

// StartsWith returns if there is any word in the Trie that starts with the given prefix.
func (t *Trie) StartsWith(prefix string) bool {
	for i := range prefix {
		c := prefix[i] - 'a'
		if child := t.children[c]; child == nil {
			return false
		}
		t = t.children[c]
	}
	return true
}
