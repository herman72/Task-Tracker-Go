// Leetcode Q 88, merge sorted array
// package main

// import (
//     "fmt"
// )

// func main() {

// 	nums1 := []int{1,2,3,0,0,0}
// 	m := 3
// 	nums2 := []int{2,5,6}
// 	n := 3
// 	merge(nums1, m, nums2, n)

// }

// func merge(nums1 []int, m int, nums2 []int, n int)  {
//     nums1 = nums1[:m] // Take only the first m elements
//     nums1 = append(nums1, nums2...) // Append nums2 elements
//     fmt.Println(nums1)
//     for i :=0; i < m + n ; i++{
//         for j :=i; j < m+n; j++{

//             if nums1[i] >= nums1[j] {

//                 nums1[i], nums1[j] = nums1[j], nums1[i]

//                 // i = j
//             }

//         }

//     }
//     fmt.Println(nums1)
// }

// Leetcode Q 27, remove element

package main

import "fmt"

func main(){
    nums := []int{0,1,2,2,3,0,4,2}
    var val int = 2
     k := removeElement(nums, val)
     fmt.Println(k)

}

func removeElement(nums []int, val int) int {

    // var listLen int = len(nums)
    // for i:=0;  i < len(nums); i++{

    //     if nums[i] == val {
    //         nums = append(nums[:i], nums[i+1:]...)
    //         i--

    //     }

    // }
    // return len(nums)
    k := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != val {
            nums[k] = nums[i]
            k++
        }
    }
    return k
    
}