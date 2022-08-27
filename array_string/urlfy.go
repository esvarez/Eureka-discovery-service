package str

/**
* Replace blank spaces for "%20"
**/
func urlfy(w string) string {
	// const replace := "%20"
	str := []byte(w)
	bte := []byte(" ")
	blk := []byte("%20")
	i := len(str) - 1
	for ; i >= 0; i-- {
		if str[i] != bte[0] {
			break
		}
	}
	j := len(str) - 1
	for j >= 0 && i >= 0 {
		if str[i] == bte[0] {
			str[j] = blk[2]
			j--
			str[j] = blk[1]
			j--
			str[j] = blk[0]
			j--
			i--
		} else {
			str[j] = str[i]
			j--
			i--
		}
	}

	return string(str)
}
