package bootcamp

import (
	"crypto/md5"
)

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// Helper function to convert byte slice to hexadecimal string
func bytesToHex(data [16]byte) string {
	hexChars := "0123456789abcdef"
	result := make([]byte, 32)
	for i, b := range data {
		result[i*2] = hexChars[b>>4]
		result[i*2+1] = hexChars[b&0x0f]
	}
	return string(result)
}

// Helper function to compute the MD5 hash of a string and return its hexadecimal representation
func computeMD5Hex(s string) string {
	hash := md5.Sum([]byte(s))
	return bytesToHex(hash)
}

// Recursive function to generate all combinations of the given charset
func generate(prefix string, length int, targetHash string) (string, bool) {
	if length == 0 {
		if computeMD5Hex(prefix) == targetHash {
			return prefix, true
		}
		return "", false
	}
	for _, char := range charset {
		result, found := generate(prefix+string(char), length-1, targetHash)
		if found {
			return result, true
		}
	}
	return "", false
}

// BruteForceHash function definition
func BruteForceHash(hash string) string {
	if len(hash) != 32 {
		return ""
	}

	// Try all lengths from 1 to 6
	for length := 1; length <= 6; length++ {
		result, found := generate("", length, hash)
		if found {
			return result
		}
	}

	return ""
}

// // Main function for testing
// func main() {
// 	fmt.Println(BruteForceHash("ab6ccd17455d5347c49606d641e0b2af")) // SALEM
// 	fmt.Println(BruteForceHash("3cbfa33db66b830bfcf47ecc956505f8")) // ALEM
// 	fmt.Println(BruteForceHash(""))                                 //
// 	fmt.Println(BruteForceHash("abc"))                              //
// }
