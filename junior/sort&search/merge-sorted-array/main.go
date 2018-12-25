package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	idx1, idx2, rzt := 0, 0, []int{}
	for idx1 < m || idx2 < n {
		if idx1 == m {
			rzt = append(rzt, nums2[idx2])
			idx2++
		} else if idx2 == n {
			rzt = append(rzt, nums1[idx1])
			idx1++
		} else {
			n1, n2 := nums1[idx1], nums2[idx2]
			if n1 <= n2 {
				rzt = append(rzt, nums1[idx1])
				idx1++
			} else {
				rzt = append(rzt, nums2[idx2])
				idx2++
			}
		}
	}
	for k, v := range rzt {
		nums1[k] = v
	}
}

func main() {

}
