package tests

import "testing"

func CompareStrings(t *testing.T, s1 string, s2 string) {

	if s1 != s2 {
		t.Errorf(s1 + " is not equal to " + s2)
	}
}

func CompareErrors(t *testing.T, e1 error, e2 error) {

	if e1 != nil && e2 != nil && e1.Error() != e2.Error() {
		t.Errorf(e1.Error() + " is not equal to " + e2.Error())
	} else if e1 == nil && e2 != nil {
		t.Errorf("nil is not equal to " + e2.Error())
	} else if e1 != nil && e2 == nil {
		t.Errorf(e1.Error() + " is not equal to nil")
	}
}
