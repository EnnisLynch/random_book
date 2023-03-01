package main

import (
	"bufio"
	"encoding/csv"
	"os"
	"strconv"
	"strings"
)

// Holds the details for a book
// note that loading bOok from new book
// will immediately load the buffer into
// memory. This code does not stream
// so keep the books small!
type Book struct {
	file_name        string
	book_text_buffer []byte
	chapters         []*Chapter
}

// This method will scan through all the chapters
// in the books and will return a string with the
// new book. Note that this method effectively
// can only be called once as it will scan
// through the chapters and not reset the buffer
// after it is finished
func (book *Book) BuildNewBook() string {
	var stringBuilder strings.Builder
	var keepGoing = true

	for keepGoing {
		var completedChapters = 0
		for _, element := range book.chapters {
			var tempWord = element.Scan()
			if tempWord == "" {
				//increment finished chapters
				completedChapters++
			} else {
				//if the sb is >0 then we are adding a word
				//thus we need a space before the word
				if stringBuilder.Len() > 0 {
					stringBuilder.WriteString(" ")
				}
				stringBuilder.WriteString(tempWord)

			}
		}
		if completedChapters == len(book.chapters) {
			//No more chapters left, exit the loop
			keepGoing = false
		}
	} //end foreach
	return stringBuilder.String()
}

// Will add a chapter to the the book's
// chapter collection
func (book *Book) AddChapter(chapter *Chapter) {
	book.chapters = append(book.chapters, chapter)
}

// Will return the book file name from the csv file name
// file_name.txt.csv will attempt to load
// file_name.txt as the book!
func GetFileNameFromCSVFileName(file_name string) string {
	var index = strings.Index(file_name, ".csv")
	if index == -1 {
		panic(".csv not found in file name")
	}
	return file_name[0:index]
}

// Will return a book and the chapters
// if the CSV is correct. Note format
// file_name.txt.csv will attempt to load
// file_name.txt as the book!
func NewBookFromCSV(file_name string) *Book {

	var fileStream, file_error = os.Open(file_name)
	var book_file_name = GetFileNameFromCSVFileName(file_name)
	checkForErrorAndFail(file_error)
	var result = NewBook(book_file_name)

	var lines, error_csv = csv.NewReader(fileStream).ReadAll()
	checkForErrorAndFail(error_csv)

	for index, record := range lines {
		if index == 0 {
			//skip the header
		} else {
			var start_index, start_error = strconv.Atoi(record[0])
			var end_index, end_error = strconv.Atoi(record[1])
			checkForErrorAndFail(start_error)
			checkForErrorAndFail(end_error)
			result.AddChapter(NewChapter(result, start_index, end_index))
		}
	}

	return result
}

// Creates a book and loads the buffer into memory
// Note the chapters are empty at this point and
// AddChapter needs to be called for each chapter (in order)
func NewBook(file_name string) *Book {
	book := new(Book)
	book.file_name = file_name
	fileStream, someError := os.ReadFile(book.file_name)
	checkForErrorAndFail(someError)
	book.book_text_buffer = []byte(fileStream)
	return book
}

// The Chapter of a book
type Chapter struct {
	start_byte int
	end_byte   int
	scanner    *bufio.Scanner
}

// Will return either the empty string or the
// next work in this chapter
func (chapter *Chapter) Scan() string {
	result := ""
	if chapter.scanner.Scan() {
		result = chapter.scanner.Text()
	}
	return result
}

// Creates a chapter of a book and will create a scanner for the
// chapter in memory
func NewChapter(book *Book, start_byte int, end_byte int) *Chapter {
	chapter := new(Chapter)
	chapter.start_byte = start_byte
	chapter.end_byte = end_byte

	//splice the book at the chapter start and end
	tempChapter := string(book.book_text_buffer[chapter.start_byte:chapter.end_byte])

	string_reader := strings.NewReader(tempChapter)
	chapter.scanner = bufio.NewScanner(string_reader)
	chapter.scanner.Split(bufio.ScanWords)

	return chapter
}

// Will panic on error
func checkForErrorAndFail(someError error) {
	if someError != nil {
		panic(someError)
	}
}
func writeToStringToFile(fileNameToCreate string, inputText string) {
	fileStream, error := os.Create(fileNameToCreate)
	checkForErrorAndFail(error)

	fileStream.WriteString(inputText)
	fileStream.Sync()
	fileStream.Close()
}
func main() {
	book := NewBookFromCSV("pelle.txt.csv")
	var result = book.BuildNewBook()
	writeToStringToFile("output.txt", result)
}
