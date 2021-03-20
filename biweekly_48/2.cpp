
//https://leetcode.com/contest/biweekly-contest-48/problems/design-authentication-manager/


class AuthenticationManager {
    map<string,int>expire;
    int ttl = -1;

    set<pair<int,string>>S;

    void process(int currTime) {
        while(!S.empty() && S.begin()->first <= currTime) {
            expire.erase(S.begin()->second);
            S.erase(S.begin());
        }
    }

public:

    AuthenticationManager(int timeToLive) { ttl = timeToLive; }

    void generate(string tokenId, int currentTime) {
        process(currentTime);
        expire[tokenId] = currentTime + ttl;
        S.insert( { expire[tokenId] , tokenId } );
    }

    void renew(string tokenId, int currentTime) {
        process(currentTime);
        if(!expire.count(tokenId)) return;
        S.erase( { expire[tokenId] , tokenId } );
        expire[tokenId] = currentTime + ttl;
        S.insert( { expire[tokenId] , tokenId } );
    }

    int countUnexpiredTokens(int currentTime) {
        process(currentTime);
        return (int)S.size();
    }
};

/**
 * Your AuthenticationManager object will be instantiated and called as such:
 * AuthenticationManager* obj = new AuthenticationManager(timeToLive);
 * obj->generate(tokenId,currentTime);
 * obj->renew(tokenId,currentTime);
 * int param_3 = obj->countUnexpiredTokens(currentTime);

 Some things I am maintaining

 the map and set contains only un expired tokens

 process removes all the expired tokens before the current time

 it is assumed that the operations follow the correct time order



 */
