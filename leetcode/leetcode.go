
package main

import (
    "fmt"
)

func main() {

	nums1 := []int{1,2,3,0,0,0}
	m := 3 
	nums2 := []int{2,5,6}
	n := 3
	merge(nums1, m, nums2, n)

}

func merge(nums1 []int, m int, nums2 []int, n int)  {
    nums1 = nums1[:m] // Take only the first m elements
    nums1 = append(nums1, nums2...) // Append nums2 elements
    fmt.Println(nums1)
    for i :=0; i < m + n ; i++{
        for j :=i; j < m+n; j++{

            if nums1[i] >= nums1[j] {
                
                nums1[i], nums1[j] = nums1[j], nums1[i]

                // i = j 
            }

        }
        

    }
    fmt.Println(nums1)
}