package bootcamp

func TextToMorse(s string) string {
	morseCode := ""

	morseMap := map[rune]string{
		'A': ".-", 'B': "-...", 'C': "-.-.", 'D': "-..", 'E': ".",
		'F': "..-.", 'G': "--.", 'H': "....", 'I': "..", 'J': ".---",
		'K': "-.-", 'L': ".-..", 'M': "--", 'N': "-.", 'O': "---",
		'P': ".--.", 'Q': "--.-", 'R': ".-.", 'S': "...", 'T': "-",
		'U': "..-", 'V': "...-", 'W': ".--", 'X': "-..-", 'Y': "-.--",
		'Z': "--..", '1': ".----", '2': "..---", '3': "...--",
		'4': "....-",
		'5': ".....",
		'6': "-....",
		'7': "--...",
		'8': "---..",
		'9': "----.",
		'0': "-----",
		'.': ".-.-.-",
		',': "--..--",
		'?': "..--..",
	}
	for _, char := range s {
		upperChar := char
		if char >= 'a' && char <= 'z' {
			upperChar -= 32
		}

		if morse, ok := morseMap[upperChar]; ok {
			morseCode += morse + " "
		}
	}
	morseCode = morseCode[:len(morseCode)-1]

	return morseCode
}

// func main() {
// 	fmt.Println(TextToMorse("SOS"))                 // ... --- ...
// 	fmt.Println(TextToMorse("salem, how are you?")) // ... .- .-.. . -- --..-- .... --- .-- .- .-. . -.-- --- ..- ..--..
// }
