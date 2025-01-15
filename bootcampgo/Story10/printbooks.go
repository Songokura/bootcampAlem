package bootcamp

import "github.com/alem-platform/ap"

// type Book struct {
// 	Name   string
// 	Author string
// 	Year   int
// }

func PutString(s string) {
	for _, v := range s {
		ap.PutRune(v)
	}
}

func PrintBooks(books []*Book) {
	lgh, lgh1 := 0, 0
	for _, v := range books {
		if len(v.Name) > lgh {
			lgh = len(v.Name)
		}
		if len(v.Author) > lgh1 {
			lgh1 = len(v.Author)
		}
	}

	if len("Name") < lgh {
		PutString("Name")
		for i := 0; i < lgh-len("Name")+1; i++ {
			ap.PutRune(' ')
		}
	} else {
		PutString("Name")
		ap.PutRune(' ')
	}

	if len("Author") < lgh1 {
		PutString("Author")
		for i := 0; i < lgh1-len("Author")+1; i++ {
			ap.PutRune(' ')
		}
	} else {
		PutString("Author")
		ap.PutRune(' ')
	}
	PutString("Year")
	ap.PutRune('\n')

	for _, v := range books {
		if len(v.Name) < lgh {
			PutString(v.Name)
			for i := 0; i < lgh-len(v.Name)+1; i++ {
				ap.PutRune(' ')
			}
		} else {
			PutString(v.Name)
			ap.PutRune(' ')
		}

		if len(v.Author) < lgh1 {
			PutString(v.Author)
			for i := 0; i < lgh1-len(v.Author)+1; i++ {
				ap.PutRune(' ')
			}
		} else {
			PutString(v.Author)
			ap.PutRune(' ')
		}

		PutString(itoa(v.Year))
		ap.PutRune('\n')
	}
}

// Helper function to convert an integer to string
func itoa(n int) string {
	var buf [20]byte
	i := len(buf)
	neg := n < 0
	if neg {
		n = -n
	}
	for n >= 10 {
		i--
		buf[i] = byte(n%10) + '0'
		n /= 10
	}
	i--
	buf[i] = byte(n) + '0'
	if neg {
		i--
		buf[i] = '-'
	}
	return string(buf[i:])
}

// Helper function to print a newline with ap.PutRune
func putNewline() {
	ap.PutRune('\n')
}

// func main() {
// 	books := []*Book{
// 		{Name: "The Kaizen Way", Author: "Robert Maurer", Year: 2009},
// 		{Name: "Dialogs", Author: "Plato", Year: -400},
// 		{Name: "Unknown", Author: "Unknown", Year: 10},
// 	}

// 	PrintBooks(books)
// }
