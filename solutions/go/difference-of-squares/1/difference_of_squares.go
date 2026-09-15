package differenceofsquares

func SquareOfSum(n int) int {
    var z int
    for i:=1;i<=n;i++{
        z+=i
    }
    return z*z
}

func SumOfSquares(n int) int {
    var z int
    for i:=1;i<=n;i++{
        z+=(i*i)
    }
    return z
}

func Difference(n int) int {
	return (SumOfSquares(n) - SquareOfSum(n)) * -1
}
