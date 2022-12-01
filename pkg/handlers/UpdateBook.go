package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/KaurKaranjot/bookslib/pkg/models"
	"github.com/gorilla/mux"
)

func (h handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	//Read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	// //Read request body

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var updatedBook models.Book
	json.Unmarshal(body, &updatedBook)

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}
	// //Iterate over all the mock books
	// for index, book := range mocks.Books {
	// 	if book.Id == id {
	// 		//Update and send response when book id matches dynamic id
	book.Title = updatedBook.Title
	book.Author = updatedBook.Author
	book.Desc = updatedBook.Desc

	h.DB.Save(&book)

	// 		mocks.Books[index] = book

	// 		break
	// 	}
	// }
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Updated")
}
