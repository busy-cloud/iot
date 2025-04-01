package app

import "testing"

func TestUnpack(t *testing.T) {
	err := Verify("app.zip")
	if err != nil {
		t.Error(err)
	}
	//_ = Check("app.zip", "license")
}
