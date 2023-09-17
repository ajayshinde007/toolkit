package tools

import "testing"

func TestTools_RandomString(t *testing.T) {
	var testTools Tools
	actual := testTools.RandomString(5)
	expected := testTools.RandomString(5)
	if actual == expected {
		t.Error("Not generating random")
	}
	if len(actual) != 5 {
		t.Error("Not generating length correctly")
	}
}
