package main

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func rotate(s []int, posi int) {
	reverse(s[:posi])
	reverse(s[posi:])
	reverse(s)
}
