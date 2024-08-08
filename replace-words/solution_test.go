package replacewords

import (
	"testing"
)

func TestReplaceWords(t *testing.T) {
	dict := []string{"cat", "bat", "rat"}
	sentence := "the cattle was rattled by the battery"
	got := ReplaceWords(dict, sentence)
	want := "the cat was rat by the bat"
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}

	dict = []string{"a", "aa", "aaa", "aaaa"}
	sentence = "a aa a aaaa aaa aaa aaa aaaaaa bbb baba ababa"
	got = ReplaceWords(dict, sentence)
	want = "a a a a a a a a bbb baba a"
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}
}
