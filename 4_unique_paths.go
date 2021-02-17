func uniquePaths(n int, m int) int {
    
    dp := make([][]int, n+1)
    for i := 0; i < n + 1; i++ { dp[i] = make([]int, m+1) }
    
    for i := 1; i < n + 1; i++ {
        for j := 1; j < m + 1; j++ {
            if i == 1 || j == 1 { dp[i][j] = 1 }
        }
    }
    
    for i := 2; i < n + 1; i++ {
        for j := 2; j < m + 1; j++ {
            dp[i][j] = dp[i-1][j] + dp[i][j-1]
        }
    }
    return dp[n][m]
    
}

/*


m X n

Right or Down


dp[i][j] -> ways to reach pos i,j

dp[i][j] = dp[i][j-1] + dp[i-1][j]

base conditions

 1 1 1 1 1
 1 2 - - -
 1 - - - -


*/
