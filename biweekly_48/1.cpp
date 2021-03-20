class Solution {
public:
    int secondHighest(string s) {
        set<int>S;
        for(char c : s) {
            if(c>='0' && c<='9'){
                S.insert(-c+48);
            }
        }
        if(!S.empty())S.erase(S.begin());
        return S.empty() ? -1 : -1 * *S.begin();
    }
};



/*
 *
 * this is a simple question, no need for explnation
 *
 * */
