package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {	
	switch card {
	case "ace": return 11
	case "two": return 2
	case "three": return 3
	case "four": return 4
	case "five": return 5
	case "six": return 6
	case "seven": return 7
	case "eight": return 8
	case "nine": return 9
	case "ten": return 10
	case "jack": return 10
	case "queen": return 10
	case "king": return 10
	default: return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	cardSum := ParseCard(card1) + ParseCard(card2)
	switch {
	// TOFIX: 
	case (cardSum) <= 11: return "H"
	// If you have a pair of aces you must always split them.
	case card1 == "ace" && card2 == "ace": return "P"
		/* If you have a Blackjack (two cards that sum up to
		   a value of 21), and the dealer does not have an
		   ace, a face card (Jack/Queen/King) or a ten then
		   you automatically win.
		*/
	case (cardSum == 21) && (dealerCard != "ace" &&
		dealerCard != "jack" &&
		dealerCard != "queen" &&
		dealerCard != "king" &&
		dealerCard != "ten"): return "W"
		/*  If the dealer does have any of those cards then
		    you'll have to stand and wait for the reveal of
		    the other card.
		*/ 
	case (cardSum == 21) && (dealerCard == "ace" ||
		dealerCard == "jack" ||
		dealerCard == "queen" ||
		dealerCard == "king" ||
		dealerCard == "ten"): return "S"
		/* If your cards sum up to a value within the range
		   [17, 20] you should always stand.
		*/
	case (cardSum >= 17) || (cardSum <= 20) && ParseCard(dealerCard) < 7: return "S"
		/* If your cards sum up to a value within the range
		   [12, 16] you should always stand unless the dealer
		   has a 7 or higher, in which case you should always
		   hit.
		*/
	case (cardSum >= 12) || (cardSum <= 16) && ParseCard(dealerCard) >= 7: return "H"
	/* If your cards sum up to a value within the range
		   [12, 16] you should always stand
	*/
	case (cardSum >= 12) || (cardSum <= 16) && ParseCard(dealerCard) < 7: return "S"
	/* If your cards sum up to 11 or lower you should
		   always hit.
	*/
	default: return "H"
	}
}
