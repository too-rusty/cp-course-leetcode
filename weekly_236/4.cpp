class MKAverage {
private:
    multiset<int> min_heap, max_heap, rest;
    int count, m , k ; long long sum;
    vector<int> elements;
public:
    MKAverage(int m, int k) {
        this->count = 0;
        this->m = m;
        this->k = k;
        this->sum = 0;
    }
    
    void addElement(int num) {
        elements.push_back(num);
        count++;
        if(count <= m) {
            rest.insert(num);
            sum += num;
        }
        
        if(count == m){
            // separate them 
            for(int i = 0; i < k; i++) {
                auto it = rest.begin();
                sum -= *it;
                min_heap.insert(*it);
                rest.erase(it);
            }
  
            for(int i = 0; i < k; i++) {
                auto it = rest.end();
                --it;
                sum -= *it;
                max_heap.insert(*it);
                rest.erase(it);
            }
            

            assert((int)rest.size() > 0);
            
        } else if(count > m){
            assert(2*k + 1 <= m);

            // already separated, first remove the element
            // and readjust the invariance accordingly
            
            int x = elements[count-m-1];
            //remove and adjust the heaps
            if(min_heap.find(x)!=min_heap.end()) {
                //rem from min heap
                auto it = min_heap.find(x);
                min_heap.erase(it);
                min_heap.insert(*rest.begin());
                sum -= *rest.begin();
                rest.erase(rest.begin());
                
            } else if (max_heap.find(x)!=max_heap.end()) {
                //rem from max heap
                auto it = max_heap.find(x);
                max_heap.erase(it);
                it = rest.end(); --it;
                sum -= *it;
                max_heap.insert(*it);
                rest.erase(it);
            } else {
                auto it = rest.find(x);
                sum -= *it;
                rest.erase(it);
            }
            //adjusted
            
            
            //insert the newer element
            int high = *max_heap.begin();
            auto it = min_heap.end(); --it;
            int low = *it;
            
            if (num < low) {
                
                auto it = min_heap.end(); --it;
                sum += *it;
                rest.insert(*it);
                min_heap.erase(it);
                min_heap.insert(num);
                assert((int)min_heap.size() == k);
                
                

            } else if (num > high) {
                
                // should go to max heap
                sum += *max_heap.begin();
                rest.insert(*max_heap.begin());
                max_heap.erase(max_heap.begin());
                max_heap.insert(num);
                assert((int)max_heap.size() == k);
                
            } else {
                rest.insert(num);
                sum += num;
            }
            
        }
    }
    
    int calculateMKAverage() {
        if (count < m) return -1;
        return sum / (m-2*k);
    }
};

/*


visualise this structure

min_first_k_elements | rest_of_array | max_last_k_elements


now if the elements are less than m then return -1

when elements are equal to m then segregate into three parts

note that at this point 2 * k + 1 < m ( according to the constraintes ) 

it means we have atleast 1 element in the middle ( in the rest part )

now when a new element is added to the array , we need to remove
the last mth element ( sliding window manner, add 1 remove 1 , distance m apart)
why, because we need to consider the last m elements only


now the elemnt removed is either from the min_heap, rest or max_heap
wherever it lies

note that if it is removed from the min_heap or the max_heap
then readjust the heaps, i.e. insert one element from the rest arr
to readjust the sizes of the extreme heaps to k

if that is the case , we again have the initial structure,one extra
element in the middle

now after the READJUSTMENT is DONE

try insertin the newer element
if the newer elemnt is smaller than the max of the min_heap
it belongs there, since that heap keeps track of the lowest k elements
if it is greater than the min of the max_heap then keep it there
since it holds the max k elements

else simply put it in the rest elements







*/