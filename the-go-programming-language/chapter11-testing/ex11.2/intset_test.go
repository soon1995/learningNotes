package intset

import (
	"testing"
)

func TestAdd(t *testing.T) {
	intset := New()
	intset.Add(1)
	intset.Add(1)
	intset.Add(1)
	if !intset.Has(1) {
		t.Errorf("Add(%d) 3 times = %v, want %v", 1, intset, "{1}")
	}
}

func TestUnionWith(t *testing.T) {
	intset := New()
	intset.Add(1)
	intset.Add(1)
	intset.Add(1)
	intset2 := New()
	intset2.Add(2)
	intset2.Add(3)
	intset2.Add(4)
  intset.UnionWith(intset2)
  if !intset.Has(4) {
		t.Errorf("{1}.UnionWith({2,3,4}) = %v, want %v", intset, "{1 2 3 4}")
  }
}
