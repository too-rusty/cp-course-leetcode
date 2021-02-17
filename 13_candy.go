func candy(A []int) int {
    n := len(A)
    left, right := make([]int, n), make([]int, n)
    var ans int
    for i:=0;i<n;i++{
        left[i],right[i]=1,1
    }
    for i:=1;i<n;i++{
        if A[i] > A[i-1]{
            left[i]=left[i-1]+1
        }
    }
    for i:=n-2;i>=0;i--{
        if A[i]>A[i+1]{
            right[i]=right[i+1]+1
        }
    }
    max := func(a, b int) int {
        if a > b { return a }
        return b
    }

    for i:=0;i<n;i++{
        ans += max(left[i],right[i])
    }
    return ans
}

/*


fixing a left value fixes everything from left to right
and similarly, fixing a right value fixes everything from right to left

*/
