package bootcamp

func MorseToText(s string) string {
	morseMap := map[string]string{
		".-": "A", "-...": "B", "-.-.": "C", "-..": "D",
		".": "E", "..-.": "F", "--.": "G", "....": "H",
		"..": "I", ".---": "J", "-.-": "K", ".-..": "L",
		"--": "M", "-.": "N", "---": "O", ".--.": "P",
		"--.-": "Q", ".-.": "R", "...": "S", "-": "T",
		"..-": "U", "...-": "V", ".--": "W", "-..-": "X",
		"-.--": "Y", "--..": "Z", ".----": "1", "..---": "2",
		"...--": "3", "....-": "4", ".....": "5", "-....": "6",
		"--...": "7", "---..": "8", "----.": "9", "-----": "0",
		"--..--": ",", ".-.-.-": ".", "..--..": "?",
	}
	var text string
	currWord := ""

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if val, ok := morseMap[currWord]; ok {
				text += val
			} else {
				text += currWord // append as-is if not in map (e.g., for spaces)
			}
			currWord = ""
			if i+1 < len(s) && s[i+1] == ' ' {
				text += " " // add space between words
				i++         // skip the next space character
			}
		} else {
			currWord += string(s[i])
		}
	}

	// Append the last word
	if val, ok := morseMap[currWord]; ok {
		text += val
	} else {
		text += currWord // append as-is if not in map
	}

	return text
}

// func main() {
// 	fmt.Println(MorseToText("... --- ..."))                                                       // SOS
// 	fmt.Println(MorseToText("... .- .-.. . -- --..-- .... --- .-- .- .-. . -.-- --- ..- ..--..")) // SALEM,HOWAREYOU?
// }
