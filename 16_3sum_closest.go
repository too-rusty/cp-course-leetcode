import "sort"
func threeSumClosest(nums []int, target int) int {
    
    sort.Ints(nums)
    n := len(nums)
    var ans int
    for i := 0; i < 3; i++ { ans += nums[i] }
    diff := abs(target - ans)
    
    for i,v := range nums {
        
        l, r := i + 1, n-1
        
        for l < r {
            
            sum := v + nums[l] + nums[r]
            
            if abs(sum - target) < diff {
                diff = abs(sum - target)
                ans = sum
            }
            
            if sum == target {
                return sum
            } else if sum < target {
                l++
            } else {
                r--
            }
            
        }
        
    }
    return ans
    
    
}

func abs(a int) int {
    if a < 0 { return -a }
    return a
}

/*

same as 3sum , only we need to find the closest value

*/
