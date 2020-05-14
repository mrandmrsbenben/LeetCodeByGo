package main

import "fmt"

/**

//执行用时 :128 ms, 在所有 Go 提交中击败了10.23%的用户
//内存消耗 :25.5 MB, 在所有 Go 提交中击败了33.33%的用户
// Trie Trie
type Trie struct {
	words  map[string]int
	leaves map[string]*Trie
}

// Constructor Initialize your data structure here.
func Constructor() Trie {
	return Trie{make(map[string]int), make(map[string]*Trie)}
}

// Insert Inserts a word into the trie.
func (this *Trie) Insert(word string) {
	if len(word) > 0 {
		this.words[word] = 1
		c := word[0:1]
		if t, ok := this.leaves[c]; ok {
			t.Insert(word[1:])
		} else {
			t = &Trie{make(map[string]int), make(map[string]*Trie)}
			t.Insert(word[1:])
			this.leaves[c] = t
		}
	}
}

// Search Returns if the word is in the trie.
func (this *Trie) Search(word string) bool {
	_, ok := this.words[word]
	return ok
}

//StartsWith Returns if there is any word in the trie that starts with the given prefix.
func (this *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 1 {
		_, ok := this.leaves[prefix]
		return ok
	}

	c := prefix[0:1]
	if t, ok := this.leaves[c]; ok {
		return t.StartsWith(prefix[1:])
	}
	return false
}
*/

// Trie Trie
type Trie struct {
	words map[string]int
	tries [26]*Trie
}

// Constructor Initialize your data structure here.
func Constructor() Trie {
	return Trie{make(map[string]int), [26]*Trie{}}
}

// Insert Inserts a word into the trie.
func (this *Trie) Insert(word string) {
	if len(word) > 0 {
		this.words[word] = 1
		i := word[0] - 'a'
		if this.tries[i] == nil {
			this.tries[i] = &Trie{make(map[string]int), [26]*Trie{}}
		}
		this.tries[i].Insert(word[1:])
	}
}

// Search Returns if the word is in the trie.
func (this *Trie) Search(word string) bool {
	_, ok := this.words[word]
	return ok
}

//StartsWith Returns if there is any word in the trie that starts with the given prefix.
func (this *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 1 {
		return this.tries[prefix[0]-'a'] != nil
	}

	i := prefix[0] - 'a'
	if this.tries[i] != nil {
		return this.tries[i].StartsWith(prefix[1:])
	}
	return false
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
func main() {
	t := Constructor()
	t.Insert("apple")
	fmt.Println("Output: ", t.Search("apple"))
	fmt.Println("Output: ", t.Search("app"))
	fmt.Println("Output: ", t.StartsWith("app"))
	t.Insert("app")
	fmt.Println("Output: ", t.Search("app"))
	fmt.Println("Output: ", t.StartsWith("appl"))
}
