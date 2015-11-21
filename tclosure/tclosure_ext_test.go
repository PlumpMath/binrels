package tclosure_test

import (
	"testing"

	"github.com/alcortesm/binrels"
)

var g = binrels.Graph{
	"1": map[string]struct{}{
		"3": struct{}{}},
	"2": map[string]struct{}{
		"1": struct{}{},
		"4": struct{}{}},
	"3": map[string]struct{}{},
	"4": map[string]struct{}{
		"2": struct{}{}},
}

func Test(t *testing.T) {
	t.Fatal("TODO: Write these tests")
}
