package _33_Number_of_One

import(
	"testing"
)


func TestShouldBe3(t *testing.T){
	if 6 != countDigitOne(13){
		t.Fail()
	}
}
