class Solution {
private:
    void dfs(int i,int j, vector<vector<char>>& grid){
        int n = (int)grid.size();
        int m = (int)grid[0].size();
        if(i < 0 || j < 0 || i >= n || j >= m) return ;
        if(grid[i][j]=='0'||grid[i][j]=='2') return ;
        grid[i][j] = '2';
        for(int dx = -1; dx <= 1; ++dx) {
            for(int dy = -1; dy <= 1; ++dy){
                if(dx*dy==0){
                    dfs(i+dx,j+dy,grid);
                }
            }
        }
    }
public:
    int numIslands(vector<vector<char>>& grid) {
        int n = (int)grid.size();
        int m = (int)grid[0].size();
        int islands=0;
        for(int i = 0; i < n; i++) {
            for(int j = 0; j < m; j++) {
                if(grid[i][j] == '1') {
                    // we encountered an unclaimed
                    // claim it 
                    dfs(i,j,grid);
                    islands++;
                }
            }
        }
        return islands;
    }
};

/*


Time complexity
---------------
O(m*n +  no of ones )

10
01


The idea is to find all the ones connected to a 1
if the current 1 is there then it must be a different component else had it been part of some previous component
it would have become 0



start dfs from the one and visit all the ones in the component , repeat for all the ones ( since they are part of an altogether diff component)
and keep the count of number of times we do this process





*/
