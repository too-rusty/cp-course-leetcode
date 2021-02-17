func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    n, m := len(obstacleGrid), len(obstacleGrid[0])
    dp := make([][]int, n+1)
    for i := 0; i < n + 1; i++ { dp[i] = make([]int, m+1) }
    if obstacleGrid[0][0] == 0 { dp[1][1] = 1 }
    
    // row 1
    for j := 2; j < m + 1; j++ {
        dp[1][j] = dp[1][j-1]
        if obstacleGrid[0][j-1] == 1 {
            dp[1][j] = 0;
        }
    }
    // col 1
    for i := 2; i < n + 1; i++ {
        dp[i][1] = dp[i-1][1]
        if obstacleGrid[i-1][1-1] == 1 {
            dp[i][1] = 0;
        }
    }
    
    for i := 2; i < n + 1; i++ {
        for j := 2; j < m + 1; j++ {
            dp[i][j] = dp[i-1][j] + dp[i][j-1]
            if obstacleGrid[i-1][j-1] == 1 {
                dp[i][j] = 0;
            }
        }
    }
    return dp[n][m]
    
}

/*

n = 1 , m = 5
- - X - -
- X - - -


1 1 0 0 0
1 0 0 0 0
*/
