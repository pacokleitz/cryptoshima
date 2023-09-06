package wordlists

func IsWordlistWord(word string, wordList []string) bool {
	for _, w := range wordList {
		if word == w {
			return true
		}
	}
	return false
}
