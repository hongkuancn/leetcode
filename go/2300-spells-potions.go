package main

import (
	"sort"
)

// func successfulPairs(spells []int, potions []int, success int64) []int {
//     res := make([]int, 0)
//     for _, s := range spells {
//         count := 0
//         for _, p := range potions {
//             if int64(s) * int64(p) >= success {
//                 count += 1
//             }
//         }
//         res = append(res, count)
//     }
//     return res
// }

func successfulPairs(spells []int, potions []int, success int64) []int {
	res := make([]int, len(spells))
	sort.Ints(potions)
	m := len(potions)
	for i, s := range spells {
		// slow convertion
		// target := int(math.Ceil(float64(success) / float64(s)))
		target := int((success + int64(s) - 1) / int64(s))
		pos := binarySearch(potions, target)
		res[i] = m - pos
	}
	return res
}

// find the minimun index that the number is not less than target
func binarySearch(arr []int, target int) int {
	l := 0
	r := len(arr) - 1
	for l < r {
		mid := l + (r-l)/2
		if arr[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	if arr[l] < target {
		return len(arr)
	}
	return l
}
