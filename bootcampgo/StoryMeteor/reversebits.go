package bootcamp

func ReverseBits(n byte) byte {
	var reversed byte = 0
	for i := 0; i < 8; i++ {
		reversed <<= 1
		reversed |= n & 1
		n >>= 1
	}
	return reversed
}

// func main() {
// 	fmt.Printf("%08b %08b\n", 5, ReverseBits(5))     // 00000101 10100000
// 	fmt.Printf("%08b %08b\n", 255, ReverseBits(255)) // 11111111 11111111
// 	fmt.Printf("%08b %08b\n", 1, ReverseBits(1))     // 00000001 10000000
// }
