package _42_Valid_Anagram

func isAnagram(s string, t string) bool {
	len_t := len(t)
	len_s := len(s)

	if (len_s != len_t){
		return false
	}

	for i, _ := range s{
		if( s[i] != t[len_t - 1 - i]){
			return false
		}
	}

	return true
}

