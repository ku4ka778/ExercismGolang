package darts
import "math"
func Score(x, y float64) int {
	z := math.Sqrt(x*x + y*y)
    if z <=1{
        return 10
    } else if z <= 5 {
        return 5
    } else if z <= 10 {
        return 1
    } else {
        return 0
    }
}
