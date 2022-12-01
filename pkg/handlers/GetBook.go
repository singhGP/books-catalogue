package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/KaurKaranjot/bookslib/pkg/models"
	"github.com/gorilla/mux"
)

func (h handler) GetBook(w http.ResponseWriter, r *http.Request) {
	//Read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	//find book by id
	var book models.Book

	if result:=h.DB.First(&book,id);result.Error !=nil{
fmt.Println(result.Error)
	}
	// for _, book := range mocks.Books{
	// 	if book.Id ==id{

	//IF ids are equal send book as a response

	// 	break
	// }
	// 	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}
