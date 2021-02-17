type MinStack struct {
    stack, pref_min []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    minStack := new(MinStack)
    minStack.stack = make([]int,0)
    minStack.pref_min = make([]int,0)
    return *minStack
}



func (this *MinStack) Push(x int)  {
    min := func(x, y int) int {
        if x < y { return x }
        return y
    }
   
    if len(this.stack) == 0 {
        this.stack = append(this.stack, x)
        this.pref_min = append(this.pref_min, x)
        return
    }
    
    n := len(this.pref_min)
    this.stack = append(this.stack, x)
    this.pref_min = append(this.pref_min, min(x, this.pref_min[n-1]))
}


func (this *MinStack) Pop()  {
    if len(this.stack) > 0 {
        this.stack = this.stack[:len(this.stack)-1]
        this.pref_min = this.pref_min[:len(this.pref_min)-1]
    }
}


func (this *MinStack) Top() int {
    n := len(this.stack)
    return this.stack[n-1]
}


func (this *MinStack) GetMin() int {
    n := len(this.pref_min)
    return this.pref_min[n-1]
}



/*



-------->

1 2 5 4 10
1 1 1 1 1

pref_min[i] = the min value in the prefix till pos i

the idea is the same as some prefix sum , instead we keep a prefix min
we store the information in a stack

So we have 2 stacks, one for storing the numbers encountered till now
and the other is the min till now



*/




/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
