package strs

import (
	"fmt"
	"sort"
)

// Contains reports whether s is within l.
func Contains(l []string, s string) bool {
	for _, v := range l {
		if v == s {
			return true
		}
	}
	return false
}

// Index returns the index of the first instance of s in l, or -1 if s is not present in l.
func Index(l []string, s string) int {
	for i, v := range l {
		if v == s {
			return i
		}
	}
	return -1
}

// RemoveAt returns a copy of the string slice l with the i-th instances removed.
// If i is not within range from 0 to len(l)-1, it returns a copy of the slice.
func RemoveAt(l []string, i int) []string {
	if i < 0 || i > len(l)-1 {
		return l
	}
	if i == len(l)-1 {
		return l[:i]
	}
	return append(l[:i], l[i+1:]...)
}

// RemoveAtE returns a copy of the string slice l with the i-th instances removed.
// If i is not within range from 0 to len(l)-1, it reports error.
func RemoveAtE(l []string, i int) ([]string, error) {
	if i < 0 || i > len(l)-1 {
		return nil, fmt.Errorf("index(%v) out of range", i)
	}
	return RemoveAt(l, i), nil
}

// Remove returns a copy of the string slice l with the first instances of the string s removed.
// If s is not within l, it returns a copy of the slice.
func Remove(l []string, s string) []string {
	if l == nil {
		return []string{}
	}
	return RemoveAt(l, Index(l, s))
}

// RemoveE returns a copy of the string slice l with the first instances of the string s removed.
// If s is not within l, it reports error.
func RemoveE(l []string, s string) ([]string, error) {
	r := RemoveAt(l, Index(l, s))
	if len(l) == len(r) {
		return nil, fmt.Errorf("'%v' is not found", s)
	}
	return r, nil
}

// RemoveN returns a copy of the string slice l with the first n non-overlapping instances of the string s removed.
// If n < 0, there is no limit on the number of remove.
func RemoveN(l []string, s string, n int) []string {
	if n < 0 {
		n = len(l)
	}
	i := 0
	r := make([]string, len(l))
	for _, v := range l {
		if v != s || n == 0 {
			r[i] = v
			i++
		} else {
			n--
		}
	}
	return r[:i]
}

// RemoveNE returns a copy of the string slice l with the first n non-overlapping instances of the string s removed.
// If n < 0, there is no limit on the number of remove. If no remove, it reports error.
func RemoveNE(l []string, s string, n int) ([]string, error) {
	r := RemoveN(l, s, n)
	if len(l) == len(r) {
		return nil, fmt.Errorf("'%v' is not found", s)
	}
	return r, nil
}

// Uniq returns a copy of the string slice l with duplication instances removed.
func Uniq(l []string) []string {
	if len(l) == 0 {
		return []string{}
	}
	r := append([]string{}, l...)
	sort.Strings(r)
	j := 0
	for i := 1; i < len(r); i++ {
		if r[j] == r[i] {
			continue
		}
		j++
		r[j] = r[i]
	}
	return r[:j+1]
}

// Sub returns new string slice of a - b.
func Sub(a, b []string) []string {
	i := 0
	r := make([]string, len(a))
	for _, v := range a {
		if !Contains(b, v) {
			r[i] = v
			i++
		}
	}
	return r[:i]
}
