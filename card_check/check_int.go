package card_check

import "strconv"

func SixTeen_dig(card_num string) int8 {
	var last_val int

	for i := 0; i < len(card_num)-1; i += 2 {

		one_v, _ := strconv.Atoi(string(card_num[i]))
		one_v *= 2

		if one_v > 9 {
			one_v = (one_v % 10) + (one_v / 10)
		}
		last_val += one_v
	}

	return 10 - int8(last_val%10)
}
