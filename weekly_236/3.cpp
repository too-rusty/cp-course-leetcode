class Solution {
public:
    int solve(int idx,int lane,vector<int>& obstacles, vector<vector<int> >& dp) {
        if(dp[lane][idx] != -1) return dp[lane][idx];
        if (idx == (int)obstacles.size()-1) return 0; // done, the frog reached the end
        int ans=(int)obstacles.size()+10; // ans initially infinite
        for(int i = 1; i < 4; i++) {
            if(obstacles[idx+1] == i || obstacles[idx] == i) continue; // cant go to that lane
            int res = (lane!=i?1:0) + solve(idx+1,i,obstacles, dp); // one jump + recursion
            ans=min(ans, res);
        }
        return dp[lane][idx]=ans;
    }
    int minSideJumps(vector<int>& obstacles) {
        obstacles.push_back(0);
        vector<vector<int> > dp(4, vector<int> ((int)obstacles.size(), -1));
        return solve(0,2,obstacles,dp); // starting from the second lane
    }
};

/*

top down DP

added one extra block at the end to be on the safe side

from [idx,lane] -> [idx+1,lane2]
where lane2 can be any other lane except the one where
there is no blockage

also we can't jump to the lane in the curr index which
has a block so be careful with that too (costed WA)

*/