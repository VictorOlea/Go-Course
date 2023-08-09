package main

import "fmt"

type Book struct {
	bTitle      string
	bAuthorName string
	bSubject    string
	book_id     int
}

func (b Book) lecture() {
	fmt.Println(b.bSubject)
}

func main() {
	var Book1 Book
	//var Book2 Book

	Book1.bTitle = "The Go Programming Language"
	Book1.bAuthorName = "Alan A. A Donovan and Brian W. Kernighan"
	Book1.bSubject = "A complete guide to Go programming"
	Book1.book_id = 6495

	printBookDetails(&Book1)

	fmt.Printf("Authors : %s\n", Book1.bAuthorName)

	Book1.lecture()

}

func printBookDetails(book *Book) {
	book.bAuthorName = "Unknow"
	fmt.Printf("\nTitle : %s\n", book.bTitle)
	fmt.Printf("Authors : %s\n", book.bAuthorName)
	fmt.Printf("Subject : %s\n", book.bSubject)
	fmt.Printf("Book ID : %d\n", book.book_id)
}
