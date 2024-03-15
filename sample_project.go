//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

type book struct {
	name  string
	count int
}
type member struct {
	name string
}
type library struct {
	books     []book
	members   []member
	histories []history
}
type history struct {
	timeOut   time.Time
	timeIn    time.Time
	user      string
	takenBook *book
}

func outBook(myLibrary *library, bookName string, userName string) {
	for j := range (myLibrary).members {
		if userName == (myLibrary).members[j].name {
			for i := range (*myLibrary).books {
				if bookName == myLibrary.books[i].name && myLibrary.books[i].count != 0 {
					myLibrary.books[i].count--
					myLibrary.histories = append(myLibrary.histories, history{
						timeOut:   time.Now(),
						user:      userName,
						takenBook: &myLibrary.books[i],
					})
					fmt.Println(userName, "borrowed", bookName, "book")
					return
				}
			}
		}
	}
	fmt.Println("sry this book or user not found in this library")
}

func returnBook(myLibrary *library, bookName string, userName string) {
	var t time.Time
	for i := range (myLibrary).histories {
		if bookName == (*myLibrary).histories[i].takenBook.name && userName == (*myLibrary).histories[i].user &&
			(*myLibrary).histories[i].timeOut != t && (*myLibrary).histories[i].timeIn == t {
			myLibrary.histories[i].takenBook.count++
			myLibrary.histories[i].timeIn = time.Now()
			fmt.Println(userName, "returned", bookName, "book")
			return
		}
	}
	fmt.Println("sry this book  not found in this library")
}

func countBook(myLibrary *library) {
	for i := range (*myLibrary).books {
		fmt.Println("count of", (*myLibrary).books[i].name, "are", (*myLibrary).books[i].count)
	}
}

func borrowedBooks(myLibrary *library, userName string) {
	var borrowedAny bool
	var t1 time.Time
	for i := range (*myLibrary).histories {
		if (*myLibrary).histories[i].user == userName && (*myLibrary).histories[i].timeIn == t1 {
			borrowedAny = true
			fmt.Println(userName, "borrowed", (*myLibrary).histories[i].takenBook.name, "at", (*myLibrary).histories[i].timeOut)
		}
	}
	if !borrowedAny {
		fmt.Println(userName, "did not borrow any book")
	}
}

func addBook(myLibrary *library, bookName string) {
	for i := range (*myLibrary).books {
		if bookName == (*myLibrary).books[i].name {
			(*myLibrary).books[i].count++
			fmt.Println("this book already exist")
			return
		}
	}
	(*myLibrary).books = append((*myLibrary).books, book{bookName, 1})
}

func addMember(myLibrary *library, userName string) {
	for i := range (*myLibrary).members {
		if userName == (*myLibrary).members[i].name {
			fmt.Println("this member exist")
			return
		}
	}
	(*myLibrary).members = append((*myLibrary).members, member{userName})
}

func bookHistory(myLibrary *library, bookName string) {
	var t time.Time
	for j := range (*myLibrary).books {
		if bookName == (*myLibrary).books[j].name {
			for i := range (*myLibrary).histories {
				if (*myLibrary).histories[i].timeIn == t {
					fmt.Println("the history of", myLibrary.histories[i].takenBook.name, "is:", "borrowed by",
						(*myLibrary).histories[i].user, "at", (*myLibrary).histories[i].timeOut)
				} else {
					fmt.Println("the history of", myLibrary.histories[i].takenBook.name, "is:", "borrowed by",
						(*myLibrary).histories[i].user, "at", (*myLibrary).histories[i].timeOut, "returned at", (*myLibrary).histories[i].timeIn)
				}

			}
		}
	}
}

func main() {
	yasLib := library{books: []book{{"farsi", 10}, {"arabi", 4}}, members: []member{{"nasrin"}, {"behrad"}}}
	countBook(&yasLib)
	outBook(&yasLib, "farsi", "behrad")
	outBook(&yasLib, "farsi", "behrad")
	outBook(&yasLib, "farsi", "behrad")
	countBook(&yasLib)
	returnBook(&yasLib, "farsi", "behrad")
	countBook(&yasLib)
	bookHistory(&yasLib, "farsi")
	borrowedBooks(&yasLib, "behrad")
	addMember(&yasLib, "Behrad")
	addBook(&yasLib, "farsI")
	fmt.Println(yasLib)
}
