package app

import "testing"

func TestPack(t *testing.T) {
	err := Pack(".", "app.zip")
	if err != nil {
		t.Error(err)
	}
}
