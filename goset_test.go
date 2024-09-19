package goset_test

import (
	"testing"

	"github.com/hedykan/goset"
)

func TestNewSet(t *testing.T) {
	goset.NewSet([]int{1, 2, 3})
	goset.NewSet([]string{"a", "b", "c"})
}

func TestNewCountSet(t *testing.T) {
	goset.NewCountSet([]int{1, 2, 3})
	goset.NewCountSet([]string{"a", "b", "c"})
}
