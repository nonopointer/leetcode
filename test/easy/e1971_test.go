package test

import "testing"

func Test1971(t *testing.T) {

}

func validPath(n int, edges [][]int, source int, destination int) bool {
	uf := make([]int, n)
	for i, _ := range uf {
		uf[i] = i
	}

	for _, v := range edges {
		a, b := v[0], v[1]
		union(a, b, uf)
	}

	return find(source, uf) == find(destination, uf)
}

func union(a int, b int, uf []int) {
	a = find(a, uf)
	b = find(b, uf)
	uf[b] = a
}

func find(x int, uf []int) int {
	for x != uf[x] {
		x = uf[x]
	}
	return uf[x]
}
