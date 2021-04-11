func findTheWinner(n int, k int) int {
    var curr int
    done := make([]bool, n)
    for times := 0; times < n - 1; times++ {
        var K int = k
        for K > 0 {
            if done[curr] {
            } else {
                K--
            }
            curr++
            curr %= n
        }
        done[(curr-1+n)%n] = true
        
    }
    for i:=0; i < n; i++ {
        if !done[i] {
            return i + 1
        }
    }
    return -1
    
}

/*

This is bruteforce

everytime remove the kth player

and start from the next one

here , i m not worrying about the exact player

just remove the kth player and shift the curr by 1 so that
we can find the next set of k unset players in the next iteration

be careful with curr index as it has been increased before breaking out



*/