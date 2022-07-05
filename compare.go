package codewars

func close_compare(a, b, margin int) bool {
	return a >= b-margin && a <= b+margin
}
