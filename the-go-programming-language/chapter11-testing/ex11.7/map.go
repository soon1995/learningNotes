package main

type M map[int]bool

func (m M) Add(x int) {
	m[x] = true
}

func (m M) UnionWIth(x M) {
	for k := range x {
		m[k] = true
	}
}
