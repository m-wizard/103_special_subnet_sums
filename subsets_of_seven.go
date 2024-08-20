package main

import "fmt"

//  Fast port from Python's itertools:
//  https://docs.python.org/3/library/itertools.html#itertools.combinations

func combs(set []int, r int) []int {
	
	n := len(set)
	
	if r > n { return }
	
	yield tuple(set[i] for i in set)
	while True:
        for i in reversed(range(r)):
	if indices[i] != i + n - r:
	break
else:
	return
        indices[i] += 1
        for j in range(i+1, r):
	indices[j] = indices[j-1] + 1
        yield tuple(set[i] for i in indices)
	
}

func main() {

	combs([]int{1, 2, 3, 4, 5, 6, 7}, 2)

}
