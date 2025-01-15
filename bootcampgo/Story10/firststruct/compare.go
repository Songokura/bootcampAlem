package firststruct

// type User struct {
// 	Name     string
// 	password string
// }

// func NewUser(name, password string) User {
// 	return User{
// 		Name:     name,
// 		password: password,
// 	}
// }

func (u User) Compare(v User) bool {
	return u.Name == v.Name && u.password == v.password
}

// func main() {
// 	user1 := NewUser("Alice", "password123")
// 	user2 := NewUser("Alice", "password123")
// 	fmt.Println("Users are identical:", user1.Compare(user2)) // Users are identical: true

// 	user3 := NewUser("Bob", "securePassword456")
// 	user4 := NewUser("Charlie", "qwerty789")
// 	fmt.Println("Users are identical:", user3.Compare(user4)) // Users are identical: false
// }
