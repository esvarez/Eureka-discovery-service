package str

/**
* Implement an algorithm to determine if a string contain unique characters
**/
func isUnique(word string) bool {
	unique := make(map[rune]bool)
	for _, v := range word {
		if _, ok := unique[v]; ok {
			return false
		}
		unique[v] = true
	}
	return true
}
