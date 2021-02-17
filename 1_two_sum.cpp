class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        vector<int> ans;
        int n = (int)nums.size();
        map<int,int>idx;
        for(int i = 0; i < n; ++i) {
            if(idx.count(target-nums[i])) {
                //found the answer
                ans.push_back(idx[target-nums[i]]);
                ans.push_back(i);
            }
            idx[nums[i]]=i;
        }
        return ans;
        
    }
};


/*

a + b = required_sum

b = required_sum - a

so if for any number a, if b is present (or seen previosuly) we are SAYING , " ok we found the answer lets push it in the vector and return it "
therefore we hash the results for every number and check when we encounter something which was already seen

so two things that need to be kept in mind, for the current number, 
check if a number corresponding to a number is seen and hash the current number



*/
