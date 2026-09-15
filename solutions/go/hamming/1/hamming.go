package hamming
import "errors"
func Distance(a, b string) (int, error) {
	rune1 := []rune(a)
    rune2 := []rune(b)
    diffCount := 0
    if len(rune1) != len(rune2) {
		return 0, errors.New("Error")
	}
    for i := 0; i < len(rune1); i++ {
		if rune1[i] != rune2[i] {
			diffCount++
		}
	}
    return diffCount, nil
}
