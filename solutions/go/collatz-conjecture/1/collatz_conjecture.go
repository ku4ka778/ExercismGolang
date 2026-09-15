package collatzconjecture
import "errors"
func CollatzConjecture(n int) (int, error) {
    if n <= 0 {
		return 0, errors.New("number must be greater than 0")
	}
    var z int
	for n>1{
        if n%2==0{
            n = n/2
        	z++
        } else{
            n = n*3+1
        	z++
        }
    }
    return z,nil
}

