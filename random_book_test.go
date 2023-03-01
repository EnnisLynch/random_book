package main

import (
	"testing"
)

func Test_Data(t *testing.T) {
	book := NewBook("data.txt")
	book.AddChapter(NewChapter(book, 0, 45))
	book.AddChapter(NewChapter(book, 46, 91))
	book.AddChapter(NewChapter(book, 92, 137))
	book.AddChapter(NewChapter(book, 138, 183))
	book.AddChapter(NewChapter(book, 184, 229))
	book.AddChapter(NewChapter(book, 230, 275))
	book.AddChapter(NewChapter(book, 276, 321))
	book.AddChapter(NewChapter(book, 322, 367))
	book.AddChapter(NewChapter(book, 368, 413))
	book.AddChapter(NewChapter(book, 414, 458))

	var result = book.BuildNewBook()
	if result != "The The The The The The The The The The quick quick quick quick quick quick quick quick quick quick brown brown brown brown brown brown brown brown brown brown fox fox fox fox fox fox fox fox fox fox jumps jumps jumps jumps jumps jumps jumps jumps jumps jumps over over over over over over over over over over the the the the the the the the the the lazy lazy lazy lazy lazy lazy lazy lazy lazy lazy dog. dog. dog. dog. dog. dog. dog. dog. dog. dog." {
		t.Error("Output doesn't match!", 0)
	}
}

func Test_GetFileNameFromCSVFileName(t *testing.T) {
	var result = GetFileNameFromCSVFileName("pelle.txt.csv")
	if result != "pelle.txt" {
		t.Error("File Name is not correct", result)
	}
}

func Test_CsvFileWorksWithBooks(t *testing.T) {
	book := NewBook("pelle.txt")
	//Chapter 1
	book.AddChapter(NewChapter(book, 4912, 48186))
	//Chapter 1 part II
	book.AddChapter(NewChapter(book, 48203, 61471))
	//Chapter 1 part III
	book.AddChapter(NewChapter(book, 61489, 114018))
	//Chapter 1 part IV
	book.AddChapter(NewChapter(book, 114035, 147708))
	//Chapter 1 part V
	book.AddChapter(NewChapter(book, 147724, 166705))
	//Chapter 1 part VI
	book.AddChapter(NewChapter(book, 166722, 181367))
	//Chapter 1 part VII
	book.AddChapter(NewChapter(book, 181385, 234047))
	//Chapter 1 part VIII
	book.AddChapter(NewChapter(book, 234067, 251007))
	//Chapter 1 part IX
	book.AddChapter(NewChapter(book, 251024, 268045))
	//Chapter 1 part X
	book.AddChapter(NewChapter(book, 268061, 285964))
	//Chapter 1 part XI
	book.AddChapter(NewChapter(book, 285981, 321193))
	//Chapter 1 part XII
	book.AddChapter(NewChapter(book, 321211, 347309))
	var result = book.BuildNewBook()

	csvBook := NewBookFromCSV("pelle.txt.csv")
	csvResult := csvBook.BuildNewBook()

	if result != csvResult {
		t.Error("The books are not the same")
	}

}
