package trie

import "testing"

func TestTrie_Insert(t *testing.T) {
	trie := NewTrie()

	// Insert some words
	trie.Insert("apple")
	trie.Insert("app")
	trie.Insert("banana")
	trie.Insert("bat")

	// Check if words are found
	if !trie.Search("apple") {
		t.Errorf("Expected 'apple' to be found")
	}
	if !trie.Search("app") {
		t.Errorf("Expected 'app' to be found")
	}
	if !trie.Search("banana") {
		t.Errorf("Expected 'banana' to be found")
	}
	if !trie.Search("bat") {
		t.Errorf("Expected 'bat' to be found")
	}

	// Check if non-existent words are not found
	if trie.Search("orange") {
		t.Errorf("Expected 'orange' to not be found")
	}
	if trie.Search("appl") {
		t.Errorf("Expected 'appl' to not be found")
	}
	if trie.Search("b") {
		t.Errorf("Expected 'b' to not be found")
	}
}

func TestTrie_Insert_WithEmptyRoot(t *testing.T) {
	trie := NewTrie()

	// Insert a word
	trie.Insert("apple")

	// Check if the word is found
	if !trie.Search("apple") {
		t.Errorf("Expected 'apple' to be found")
	}

	// Check if non-existent words are not found
	if trie.Search("orange") {
		t.Errorf("Expected 'orange' to not be found")
	}
	if trie.Search("appl") {
		t.Errorf("Expected 'appl' to not be found")
	}
	if trie.Search("b") {
		t.Errorf("Expected 'b' to not be found")
	}
}

func TestTrie_Insert_WithExistingRoot(t *testing.T) {
	trie := NewTrie()

	// Insert a word
	trie.Insert("apple")

	// Insert another word
	trie.Insert("app")

	// Check if the words are found
	if !trie.Search("apple") {
		t.Errorf("Expected 'apple' to be found")
	}
	if !trie.Search("app") {
		t.Errorf("Expected 'app' to be found")
	}

	// Check if non-existent words are not found
	if trie.Search("orange") {
		t.Errorf("Expected 'orange' to not be found")
	}
	if trie.Search("appl") {
		t.Errorf("Expected 'appl' to not be found")
	}
	if trie.Search("b") {
		t.Errorf("Expected 'b' to not be found")
	}
}

func TestTrie_Insert_WithLongerWord(t *testing.T) {
	trie := NewTrie()

	// Insert a word
	trie.Insert("apple")

	// Insert a longer word
	trie.Insert("applet")

	// Check if the words are found
	if !trie.Search("apple") {
		t.Errorf("Expected 'apple' to be found")
	}
	if !trie.Search("applet") {
		t.Errorf("Expected 'applet' to be found")
	}

	// Check if non-existent words are not found
	if trie.Search("orange") {
		t.Errorf("Expected 'orange' to not be found")
	}
	if trie.Search("appl") {
		t.Errorf("Expected 'appl' to not be found")
	}
	if trie.Search("b") {
		t.Errorf("Expected 'b' to not be found")
	}
}

func TestTrie_Insert_WithMultipleWords(t *testing.T) {
	trie := NewTrie()

	// Insert multiple words
	trie.Insert("apple")
	trie.Insert("app")
	trie.Insert("banana")
	trie.Insert("bat")
	trie.Insert("batmobile")
	trie.Insert("bike")

	// Check if the words are found
	if !trie.Search("apple") {
		t.Errorf("Expected 'apple' to be found")
	}
	if !trie.Search("app") {
		t.Errorf("Expected 'app' to be found")
	}
	if !trie.Search("banana") {
		t.Errorf("Expected 'banana' to be found")
	}
	if !trie.Search("bat") {
		t.Errorf("Expected 'bat' to be found")
	}
	if !trie.Search("batmobile") {
		t.Errorf("Expected 'batmobile' to be found")
	}
	if !trie.Search("bike") {
		t.Errorf("Expected 'bike' to be found")
	}

	// Check if non-existent words are not found
	if trie.Search("orange") {
		t.Errorf("Expected 'orange' to not be found")
	}
	if trie.Search("appl") {
		t.Errorf("Expected 'appl' to not be found")
	}
	if trie.Search("b") {
		t.Errorf("Expected 'b' to not be found")
	}
	if trie.Search("batp") {
		t.Errorf("Expected 'batp' to not be found")
	}
	if trie.Search("bi") {
		t.Errorf("Expected 'bi' to not be found")
	}
}
