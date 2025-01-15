package firststruct

// import "fmt"

// type User struct {
// 	Name     string
// 	password string
// }

// // NewUser function to create a new User instance
// func NewUser(name, password string) User {
// 	return User{
// 		Name:     name,
// 		password: password,
// 	}
// }

func (u User) PasswordReliability() string {
	if u.password == "" {
		return "undefined"
	}

	criteriaMet := 0

	if len(u.password) >= 8 {
		criteriaMet++
	}

	hasUpper := false
	for _, v := range u.password {
		if 'A' <= v && v <= 'Z' {
			hasUpper = true
			break
		}
	}
	if hasUpper {
		criteriaMet++
	}

	hasLower := false
	for _, v := range u.password {
		if 'a' <= v && v <= 'z' {
			hasLower = true
			break
		}
	}

	if hasLower {
		criteriaMet++
	}

	hasDigit := false
	for _, v := range u.password {
		if '0' <= v && v <= '9' {
			hasDigit = true
			break
		}
	}

	if hasDigit {
		criteriaMet++
	}

	hasSpecial := false
	specialChars := "~!@#$%^&*()-_=+[{]}\\|;:'\",<.>/?"

	for _, v := range u.password {
		if !isLetterOrDigit(v) && containsRune(specialChars, v) {
			hasSpecial = true
			break
		}
	}

	if hasSpecial {
		criteriaMet++
	}

	switch criteriaMet {
	case 5:
		return "strong"
	case 3, 4:
		return "medium"
	case 1, 2:
		return "easy"
	default:
		return "undefined"
	}
}

func isLetterOrDigit(r rune) bool {
	return ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z') || ('0' <= r && r <= '9')
}

func containsRune(s string, r rune) bool {
	for _, c := range s {
		if c == r {
			return true
		}
	}
	return false
}

// func main() {
// 	user := NewUser("Alice", "StrongPassword123$")

// 	reliability := user.PasswordReliability()
// 	fmt.Println("Password reliability:", reliability)
// 	// Password reliability: strong
// }
