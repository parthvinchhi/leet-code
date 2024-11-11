package easy

func FindtheIndex(haystack, needle string) int {
	ln := len(needle)
	lh := len(haystack)

	if lh < ln {
		return -1
	}

	// for i := range haystack {
	// 	if i < ln || i <= ln {
	// 		if haystack[i:i+ln] == needle {
	// 			return i
	// 		}
	// 	}
	// }

	for i := range haystack {
		if i < ln && haystack[i:i+ln] == needle {
			return i
		} else if i == ln && haystack[i:] == needle {
			return i
		}
	}

	return -1
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
