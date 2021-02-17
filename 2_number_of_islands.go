
func dfs(i,j int, grid [][]byte) {
    n,m := len(grid), len(grid[0])
    if i<0 || j<0 || i>=n || j>=m { return }
    if grid[i][j] == '0' || grid[i][j] == '2' { return }
    grid[i][j] = '2'
    dfs(i+1,j,grid)
    dfs(i-1,j,grid)
    dfs(i,j+1,grid)
    dfs(i,j-1,grid)
}

func numIslands(grid [][]byte) int {
    var islands int = 0
    n,m := len(grid), len(grid[0])
    for i:=0; i<n; i++ {
        for j:=0; j<m; j++ {
            if grid[i][j] == '1' {
                dfs(i,j,grid)
                islands++
            }
        }
    }
    return islands
}


/*

same code but in golang


*/
