package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "king", "queen", "jack":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	v1 := ParseCard(card1)
	v2 := ParseCard(card2)
	vDealer := ParseCard(dealerCard)
	handSum := v1 + v2
	if card1 == "ace" && card2 == "ace" {
		return "P"
	}
	if handSum == 21 {
		if vDealer >= 10 {
			return "S"
		}
		return "W"
	}
	if handSum >= 17 {
		return "S"
	}
	if handSum >= 12 && handSum <= 16 {
		if vDealer < 7 {
			return "S"
		}
		return "H"
	}
	return "H"
}