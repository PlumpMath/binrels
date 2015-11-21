package tclosure

import (
	. "github.com/alcortesm/binrels"
)

type tclosure func(g Graph) Graph

func Warshall(g Graph) Graph {
	for k := range g {
		for i := range g {
			for j := range g {
				if g.Get(i, j) || g.Get(i, k) && g.Get(k, j) {
					g.Add(i, j)
				}
			}
		}
	}
	return g
}
