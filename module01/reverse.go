package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
// func Reverse(word string) string {
// 	var revd string
// 	for i := len(word) - 1; i >= 0; i-- {
// 		revd += string(word[i])
// 	}
// 	return revd
// }

// A more efficient approach using 'strings.Builder' to concatenate strings to the previous (above) solutions
// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
// func Reverse(word string) string {
// 	var revd strings.Builder
// 	for i := len(word) - 1; i >= 0; i-- {
// 		revd.WriteByte(word[i])
// 	}
// 	return revd.String()
// }

//
// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
// func Reverse(word string) string {
// 	var revd string
// 	for i := 0; i < len(word); i++ {
// 		revd = string(word[i]) + revd
// 	}
// 	return revd
// }

// That approach supports Rune what it is the best alternative
// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {
	var revd string
	for _, v := range word {
		revd = string(v) + revd
	}
	return revd
}
