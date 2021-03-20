class Solution {
public:
    int count_bits(int x) {
        int cnt=0;
        while(x > 0) {
            if(x&1)++cnt;
            x >>= 1;
        }
        return cnt;
    }
    
    int maxScore(vector<int>& nums) {
        int n = (int)nums.size(); n >>= 1;
        vector< vector< int > > dp(n+1, vector<int> (1<<(2*n)) );
        
        for(int i = 1; i < n + 1; i++ ) {
            for(int mask = 0; mask < (1 << (2*n)); mask++) {
                if(count_bits(mask) == 2*i-2) {
                    
                    for(int j = 0; j < 2*n; j++) {
                        for(int k = 0; k < 2*n; k++) {
                            if(!(mask&(1<<j)) && !(mask&(1<<k))) {
                                //take these two nums
                                dp[i][mask|(1<<j)|(1<<k)] = max(dp[i][mask|(1<<j)|(1<<k)],
                                                               dp[i-1][mask] + i*__gcd(nums[j], nums[k]) );
                            }
                        }
                    }
                }
            }
        }
        
        return dp[n][(1<<(2*n))-1];
        
    }
};


/*

final answer is dp[n][(1<<2n)-1] when we have picked everything

the number of set bits in ith turn must be 2*i-2

this is a dp + bitmasking question



https://leetcode.com/contest/biweekly-contest-48/problems/maximize-score-after-n-operations/


*/
