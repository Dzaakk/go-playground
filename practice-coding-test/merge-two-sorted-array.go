package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	tmp := make([]int, 0, m+n)
	i, j := 0, 0

	for i < m && j < n {
		if nums1[i] < nums2[j] {
			tmp = append(tmp, nums1[i])
			i++
		} else {
			tmp = append(tmp, nums2[j])
			j++
		}
	}

	tmp = append(tmp, nums1[i:m]...)
	tmp = append(tmp, nums2[j:n]...)
	copy(nums1, tmp)
}
