package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/KaurKaranjot/bookslib/pkg/models"
)

func (h handler) AddBook(w http.ResponseWriter, r *http.Request) {
	//Read to request
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var book models.Book
	json.Unmarshal(body, &book)

	// Append to the book mocks
	if result := h.DB.Create(&book); result.Error != nil {
		fmt.Println(result.Error)
	}
	// book.Id=rand.Intn(100)
	// mocks.Books = append(mocks.Books, book)
	//Send a 201 craeted response
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Created")
}
