func merge(nums1 []int, m int, nums2 []int, n int)  {
    
    tmp := make([]int, m+n)
    
    var i,j,k int
    for i<m && j<n{
        if nums1[i] < nums2[j] {
            tmp[k]=nums1[i]
            i++
        } else {
            tmp[k]=nums2[j]
            j++
        }
        k++
    }
    for i<m{
        tmp[k]=nums1[i]
        i++
        k++
    }
    for j<n{
        tmp[k]=nums2[j]
        j++
        k++
    }
    
    
    copy(nums1, tmp)
    
}

/*

here we merge two sorted arrays
this is an easy and a pretty intuitive problem

*/
