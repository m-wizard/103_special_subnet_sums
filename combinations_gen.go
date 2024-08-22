package main

import (
	"fmt"
	"slices"
	"strconv"
)

func main() {

	base_elems := []int{1, 2, 3, 4, 5, 6, 7}

	all_combs_of_seven := make([][]int, 0, 128)

	for i := 1 ; i < 7 ; i++ {

		// fmt.Println(i)

		all_combs_of_seven = append(all_combs_of_seven, gen_comb_set(base_elems, i, make([][]int, 0, 128))...)

	}
	
	all_combs_of_seven = append(all_combs_of_seven, base_elems)

	//  To re-use gen_comb_set(), convert output to compatible type.

	acs_ints := convert_all_slices_to_ints(all_combs_of_seven)

	acs_choose_two := gen_comb_set(acs_ints, 2, make([][]int, 0, 8010))

	find_disjoint_sets(acs_choose_two)

	// fmt.Println(acs_choose_two)

	//  127_c_2_combs = make([][]int, 0, 8002)
	//  fmt.Println(gen_comb_set([]int{1, 2, 3, 4, 5, 6, 7}, 3))

}

//  Ported from Python's barely readable itertools:
//  https://docs.python.org/3/library/itertools.html#itertools.combinations

func gen_comb_set(set []int, r int, combs_set [][]int) [][]int {

	curr_comb := []int{}

	curr_comb_overlay := []int{}
	
	final_comb := []int{}

	if r > len(set) { return combs_set }

	n := 0

	//  Initialise some sets.
	
	for i := range r {

		curr_comb = append(curr_comb, i)

		curr_comb_overlay = append(curr_comb_overlay, set[curr_comb[i]])

		final_comb = append(final_comb, i + len(set) - r)

	}

	combs_set = append(combs_set, curr_comb_overlay)

	final_comb_hit := false

	for final_comb_hit == false {

		curr_comb = iterate_comb(curr_comb, final_comb, r)

		n += 1

		//  Replace combination with unique values provided, and store.

		curr_comb_overlay := []int{}

		for j := range r {

			curr_comb_overlay = append(curr_comb_overlay, set[curr_comb[j]])

		}

		combs_set = append(combs_set, curr_comb_overlay)

		final_comb_hit = slices.Equal(curr_comb, final_comb)
		
	}

	return combs_set
	
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

func convert_all_slices_to_ints(set [][]int) []int {

	all := make([]int, 0, 8002)

	// fmt.Println(set)

	for _, v := range set {

		all = append(all, convert_digits_to_int(v))

	}

	return all
	
}

func convert_digits_to_int(set []int) int {

	//  Ripped from Stack Overflow in lieu of standard library functionality.
	//  https://stackoverflow.com/questions/44729587

	res := 0
	
	op := 1
	
	for i := len(set) - 1; i >= 0; i-- {
		
		res += set[i] * op
		
		op *= 10
		
	}
	
	return res

}

func find_disjoint_sets(sets [][]int) [][]int {

	disjoint_sets := make([][]int, 0, 8010)

	for _, v := range sets {

		fmt.Println(v)
		
	}

	return disjoint_sets
	
}

func chk_for_same_digits(set []int) {

	

}
