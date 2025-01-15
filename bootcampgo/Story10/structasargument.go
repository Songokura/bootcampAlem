package bootcamp

import (
	"bootcamp/firststruct"

	"github.com/alem-platform/ap"
)

func PrintF(s string) {
	for _, v := range s {
		ap.PutRune(v)
	}
	ap.PutRune('\n')
}

func PrintUserInfo(u firststruct.User) {
	is := "no"
	PrintF("Name: " + u.Name)
	passw := u.GetPassword()
	if passw != "" {
		is = "yes"
	}
	PrintF("HasPassword: " + is)
}

// func main() {
// 	userAlem := firststruct.NewUser("Alem", "hello.alem")
// 	PrintUserInfo(userAlem)
// 	// Output:
// 	// Name: Alem
// 	// HasPassword: yes

// 	userDias := firststruct.NewUser("Dias", "")
// 	PrintUserInfo(userDias)
// 	// Output:
// 	// Name: Dias
// 	// HasPassword: no
// }
