package easy

func FindtheIndex(haystack, needle string) int {
	mainLen := len(haystack)
	subLen := len(needle)

	// Loop through the main string to find the start of the substring
	for i := 0; i <= mainLen-subLen; i++ {
		match := true
		for j := 0; j < subLen; j++ {
			if haystack[i+j] != needle[j] {
				match = false
				break
			}
		}
		if match {
			return i // Return the index if we found the substring
		}
	}
	return -1 // Return -1 if the substring is not found
}

// func strStr(haystack string, needle string) int {
//     l := len(needle)
// 	for i := range haystack {
// 		if haystack[i:i+l] == needle {
// 			return i
// 		}
// 	}
// 	return -1
// }
