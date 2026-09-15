package luhn
import "unicode"
func Valid(id string) bool {
	sum := 0
	isSecond := false
	digitCount := 0
    
	for i:=len(id) - 1; i >=0; i-- {
		r:= rune(id[i])

		if unicode.IsSpace(r){
			continue
		}

		if !unicode.IsDigit(r){
			return false
		}

		digit := int(r - '0')
		digitCount++
        
		if isSecond {
			digit*=2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		isSecond = !isSecond
	}
    if digitCount <= 1 {
		return false
	}
	return sum%10 == 0
}