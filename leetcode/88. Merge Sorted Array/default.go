package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	for n1, n2, l := m-1, n-1, m+n-1; n2 >= 0; l-- {
		if n1 >= 0 && nums1[n1] > nums2[n2] {
			nums1[l] = nums1[n1]
			n1--
		} else {
			nums1[l] = nums2[n2]
			n2--
		}
	}
}
