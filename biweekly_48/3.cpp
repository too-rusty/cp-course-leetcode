class Solution {
public:
    int getMaximumConsecutive(vector<int>& coins) {
        int mns = 0, mxs = 0;
        sort(coins.begin(), coins.end());
        for(int i : coins) {
            int ni = mns + i;
            int nj = mxs + i;
            if(ni<=mxs+1)mxs=max(nj,mxs);
            else break;
        }
        return mxs-mns+1;
    }
};


/*
 *
 * i was quite confused in this question
 * basically maintain a range which we can make the sums
 * and merge the new range to the older range
 *
 *
 * https://leetcode.com/contest/biweekly-contest-48/problems/maximum-number-of-consecutive-values-you-can-make/
 *
 * */
