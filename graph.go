package binrels

type Graph map[string]set

// The function add adds the relationship "a is related to b" to the
// graph. This function is idempotent.
func (g Graph) Add(a, b string) {
	related, ok := g[a]
	if !ok {
		related = make(set)
	}
	related.add(b)
}

// The function get returns if a is related to b in the graph g.
func (g Graph) Get(i, j string) bool {
	rels, ok := g[i]
	if !ok {
		return false
	}
	_, ok = rels[j]
	return ok
}
