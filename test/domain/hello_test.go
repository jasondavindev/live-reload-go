package domain_test

import (
	"testing"

	"github.com/jasondavindev/hacktoberfest-2020/domain"
)

func TestHelloFunc(t *testing.T) {
	text := "hello world"

	if domain.HelloFunc(text) != text {
		t.Errorf("Failed. Expected %s to equal %s", text, text)
	}
}
