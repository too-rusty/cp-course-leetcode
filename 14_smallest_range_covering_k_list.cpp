class Solution {
    
    struct Data {
        int val, idx, lidx;
        Data(int val,int idx,int lidx) : val(val), idx(idx), lidx(lidx) {}
        bool operator<(const Data& other)const {
            if(val < other.val)   return true;
            if(val > other.val)   return false;
            return lidx < other.lidx;
        }
    };
    
public:
    
    vector<int> smallestRange(vector<vector<int>>& nums) {
        set<Data> S;
        const int inf = 2e9;
        vector<int> ans( {0,inf} );
        for(int i=0;i<(int)nums.size();i++) S.insert(Data(nums[i][0],0,i));
        
        while(1){
            auto z=S.begin();
            int tmpa=S.begin()->val, tmpb=S.rbegin()->val;
            
            if(tmpb - tmpa < ans[1] - ans[0])           ans[0]=tmpa,ans[1]=tmpb;
            // if(tmpb-tmpa==ans[1]-ans[0] && tmpa<ans[0]) ans[0]=tmpa,ans[1]=tmpb;
            
            if(z -> idx+1 == (int)nums[z->lidx].size()) break;
            
            Data new_d( nums[ z->lidx ][ z->idx + 1 ] , z->idx + 1 , z->lidx );
            S.erase( z ) ;
            S.insert( new_d ) ;
        }
        return ans;
    }
    
};



/*


l1 [ 0 , 10 , 44 ]
l2 [ 2 , 3 , 32 ]
l3 [ 1 , 23 , 24]


i need to fix the first element and check for all of them
 l , r

0 ,1, 2    0-2
10 ,1, 2   1-10
10, 23, 2  2-23
10, 23, 3  3-23


we can fix the l value and find the lowest r such that one element from all the lists is covered

this is brute force and in many ways is how we can solve a problem, think of the brute force first
after that we can optimise the process by greedily selecting the lowest l value and the r value
i have used sets in cpp for this problem but we can use heaps in golang for the same


*/
