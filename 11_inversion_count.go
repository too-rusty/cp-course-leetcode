
func isIdealPermutation(A []int) bool {
    var n int = len(A)
    var local_inv int

    for i:=1;i<n;i++{
        if A[i-1]>A[i] { local_inv++ }
    }
    global_inv := mergeSort(A,0,n-1)
    
    fmt.Println(local_inv, global_inv)
    return local_inv == global_inv
}

func mergeSort(A []int, l, r int) int {
    if l >= r { return 0 }
    mid := (l + r )/2
    cnt1 := mergeSort(A, l, mid)
    cnt2 := mergeSort(A, mid+1, r)
    cnt3 := merge(A, l, r)
    return cnt1 + cnt2 + cnt3
}


func merge(a []int, l, r int) int {
    mid := (l+r)/2
    i, j, k := l, mid + 1, 0
    tmp := make([]int, r-l+1)
    var inversions int
    for i<mid+1 && j<r+1 {
        if a[i] <= a[j] {
            tmp[k] = a[i]
            inversions += j-(mid+1)
            i++
        } else {
            tmp[k] = a[j]
            j++
        }
        k++
    }
    for j<r+1{
        tmp[k]=a[j]
        j++
        k++
    }
    for i<mid+1{
        tmp[k]=a[i]
        inversions += j - (mid+1)
        i++
        k++
    }
    // assert k == r-l+1
    for i:=0; i<k; i++{
        a[i+l]=tmp[i]
    }
    return inversions
    
}






/*

observation 1 ->   0 <= ai < N , permutation



1 0 -> 1 inversion

3 2 1 -> 3 global inversions , 2 local inversions

i < j , a[i] > a[j] - condition for an inversion (global)


we count the number of inversions while calling the merge function

this is simply a variation of mergesort and while merging, we count the number of inversions



*/
