package binrels

import "testing"

func TestNew(t *testing.T) {
	s := newSet()
	if s == nil {
		t.Error("newSet returned nil")
	}
}

type equalTest struct {
	i1 set
	i2 set
	e  bool
}

var (
	empty1     = newSet()
	empty2     = newSet()
	equalTests = [...]equalTest{
		{i1: empty1, i2: empty1, e: true},
		{i1: empty1, i2: empty2, e: true},
		{i1: empty2, i2: empty1, e: true},
		{i1: empty1, i2: newSet().add("b"), e: false},
		{i1: newSet().add("a"), i2: empty1, e: false},
		{i1: newSet().add("a"), i2: newSet().add("b"), e: false},
		{i1: newSet().add("a").add("b"), i2: newSet().add("b"), e: false},
		{i1: newSet().add("a").add("b"), i2: newSet().add("b").add("a"), e: true},
		{i1: newSet().add("a", "b"), i2: newSet().add("b", "a"), e: true},
		{i1: newSet().add("a", "b", "c"), i2: newSet().add("b").add("c", "a"), e: true},
	}
)

func TestCmp(t *testing.T) {
	for i, v := range equalTests {
		o := v.i1.equal(v.i2)
		if o != v.e {
			t.Errorf("%d) expected %t obtained %t", i, v.e, o)
		}
	}
}

type addTest struct {
	start    set
	toAdd    []string
	expected set
}

var (
	addTests = [...]addTest{
		{
			start:    newSet(),
			toAdd:    []string{},
			expected: newSet(),
		},
		{
			start: newSet(),
			toAdd: []string{"a"},
			expected: map[string]struct{}{
				"a": struct{}{},
			},
		},
		{
			start: newSet(),
			toAdd: []string{"a", "b"},
			expected: map[string]struct{}{
				"a": struct{}{},
				"b": struct{}{},
			},
		},
		{
			start: newSet().add("a"),
			toAdd: []string{},
			expected: map[string]struct{}{
				"a": struct{}{},
			},
		},
		{
			start: newSet().add("a"),
			toAdd: []string{"b"},
			expected: map[string]struct{}{
				"a": struct{}{},
				"b": struct{}{},
			},
		},
		{
			start: newSet(),
			toAdd: []string{"a", "a"},
			expected: map[string]struct{}{
				"a": struct{}{},
			},
		},
		{
			start: newSet(),
			toAdd: []string{"a", "a", "a", "b", "a", "a", "b", "c", "a", "b"},
			expected: map[string]struct{}{
				"a": struct{}{},
				"b": struct{}{},
				"c": struct{}{},
			},
		},
	}
)

func TestAdd(t *testing.T) {
	for i, v := range addTests {
		v.start.add(v.toAdd...)
		if !v.expected.equal(v.start) {
			t.Errorf("%d)\n\texpected %v\n\tobtained %v", i, v.expected, v.start)
		}
	}
}

type delTest struct {
	start    set
	toDel    []string
	expected set
}

var (
	delTests = [...]delTest{
		{
			start:    newSet(),
			toDel:    []string{},
			expected: newSet(),
		},
		{
			start:    newSet().add("a"),
			toDel:    []string{},
			expected: newSet().add("a"),
		},
		{
			start:    newSet().add("a", "b"),
			toDel:    []string{},
			expected: newSet().add("a", "b"),
		},
		{
			start:    newSet(),
			toDel:    []string{"a", "b"},
			expected: newSet(),
		},
		{
			start:    newSet().del("a"),
			toDel:    []string{},
			expected: newSet(),
		},
		{
			start:    newSet().add("a"),
			toDel:    []string{"b"},
			expected: newSet().add("a"),
		},
		{
			start:    newSet().add("a"),
			toDel:    []string{"a", "a"},
			expected: newSet(),
		},
		{
			start:    newSet().add("a", "b", "c", "d", "e"),
			toDel:    []string{"a", "c", "a", "e", "c"},
			expected: newSet().add("b", "d"),
		},
		{
			start:    newSet().add("a", "b", "c", "d", "e"),
			toDel:    []string{"a", "c", "a", "e", "c", "bla", "e", "a", "b", "d"},
			expected: newSet(),
		},
	}
)

func TestDel(t *testing.T) {
	for i, v := range delTests {
		v.start.del(v.toDel...)
		if !v.expected.equal(v.start) {
			t.Errorf("%d)\n\texpected %v\n\tobtained %v", i, v.expected, v.start)
		}
	}
}
