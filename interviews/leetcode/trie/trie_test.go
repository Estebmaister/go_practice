package trie

import (
	"fmt"
	"os"
	"testing"
)

// go test ./interviews/leetcode/trie/...
type fileLogger struct {
	messagesFile *os.File
}

func (l *fileLogger) logf(format string, a ...any) {
	if l.messagesFile == nil {
		l.messagesFile, _ = os.Create("tests.out")
	}
	fmt.Fprintf(l.messagesFile, format, a...)
}

var l = &fileLogger{}

func TestTrie(t *testing.T) {
	trie := NewTrie()
	l.logf("Constructor: %v\n", trie)
	trie.Insert("apple")
	l.logf("First insert (apple): %v\n", trie)
	l.logf("Child a(%v): %v\n", "a"[0]-'a', trie.children["a"[0]-'a'])
	trie.Search("apple")           // return True
	appFound := trie.Search("app") // return False
	l.logf("Searching app, stored apple: %v\n", appFound)
	trie.Search("apples")  // return False
	trie.StartsWith("app") // return True
	trie.StartsWith("ban") // return False
	trie.Insert("app")
	trie.Search("app") // return True
	l.logf("Second insert (app): %v\n", trie)
	trie.Insert("banana")
	l.logf("Third insert (banana): %v\n", trie)
	l.logf(
		"Child ap(p) is end of a word: %v\n",
		*trie.children[0].children["p"[0]-'a'].children["p"[0]-'a'])
}
