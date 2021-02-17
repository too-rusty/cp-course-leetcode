class Solution {
public:
    int candy(vector<int>& ratings) {
        int n = (int)ratings.size();
        vector<pair<int,int>>A;
        for(int i=0;i<n;i++){
            A.push_back({ratings[i],i});
        }
        sort(A.begin(),A.end());
        vector<int> ans(n,0);
        for(auto i : A){
            int idx=i.second;
            int mx=0;
            
            if(idx>0 && ratings[idx]>ratings[idx-1]){
                mx=max(ans[idx-1],mx);
            }
            if(idx<n-1 && ratings[idx]>ratings[idx+1]){
                mx=max(ans[idx+1],mx);
            }
            ans[idx]=mx+1;
            // cout<<"mx "<<ans[idx]<<" i.se " <<i.second<<endl;
        }
        // for(int i=0;i<n;i++)cout<<ans[i]<<" ";
        // cout<<endl;
        int res = accumulate(ans.begin(),ans.end(),0);
        return res;
    }
};

/*

this is a lengthier implementation but more intuitive

we allocate the cady to the lowest height guy and then some more to the guy with higher height
and so on


*/
