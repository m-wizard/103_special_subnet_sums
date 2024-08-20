package main

import (
	"fmt"
	"slices"
)

//  Slow port from Python's cryptic itertools:
//  https://docs.python.org/3/library/itertools.html#itertools.combinations

func gen_all_combs(set []int, r int) map[int][]int {

	curr_comb := []int{}

	final_comb := []int{}

	all_combs := make(map [int][]int)

	if r > len(set) { return all_combs }

	n := 0

	//  Initialise some sets.
	
	for i := range r {

		all_combs[n] = append(all_combs[n], set[i])

		curr_comb = append(curr_comb, i)

		final_comb = append(final_comb, i + len(set) - r)

	}

	final_comb_hit := false

	for final_comb_hit == false {

		curr_comb = iterate_comb(curr_comb, final_comb, r)

		n += 1

		//  Replace combination with unique values provided, and store.

		for j := range r {

			all_combs[n] = append(all_combs[n], set[curr_comb[j]])

		}

		final_comb_hit = slices.Equal(curr_comb, final_comb)
		
	}

	return all_combs
	
}

func iterate_comb(curr_comb []int, final_comb []int, r int) []int {

	//  Find first non-match, iterating from end of combinations, backwards.

	a := len(final_comb) - 1

	found_mismatch := false

	for found_mismatch == false && a >= 0 {

		if curr_comb[a] != final_comb[a] && found_mismatch == false {

			found_mismatch = true

		} else { a -= 1 }

	}

	//  Iterate the current combination, and reset subsequent numbers within it.

	curr_comb[a] += 1

	b := a + 1

	for b < len(curr_comb) {

		curr_comb[b] = curr_comb[b - 1] + 1

		b += 1

	}

	return curr_comb

}


func main() {

	fmt.Println(gen_all_combs([]int{1, 2, 3, 4, 5, 6, 7}, 3))

}
