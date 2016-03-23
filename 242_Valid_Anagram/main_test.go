package _42_Valid_Anagram

import(
	"testing"
)

func Test_Should_return_true(t *testing.T){
	s2 := "abc"
	s1 := "cba"
	result := isAnagram(s1, s2)
	if (result) {
		t.Log("right")
	}else{
		t.Error("wrong")
	}
}

func Test_Should_return_false(t *testing.T){
	s2 := "abc"
	s1 := "cbaa"
	result := isAnagram(s1, s2)
	if (result) {
		t.Error("wrong")
	}else{
		t.Log("right")
	}
}