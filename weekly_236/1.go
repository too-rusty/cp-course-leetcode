func arraySign(nums []int) (ans int) {
    cnt := 0
    for _, v := range nums {
        if v < 0 {
            cnt++
        } else if v == 0 {
            return
        }
    }
    if cnt % 2 == 0 {
        ans = 1
    } else {
        ans = -1
    }
    return
}

/*
if atleast 1 zero then 0
if num of -ve numbers is even then positive else -ve
*/