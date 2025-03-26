package iot

import (
	"regexp"
	"testing"
)

func TestGoo(t *testing.T) {
	paramsRegex := regexp.MustCompile(`\{\w+\}`)
	resl := paramsRegex.ReplaceAllStringFunc("abc ${ab} 中 ${cc}", func(s string) string {
		return "[" + s + "]"
	})

	t.Log(resl)

}
