package app

import "testing"

func TestPack(t *testing.T) {
	err := Pack("../bin", "app.zip")
	if err != nil {
		t.Error(err)
	}
	//_ = Check("app.zip", "license")
}
