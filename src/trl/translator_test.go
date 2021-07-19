package trl

import "testing"

func TestTranslate(t *testing.T) {
	trl := translator{}
	if trl.translate("") != "" {
		t.Fatal("failed")
	}
}
