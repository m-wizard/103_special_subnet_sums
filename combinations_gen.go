package main

import (
	"fmt"
	"slices"
)

//  Slow port from Python's itertools:
//  https://docs.python.org/3/library/itertools.html#itertools.combinations

func combs(set []int, r int) []int {

	fmt.Println("combs()")
	
	if r > len(set) { return []int{} }

	curr_comb := []int{}

	final_comb := []int{}

	all_combs := make(map [int][]int)

	n := 0

	//  Initialise some sets.
	
	for i := range r {

		curr_comb = append(curr_comb, i)

		all_combs[n] = append(all_combs[n], set[i])

		final_comb = append(final_comb, i + len(set) - r)

	}

	fmt.Println(curr_comb, final_comb)

	final_comb_hit := false

	for final_comb_hit == false {

		//  Find first non-match

		a := len(final_comb) - 1

		found_mismatch := false

		for found_mismatch == false && a >= 0 {

			if curr_comb[a] != final_comb[a] && found_mismatch == false {

				found_mismatch = true

			} else { a -= 1 }

			fmt.Println(a)
			
		}

		curr_comb[a] += 1

		b := a + 1

		for b < len(curr_comb) {

			curr_comb[b] = curr_comb[b - 1] + 1

			b += 1

		}

		n += 1

		for c := range r {

			all_combs[n] = append(all_combs[n], set[curr_comb[c]])

		}

		final_comb_hit = slices.Equal(curr_comb, final_comb)
		
	}

	fmt.Println(all_combs, n)

	

	
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

	return set
	
}

func main() {

	combs([]int{1, 2, 3, 4, 5, 6, 7}, 3)

}
