package _42_Valid_Anagram

//import (
//	"sort"
//	"strings"
//)


//func SortString(w string) string {
//	s := strings.Split(w, "")
//	sort.Strings(s)
//	return strings.Join(s, "")
//}
//
//func isAnagram(s string, t string) bool {
//	len_t := len(t)
//	len_s := len(s)
//	if (len_s != len_t){
//		return false
//	}
//
//	t = SortString(t)
//	s = SortString(s)
//
//	if (s != t){
//		return  false
//	}
//	return true
//}

func isAnagram(s string, t string) bool {
	var counterS [256]int
	var counterT [256]int
	for _, c := range s{
		counterS[c]++;
	}
	for _, c := range t{
		counterT[c]++;
	}
	if counterS == counterT{
		return true
	}
	return false
}