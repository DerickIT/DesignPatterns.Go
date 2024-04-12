package main

import "fmt"

/*
输入: nums1 = [1,2,2,1], nums2 = [2,2]

输出: [2,2]



*/

func main() {

	arr := intersect([]int{12, 3, 4, 1, 2}, []int{23, 5, 2, 3, 4})
	fmt.Println(arr)

}
func intersect(nums1 []int, nums2 []int) []int {
	m0 := map[int]int{}
	for _, v := range nums1 {
		//遍历nums1，初始化map
		m0[v] += 1
	}
	k := 0
	for _, v := range nums2 {
		//如果元素相同，将其存入nums2中，并将出现次数减1
		if m0[v] > 0 {
			m0[v] -= 1
			nums2[k] = v
			k++
		}
	}
	return nums2[0:k]
}
