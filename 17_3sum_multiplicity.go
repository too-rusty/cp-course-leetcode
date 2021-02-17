func threeSumMulti(arr []int, target int) (ans int) {

    cnt := make(map[int]int)
    n := len(arr)
    for i := n-2; i>=0; i-- {
        required_sum := target - arr[i]
        if tot,found := cnt[required_sum]; found { ans = add(ans, tot) }
        var j int = i
        for k := j + 1; k < n; k++ { cnt[arr[j] + arr[k]]++ }
    }
    return
    
}

const mod = int(1e9) + 7
func add(a , b int) int {
    a += b
    a %= mod
    return a
}


/*


ai + aj + ak = target

aj + ak = target - ai

same as the previous 2 questions
only thing is that we are fixing ai
and also hashing the results like 2sum


*/
