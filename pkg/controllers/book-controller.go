package controllers

import (
	"fmt"
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/ctycho/go_bookstore/pkg/models"
	"github.com/ctycho/go_bookstore/pkg/utils"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	allBooks := models.GetBooks()
	res, _ := json.Marshal(allBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}