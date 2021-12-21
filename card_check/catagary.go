package card_check

func Category(bin string) string {
	if bin[:2] == "34" || bin[:2] == "37" {
		return "American Express"
	} else if bin[:5] == "5610" {
		return "Bankcard"
	} else if bin[:4] == "500" {
		return "BMO ABM Card"
	} else if bin[:5] == "4506" {
		return "Canadian Imperial Bank of Commerce Advantage Debit Card"
	} else if bin[:2] == "31" {
		return "CHina T-union"
	} else if bin[:2] == "50" || bin[:2] == "56" {
		return "Maestro"
	} else if bin[:2] == "60" || bin[:5] == "6521" || bin[:5] == "6522" {
		return "Rupay"
	} else if (bin[:5] >= "2221" && bin[0:5] <= "22004") || (bin[:2] >= "51" && bin[0:2] <= "55") {
		return "Mastercard"
	} else if bin[:1] == "4" {
		return "Visa"
	} else {
		return "card name not found"
	}
}
