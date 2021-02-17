type LRUCache struct {
    kv_store            map[int]int
    latest_index        map[int]int     // latest_index -> latest index of the key in the queue
    capacity, volume    int             // to determine wheter we evict or not
    curr_idx            int
    queue               []int
}


func Constructor(capacity int) LRUCache {
    return LRUCache {
        kv_store : make(map[int]int) ,
        latest_index : make(map[int]int) ,
        capacity : capacity ,
        queue : make([]int, 0) ,
    }
}


func (this *LRUCache) Get(key int) int {
    if val, found := this.kv_store[key]; found {
        this.queue = append(this.queue, key)
        this.latest_index[key] = len(this.queue) - 1
        return val
    }
    return -1
}


func (this *LRUCache) Put(key int, value int)  {
    if _, found := this.kv_store[key]; found {
        
        this.queue = append(this.queue, key)
        this.latest_index[key] = len(this.queue) - 1
        
        this.kv_store[key] = value
        
    } else {
        
        this.queue = append(this.queue, key)
        this.latest_index[key] = len(this.queue) - 1
        this.kv_store[key] = value
        
        this.volume++
        
        if this.volume > this.capacity {
            // evict LRU
            defer func(){ this.volume-- }()
            
            for {
                key2 := this.queue[this.curr_idx]
                if this.latest_index[key2] == this.curr_idx {
                    
                    delete(this.latest_index, key2)
                    delete(this.kv_store, key2)
                    
                    
                    this.curr_idx++
                    break
                }
                this.curr_idx++
            }
        }
        
    }
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

/*


we can do this question in quite a few number of ways
i used a queue like structure, (its simply an array and two maps to store the indices)

it may happen that something gets updated and we ignore its previous value in the array ( queue )
and move on to find the latest value, everything before that is removed or deprecated

we are traversing once over the whole array lazily, whenever its required, hence the complexity is because of the map only.




*/
