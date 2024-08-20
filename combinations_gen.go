package main

import (
	"fmt"
	"errors"
)

//  Fast port from Python's itertools:
//  https://docs.python.org/3/library/itertools.html#itertools.combinations

func combs(set []int, r int) ([]int, error) {

	fmt.Println("combs()")
	
	if r > len(set) { return []int{}, errors.New("test") }

	indices := []int{}

	all_combs := make(map [int][]int)

	n := 0
	
	for i := range r {

		indices = append(indices, i)

		all_combs[n] = append(all_combs[n], set[i])

	}

	fmt.Println(indices, all_combs[n])

	
// 	yield tuple(set[i] for i in set)
// 	while True:
//         for i in reversed(range(r)):
// 	if indices[i] != i + n - r:
// 	break
// else:
// 	return
//         indices[i] += 1
//         for j in range(i+1, r):
// 	indices[j] = indices[j-1] + 1
	//         yield tuple(set[i] for i in indices)

	return set, nil
	
}

func main() {

	combs([]int{1, 2, 3, 4, 5, 6, 7}, 3)

}
