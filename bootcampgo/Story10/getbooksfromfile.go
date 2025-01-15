package bootcamp

import (
	"bufio"
	"os"
)

type Book struct {
	Name   string
	Author string
	Year   int
}

func GetBooksFromCsv(path string) []*Book {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines [][]string

	for scanner.Scan() {
		line := scanner.Text()
		var record []string
		field := ""

		for _, char := range line {
			if char == ',' {
				record = append(record, field)
				field = ""
			} else {
				field += string(char)
			}
		}
		record = append(record, field)
		lines = append(lines, record)
	}

	if len(lines) < 2 {
		return nil
	}

	header := lines[0]
	var nameIndex, authorIndex, yearIndex int
	for i, h := range header {
		if h == "Name" {
			nameIndex = i
		} else if h == "Author" {
			authorIndex = i
		} else if h == "Year" {
			yearIndex = i
		}
	}

	var books []*Book

	for _, record := range lines[1:] {
		if len(record) != len(header) {
			continue
		}

		year := 0
		for _, char := range record[yearIndex] {
			year = year*10 + int(char-'0')
		}

		book := &Book{
			Name:   record[nameIndex],
			Author: record[authorIndex],
			Year:   year,
		}
		books = append(books, book)
	}

	return books
}

// func main() {
// 	books := GetBooksFromCsv("books.csv")
// 	for _, b := range books {
// 		fmt.Println(b.Name, b.Author, b.Year)
// 	}

// 	books2 := GetBooksFromCsv("books2.csv")
// 	for _, b := range books2 {
// 		fmt.Println(b.Name, b.Author, b.Year)
// 	}
// }
