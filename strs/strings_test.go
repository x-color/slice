package strs_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	. "github.com/x-color/slice/strs"
)

func TestContains(t *testing.T) {
	testcase := []struct {
		l   []string
		s   string
		out bool
	}{
		{[]string{"first"}, "first", true},
		{[]string{"first", "first", "second"}, "first", true},
		{[]string{"first", "second", "third"}, "first", true},
		{[]string{"first", "second", "third"}, "second", true},
		{[]string{"first", "second", "third"}, "third", true},
		{[]string{"first", "second", "third"}, "", false},
		{[]string{"first", "second", "third"}, "fourth", false},
		{[]string{"first", "second", "third"}, "i", false},
		{[]string{}, "first", false},
		{nil, "first", false},
	}
	for _, tc := range testcase {
		act := Contains(tc.l, tc.s)
		if act != tc.out {
			t.Errorf("In(%v, %v) want %v but %v", tc.l, tc.s, tc.out, act)
		}
	}
}

func TestIndex(t *testing.T) {
	testcase := []struct {
		l   []string
		s   string
		out int
	}{
		{[]string{"first"}, "first", 0},
		{[]string{"first", "first", "second"}, "first", 0},
		{[]string{"first", "second", "third"}, "first", 0},
		{[]string{"first", "second", "third"}, "second", 1},
		{[]string{"first", "second", "third"}, "third", 2},
		{[]string{"first", "second", "third"}, "", -1},
		{[]string{"first", "second", "third"}, "fourth", -1},
		{[]string{"first", "second", "third"}, "i", -1},
		{[]string{}, "first", -1},
		{nil, "first", -1},
	}
	for _, tc := range testcase {
		act := Index(tc.l, tc.s)
		if act != tc.out {
			t.Errorf("Index(%v, %v) want %v but %v", tc.l, tc.s, tc.out, act)
		}
	}
}

func TestRemove(t *testing.T) {
	testcase := []struct {
		l   []string
		s   string
		out []string
	}{
		{[]string{"first"}, "first", []string{}},
		{[]string{"first", "first", "second"}, "first", []string{"first", "second"}},
		{[]string{"first", "second", "third"}, "first", []string{"second", "third"}},
		{[]string{"first", "second", "third"}, "second", []string{"first", "third"}},
		{[]string{"first", "second", "third"}, "third", []string{"first", "second"}},
		{[]string{"first", "second", "third"}, "", []string{"first", "second", "third"}},
		{[]string{"first", "second", "third"}, "fourth", []string{"first", "second", "third"}},
		{[]string{"first", "second", "third"}, "i", []string{"first", "second", "third"}},
		{[]string{}, "first", []string{}},
		{nil, "first", []string{}},
	}
	for _, tc := range testcase {
		act := Remove(tc.l, tc.s)
		if d := cmp.Diff(act, tc.out); d != "" {
			t.Errorf("Remove(%v, %v) want %v but %v", tc.l, tc.s, tc.out, act)
		}
	}
}

func TestRemoveE(t *testing.T) {
	testcase := []struct {
		l   []string
		s   string
		out []string
		err bool
	}{
		{[]string{"first"}, "first", []string{}, false},
		{[]string{"first", "first", "second"}, "first", []string{"first", "second"}, false},
		{[]string{"first", "second", "third"}, "first", []string{"second", "third"}, false},
		{[]string{"first", "second", "third"}, "second", []string{"first", "third"}, false},
		{[]string{"first", "second", "third"}, "third", []string{"first", "second"}, false},
		{[]string{"first", "second", "third"}, "", nil, true},
		{[]string{"first", "second", "third"}, "fourth", nil, true},
		{[]string{"first", "second", "third"}, "i", nil, true},
		{[]string{}, "first", nil, true},
		{nil, "first", nil, true},
	}
	for _, tc := range testcase {
		act, err := RemoveE(tc.l, tc.s)
		if d := cmp.Diff(act, tc.out); d != "" {
			t.Errorf("RemoveE(%v, %v) want %v but %v", tc.l, tc.s, tc.out, act)
		}
		if (err != nil) != tc.err {
			t.Errorf("RemoveE(%v, %v) want error but not", tc.l, tc.s)
		}
	}
}

func TestRemoveN(t *testing.T) {
	testcase := []struct {
		l   []string
		s   string
		n   int
		out []string
	}{
		{[]string{"first"}, "first", 1, []string{}},
		{[]string{"first", "first", "second"}, "first", 0, []string{"first", "first", "second"}},
		{[]string{"first", "first", "second"}, "first", 1, []string{"first", "second"}},
		{[]string{"first", "first", "second"}, "first", 2, []string{"second"}},
		{[]string{"first", "first", "second"}, "first", 3, []string{"second"}},
		{[]string{"first", "first", "second"}, "first", -1, []string{"second"}},
		{[]string{"first", "first", "second"}, "first", -2, []string{"second"}},
		{[]string{"first", "second", "third"}, "first", -1, []string{"second", "third"}},
		{[]string{"first", "second", "third"}, "second", -1, []string{"first", "third"}},
		{[]string{"first", "second", "third"}, "third", -1, []string{"first", "second"}},
		{[]string{"first", "second", "third"}, "", -1, []string{"first", "second", "third"}},
		{[]string{"first", "second", "third"}, "fourth", -1, []string{"first", "second", "third"}},
		{[]string{"first", "second", "third"}, "i", -1, []string{"first", "second", "third"}},
		{[]string{}, "first", -1, []string{}},
		{nil, "first", -1, []string{}},
	}
	for _, tc := range testcase {
		act := RemoveN(tc.l, tc.s, tc.n)
		if d := cmp.Diff(act, tc.out); d != "" {
			t.Errorf("RemoveN(%v, %v, %v) want %v but %v", tc.l, tc.s, tc.n, tc.out, act)
		}
	}
}

func TestRemoveNE(t *testing.T) {
	testcase := []struct {
		l   []string
		s   string
		n   int
		out []string
		err bool
	}{
		{[]string{"first"}, "first", 1, []string{}, false},
		{[]string{"first", "first", "second"}, "first", 0, nil, true},
		{[]string{"first", "first", "second"}, "first", 1, []string{"first", "second"}, false},
		{[]string{"first", "first", "second"}, "first", 2, []string{"second"}, false},
		{[]string{"first", "first", "second"}, "first", 3, []string{"second"}, false},
		{[]string{"first", "first", "second"}, "first", -1, []string{"second"}, false},
		{[]string{"first", "second", "third"}, "first", -1, []string{"second", "third"}, false},
		{[]string{"first", "second", "third"}, "second", -1, []string{"first", "third"}, false},
		{[]string{"first", "second", "third"}, "third", -1, []string{"first", "second"}, false},
		{[]string{"first", "second", "third"}, "", -1, nil, true},
		{[]string{"first", "second", "third"}, "fourth", -1, nil, true},
		{[]string{"first", "second", "third"}, "i", -1, nil, true},
		{[]string{}, "first", -1, nil, true},
		{nil, "first", -1, nil, true},
	}
	for _, tc := range testcase {
		act, err := RemoveNE(tc.l, tc.s, tc.n)
		if !cmp.Equal(act, tc.out) {
			t.Errorf("RemoveNE(%v, %v, %v) want %v but %v", tc.l, tc.s, tc.n, tc.out, act)
		}
		if (err != nil) != tc.err {
			t.Errorf("RemoveNE(%v, %v, %v) want error but not", tc.l, tc.s, tc.n)
		}
	}
}

func TestUniq(t *testing.T) {
	testcase := []struct {
		l   []string
		out []string
	}{
		{[]string{"a"}, []string{"a"}},
		{[]string{"a", "b"}, []string{"a", "b"}},
		{[]string{"a", "b", "c"}, []string{"a", "b", "c"}},
		{[]string{"a", "a", "c"}, []string{"a", "c"}},
		{[]string{"a", "a", "a"}, []string{"a"}},
		{[]string{}, []string{}},
		{nil, []string{}},
	}
	for _, tc := range testcase {
		act := Uniq(tc.l)
		if !cmp.Equal(act, tc.out, cmpopts.SortSlices(less)) {
			t.Errorf("Uniq(%v) want %v but %v", tc.l, tc.out, act)
		}
	}
}

func TestSub(t *testing.T) {
	testcase := []struct {
		a   []string
		b   []string
		out []string
	}{
		{[]string{"a"}, []string{"a"}, []string{}},
		{[]string{"a", "b"}, []string{"a"}, []string{"b"}},
		{[]string{"a", "b", "c"}, []string{"a", "c"}, []string{"b"}},
		{[]string{"a", "b", "c"}, []string{"a", "c", "d"}, []string{"b"}},
		{[]string{"a", "b", "c"}, []string{"a", "b", "c", "d"}, []string{}},
		{[]string{"a", "b", "c"}, []string{}, []string{"a", "b", "c"}},
		{[]string{}, []string{"a", "b", "c"}, []string{}},
		{[]string{}, []string{}, []string{}},
		{[]string{}, nil, []string{}},
		{nil, []string{}, []string{}},
	}
	for _, tc := range testcase {
		act := Sub(tc.a, tc.b)
		if !cmp.Equal(act, tc.out, cmpopts.SortSlices(less)) {
			t.Errorf("Sub(%v, %v) want %v but %v", tc.a, tc.b, tc.out, act)
		}
	}
}

func less(a, b string) bool {
	return a < b
}
