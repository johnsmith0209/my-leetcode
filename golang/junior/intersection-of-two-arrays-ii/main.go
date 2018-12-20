package main

import "fmt"

// fast sort
func sort(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return nums
	}
	mid, mids, bigger, smaller := nums[0], []int{}, []int{}, []int{}
	for _, v := range nums {
		if v > mid {
			bigger = append(bigger, v)
		} else if v < mid {
			smaller = append(smaller, v)
		} else {
			mids = append(mids, v)
		}
	}
	bigger = sort(bigger)
	smaller = sort(smaller)
	return append(smaller, append(mids, bigger...)...)
}

func intersect(nums1 []int, nums2 []int) []int {
	fmt.Println("before sort", nums1, nums2)
	nums1 = sort(nums1)
	nums2 = sort(nums2)
	fmt.Println("after sort", nums1, nums2)
	var rzt []int
	longerNums, shorterNums := nums1, nums2
	if len(nums1) < len(nums2) {
		longerNums, shorterNums = nums2, nums1
	}
	lL, sL, lIdx, sIdx := len(longerNums), len(shorterNums), 0, 0
	for lIdx < lL && sIdx < sL {
		l, s := longerNums[lIdx], shorterNums[sIdx]
		// fmt.Println("lIdx, l", lIdx, l, "sIdx, s", sIdx, s)
		if l == s {
			rzt = append(rzt, l)
			lIdx++
			sIdx++
			// fmt.Println("equal, next ", lIdx, sIdx)
		} else if l < s {
			lIdx++
			// fmt.Println("l < s, next ", lIdx, sIdx)
		} else {
			sIdx++
			// fmt.Println("l > s, next ", lIdx, sIdx)
		}
	}
	return rzt
}

func intersect2(nums1, nums2 []int) []int {
	fmt.Println(nums1)
	fmt.Println(nums2)
	var rzt []int
	l1, l2, idx1, idx2 := len(nums1), len(nums2), 0, 0
	for idx1 < l1 && idx2 < l2 {
		n1, n2, n1EIdx, n2EIdx, step1, step2 := nums1[idx1], nums2[idx2], -1, -1, 0, 0
		fmt.Println(n1, n2, idx1, idx2)
		if n1 == n2 {
			fmt.Println("same", n1, "goto next", idx1+1, idx2+1)
			rzt = append(rzt, n1)
			idx1++
			idx2++
		}
		fmt.Println("1 here")
		for i := idx2; i < l2; i++ {
			step1++
			if nums2[i] == n1 {
				n1EIdx = i
				break
			}
		}
		fmt.Println("2 here")
		for i := idx1; i < l1; i++ {
			step2++
			if nums1[i] == n2 {
				n2EIdx = i
			}
		}
		fmt.Println("3 here")
		// 都没找到相同
		if n1EIdx == -1 && n2EIdx == -1 {
			fmt.Println("no same, break")
			n1EIdx++
			break
		}
		fmt.Println("4 here", n1EIdx, n2EIdx, step1, step2)
		if (n1EIdx <= n2EIdx && n1EIdx >= 0) || (n1EIdx > 0 && n2EIdx < 0) {
			rzt = append(rzt, nums2[n1EIdx])
			idx2 = n1EIdx
			fmt.Println("same", nums2[n1EIdx], "goto next", idx1+1, idx2+1)
			idx1++
			idx2++
		} else if (n1EIdx > n2EIdx && n2EIdx >= 0) || (n2EIdx > 0 && n1EIdx < 0) {
			rzt = append(rzt, nums1[n2EIdx])
			idx1 = n2EIdx
			fmt.Println("same", nums1[n2EIdx], "goto next", idx1+1, idx2+1)
			idx1++
			idx2++
		}
		fmt.Println("5 here")
	}
	return rzt
}

func main() {
	a := []int{4, 6, 1, 21, 4, 5125, 612, 3, 4161, 3, 5134, 123, 312}
	b := []int{1, 123, 3, 123, 12, 4, 312, 3, 24, 4}
	// fmt.Println(sort(a))
	fmt.Println(intersect(a, b))
}
