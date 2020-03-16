// https://leetcode.com/problems/merge-sorted-array
// Just sol to the problem, It does not include the I/O part

func merge(nums1 []int, m int, nums2 []int, n int)  {

    k := m + n -1
    m, n = m-1, n-1

    for m >= 0 && n >= 0 {
        if nums1[m] > nums2[n] {
            nums1[k] = nums1[m]
            m--
        } else {
            nums1[k] = nums2[n]
            n--
        }
        k--
    }

    for m >= 0 {
        nums1[k] = nums1[m]
        k--
        m--
    }

    for n >= 0 {
        nums1[k] = nums2[n]
        k--
        n--
    }
}
