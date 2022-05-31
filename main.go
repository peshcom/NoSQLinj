package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func register(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
		return
	}
	var req user
	err = json.Unmarshal(reqBody, &req)
	if err != nil {
		fmt.Fprintf(w, "Unmarshal error")
		return
	}

	result, err := userCollection.InsertOne(ctx, req)
	fmt.Println(result)
	if err != nil {
		fmt.Fprintf(w, "Error insert row")
		return
	}

	err = json.NewEncoder(w).Encode(req)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func auth(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
		return
	}

	var req user
	err = json.Unmarshal(reqBody, &req)
	if err != nil {
		fmt.Fprintf(w, "Unmarshal error")
		return
	}

	// auth
	success, err := checkUser(req.Username, req.Password)
	if (!success) || (err != nil) {
		fmt.Fprintf(w, "Wrong username or password")
		return
	}

	fmt.Fprintf(w, "AUTH SUCCESSFUL")
}

func main() {
	fmt.Println("Start web server")
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/auth", auth).Methods("POST")
	router.HandleFunc("/register", register).Methods("POST")

	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatal(err)
	}
}
