package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/KaurKaranjot/bookslib/pkg/models"
	"github.com/gorilla/mux"
)

func (h handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	//Read the dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	//Find the book by id
	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}
	//Delet that book
	h.DB.Delete(&book)
	// for index, book := range mocks.Books {
	// 	if book.Id == id {

	// 		//Delete book and send response if the book id matches dynamically
	// 		mocks.Books = append(mocks.Books[:index], mocks.Books[index+1:]...)

	// 		break
	// 	}
	// }
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")
}
