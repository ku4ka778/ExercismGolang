package isbnverifier

import "unicode"

func IsValidISBN(isbn string) bool {
	sum := 0
	weight := 10
	digitCount := 0

	for i := 0; i < len(isbn); i++ {
		r := rune(isbn[i])

		if r == '-' {
			continue
		}

		if r == 'X' || r == 'x' {
			if digitCount == 9 {
				sum += 10 * weight
				weight--
				digitCount++
				continue
			}
			return false
		}

		if !unicode.IsDigit(r) {
			return false
		}

		digit := int(r - '0')
		sum += digit * weight
		weight--
		digitCount++
	}

	if digitCount != 10 {
		return false
	}

	return sum%11 == 0
}
