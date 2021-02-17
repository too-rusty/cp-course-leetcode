func integerBreak(N int) int {
    var ans int = 1
    for parts := 2; parts < N; parts++ {
        var val int = N / parts
        var left int = N % parts
        // N = parts * val + left
        var tmp int = 1
        for i := 0; i < parts; i++{
            if i < left {
                tmp *= (val + 1)
            } else {
                tmp *= val
            }
        }
        ans = max (ans, tmp)
    }
    return ans
}

// max(arr...)
// max(1,2,3,5,5)

func max(x ...int) int{
    mx := x[0]
    for _,v := range x {
        if v > mx {
            mx = v
        }
    }
    return mx
}


/*

A

a1 a2 a3 ... an

a1 + a2 + a3 .... + an = A
n >= 2
maximise the product

n < 60

x * ( N - x ) = P
---- Calculus Proof ------
dp/dx = 0
=> x = N / 2
3 => 1 1 1 

mod = 3 % 2 = 1

3 => 1 1+1

N = 8 , parts = 3

8%3 = 2
2 3 3


The proof is kind of intuitive or from calculus (as explained above)

the intuitive proof is that since we need to 
maximise the product we can try to make the numbers equal
like for eg , 4 => 2 + 2 => 2 * 2 = 4 but if we use 3, 1 the product is lesser ( 3 < 4 )
hence try to break into equal parts to maximise the product

*/


