package replacewords

import (
	"sort"
	"strings"
)

func ReplaceWords(dictionary []string, sentence string) string {
	sort.Slice(dictionary, func(i, j int) bool {
		return len(dictionary[i]) < len(dictionary[j])
	})

	words := strings.Split(sentence, " ")

	for i, word := range words {
		for _, root := range dictionary {
			if strings.HasPrefix(word, root) {
				words[i] = root
				break
			}
		}
	}

	return strings.Join(words, " ")
}
