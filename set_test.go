package set

import (
	"reflect"
	"sort"
	"testing"
)

func TestAddAndContains(t *testing.T) {
	var s Set[int]
	if s.Contains(1) {
		t.Error("Contains(1) on empty set\n got: true\nwant: false")
	}
	s.Add(1)
	if !s.Contains(1) {
		t.Error("Contains(1)\n got: false\nwant: true")
	}
	s.Add(2)
	if !s.Contains(2) {
		t.Error("Contains(2)\n got: false\nwant: true")
	}
}

func TestAddAll(t *testing.T) {
	var a Set[int]
	b := Of(1, 2, 3)
	a.AddAll(b)
	if !a.ContainsAll(b) {
		t.Error("ContainsAll\n got: false\nwant: true")
	}
	b.Add(4)
	if a.ContainsAll(b) {
		t.Error("ContainsAll\n got: true\nwant: false")
	}
}

func TestAsSlice(t *testing.T) {
	got := Of(1, 2, 3).AsSlice()
	sort.Ints(got)
	want := []int{1, 2, 3}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("AsSlice\n got: %v\nwant: %v", got, want)
	}
}

func TestEquals(t *testing.T) {
	a := New[int]()
	if !a.Equal(a) {
		t.Error("Equal to itself\n got: false\nwant: true")
	}
	b := New[int]()
	if !a.Equal(b) {
		t.Error("Equal with empty sets\n got: false\nwant: true")
	}
	a.Add(1)
	if a.Equal(b) {
		t.Error("Equal\n got: true\nwant: false")
	}
	b.Add(1)
	if !a.Equal(b) {
		t.Error("Equal\n got: false\nwant: true")
	}
	var c *Set[int]
	var d *Set[int]
	if !c.Equal(d) {
		t.Error("Equal with nil pointers\n got: false\nwant: true")
	}
	e := Of(1, 2, 3)
	if c.Equal(e) {
		t.Error("Equal with nil receiver\n got: true\nwant: false")
	}
	if e.Equal(c) {
		t.Error("Equal with nil argument\n got: true\nwant: false")
	}
}

func TestLen(t *testing.T) {
	var s Set[int]
	if got := s.Len(); got != 0 {
		t.Errorf("Len\n got: %d\nwant: 0", got)
	}
	s.Add(1)
	if got := s.Len(); got != 1 {
		t.Errorf("Len\n got: %d\nwant: 1", got)
	}
	s.Add(2)
	s.Add(2)
	if got := s.Len(); got != 2 {
		t.Errorf("Len\n got: %d\nwant: 2", got)
	}
}

func TestContainsAll(t *testing.T) {
	a := Of(1, 2, 3, 4, 5)
	b := Of(1, 2, 3, 4, 5)
	if !a.ContainsAll(b) {
		t.Error("ContainsAll\n got: false\nwant: true")
	}
	b.Add(6)
	if a.ContainsAll(b) {
		t.Error("ContainsAll\n got: true\nwant: false")
	}
}

func TestIsEmpty(t *testing.T) {
	var s Set[int]
	if !s.IsEmpty() {
		t.Error("IsEmpty\n got: false\nwant: true")
	}
	s.Add(1)
	if s.IsEmpty() {
		t.Error("IsEmpty\n got: true\nwant: false")
	}
}

func TestIsSubsetOf(t *testing.T) {
	xs := Of(1, 2, 3, 4, 5)
	ys := Of(1, 2, 3, 4, 5)
	if !xs.IsSubsetOf(ys) {
		t.Error("IsSubsetOf\n got: false\nwant: true")
	}
	if !ys.IsSubsetOf(xs) {
		t.Error("IsSubsetOf\n got: false\nwant: true")
	}
	ys.Add(6)
	if !xs.IsSubsetOf(ys) {
		t.Error("IsSubsetOf\n got: false\nwant: true")
	}
	if ys.IsSubsetOf(xs) {
		t.Error("IsSubsetOf\n got: true\nwant: false")
	}
}

func TestIsProperSubsetOf(t *testing.T) {
	xs := Of(1, 2, 3, 4, 5)
	ys := Of(1, 2, 3, 4, 5)
	if xs.IsProperSubsetOf(ys) {
		t.Error("IsProperSubsetOf\n got: true\nwant: false")
	}
	ys.Add(6)
	if !xs.IsProperSubsetOf(ys) {
		t.Error("IsProperSubsetOf\n got: false\nwant: true")
	}
}

func TestIsSupersetOf(t *testing.T) {
	xs := Of(1, 2, 3, 4, 5)
	ys := Of(1, 2, 3, 4, 5)
	if !xs.IsSupersetOf(ys) {
		t.Error("IsSupersetOf\n got: false\nwant: true")
	}
	ys.Add(6)
	if xs.IsSupersetOf(ys) {
		t.Error("IsSupersetOf\n got: true\nwant: false")
	}
	xs.Add(6)
	xs.Add(7)
	if !xs.IsSupersetOf(ys) {
		t.Error("IsSupersetOf\n got: false\nwant: true")
	}
}

func TestIsProperSupersetOf(t *testing.T) {
	xs := Of(1, 2, 3)
	ys := Of(1, 2, 3)
	if xs.IsProperSupersetOf(ys) {
		t.Error("IsProperSupersetOf\n got: true\nwant: false")
	}
	xs.Add(4)
	if !xs.IsProperSupersetOf(ys) {
		t.Error("IsProperSupersetOf\n got: false\nwant: true")
	}
	if ys.IsProperSupersetOf(xs) {
		t.Error("IsProperSupersetOf\n got: true\nwant: false")
	}
}

func TestRemove(t *testing.T) {
	elements := []int{1, 2, 3}
	xs := Of(elements...)
	for _, e := range elements {
		xs.Remove(e)
		if xs.Contains(e) {
			t.Errorf("Contains(%d)\n got: true\nwant: false", e)
		}
	}
}

func TestRemoveAll(t *testing.T) {
	a := Of(1, 2, 3)
	removed := []int{1, 2}
	a.RemoveAll(Of(removed...))
	for _, i := range removed {
		if a.Contains(i) {
			t.Errorf("Contains(%d)\n got: true\nwant: false", i)
		}
	}
}

func TestDifference(t *testing.T) {
	a := Of(1, 2, 3)
	b := Of(2, 3, 4)
	if got, want := Difference(a, b), Of(1); !got.Equal(want) {
		t.Errorf("Difference\n got: %v\nwant: %v", got.AsSlice(), want.AsSlice())
	}
}

func TestSymmetricDifference(t *testing.T) {
	a := Of(1, 2, 3)
	b := Of(2, 3, 4)
	if got, want := SymmetricDifference(a, b), Of(1, 4); !got.Equal(want) {
		t.Errorf("SymmetricDifference\n got: %v\nwant: %v", got.AsSlice(), want.AsSlice())
	}
}

func TestIntersection(t *testing.T) {
	a := Of(1, 2, 3)
	b := Of(2, 3, 4)
	if got, want := Intersection(a, b), Of(2, 3); !got.Equal(want) {
		t.Errorf("Intersection\n got: %v\nwant: %v", got.AsSlice(), want.AsSlice())
	}
}

func TestUnion(t *testing.T) {
	a := Of(1, 2, 3)
	b := Of(2, 3, 4)
	if got, want := Union(a, b), Of(1, 2, 3, 4); !got.Equal(want) {
		t.Errorf("Union\n got: %v\nwant: %v", got.AsSlice(), want.AsSlice())
	}
}
