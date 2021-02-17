class Solution {
public:
    vector<vector<int>> threeSum(vector<int>& nums) {
        vector<vector<int>>ans;
        int n=(int)nums.size();
        if(n==0)return ans;
        sort(nums.begin(),nums.end());
        unordered_map<int,bool> seen;
        map<pair<int,pair<int,int>>, bool> triplet;
        seen[nums[0]]=true;
        map<int,int>cnt;
        for(int i = 1;i < n; i++){
            //fix index i and check for all j
            if (cnt.count(nums[i]) && cnt[nums[i]]==3)continue;
            for(int j = i + 1; j < n; j++) {
                int sum = nums[i] + nums[j];
                if(seen.count(-sum)) {
                    vector<int> tmp({-sum,nums[i],nums[j]});
                    sort(tmp.begin(),tmp.end());
                    triplet[{tmp[0],{tmp[1],tmp[2]}}]=true;
                }
            }
            //done with i
            seen[nums[i]]=true;
            cnt[nums[i]]++;
        }
        // i.first = key , i.second = value (true)
        for(auto i:triplet) {
            pair<int,pair<int,int>>key=i.first;
            vector<int> tmp({key.first, key.second.first, key.second.second});
            ans.push_back(tmp);
        }
        return ans;
        
        
//         sort(nums.begin(),nums.end());
//         // for(int i :nums)cout<<i<< " ";
//         // cout<<endl;
//         for (int i = 0; i < n; i++) {
//             int l = i+1;
//             int r = n-1;
            
//             if(i>0&&nums[i]==nums[i-1])continue;
            
//             while(l < r) {
                
//                 int sum = nums[l] + nums[r] + nums[i];
//                 // cout<<sum<<endl;
//                 if ( sum == 0 ) {
                    
//                     vector<int> tmp({nums[i],nums[l],nums[r]});
//                     ans.push_back(tmp);
//                     ++l;
//                     --r;

//                 } else if ( sum > 0 ) {
//                     --r;
//                 } else {
//                     ++l;
//                 }
//             }
//         }
//         sort(ans.begin(),ans.end());
//         if((int)ans.size()==0)return ans;
//         vector<vector<int>>res;
//         res.push_back(ans[0]);
//         for (auto& z:ans)if(z!=res.back())res.push_back(z);
//         return res;
    }

};

/*

0 0 0 0
{ 0 0 0 }  x 3
-1 , 0 , 1
1 , 0 , -1


well this question took quite some time

i was missing the fact that we can have 0 0 0 values
i.e. similar values in a row



so no need to repeat for values that are already seen

note the method in the comments, we can use the two pointer approach to solve the problem too
and we will be using this approach extensively later on

*/
