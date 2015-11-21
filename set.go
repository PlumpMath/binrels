package binrels

import "reflect"

// An unordered collection of non-duplicated elements.
type set map[string]struct{}

func newSet() set {
	return make(set)
}

// Returns if set s is equal to set o.
// Two sets are equal if the have the same elements.
// This function is commutative.
func (s set) equal(o set) bool {
	return reflect.DeepEqual(s, o)
}

// Adds an element e to the set s.
// Returns the modified set, so the call can be chained.
// This function is idempotent.
func (s set) add(e ...string) set {
	for _, v := range e {
		s[v] = struct{}{}
	}
	return s
}

// Remove element e from the set s.
// Returns the modified set, so the call can be chained.
// This function is idempotent.
func (s set) del(e ...string) set {
	for _, v := range e {
		delete(s, v)
	}
	return s
}
