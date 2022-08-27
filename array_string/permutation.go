package str

/**
* Validate if two string are permitation
 */
func isPermutation(str, str2 string) bool {
	if len(str) != len(str2) {
		return false
	}

	m1 := make(map[byte]int)
	m2 := make(map[byte]int)

	for i := 0; i < len(str); i++ {
		m1[str[i]]++
		m2[str2[i]]++
	}

	if len(m1) != len(m2) {
		return false
	}

	for k, v := range m1 {
		if v2, ok := m2[k]; ok {
			if v == v2 {
				continue
			}
		}
		return false
	}

	return true
}
