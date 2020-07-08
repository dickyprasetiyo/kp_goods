package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var err error

type SuperPower struct {
	Age    int64 `json:"age"`
	Errors struct {
		Detail string `json:"detail"`
		Source struct {
			Pointer string `json:"pointer"`
		} `json:"source"`
		Status string `json:"status"`
		Title  string `json:"title"`
	} `json:"errors"`
	Name   string `json:"name"`
	Powers struct {
		Task1 string `json:"Task1"`
		Task2 string `json:"Task2"`
		Task3 string `json:"Task3"`
	} `json:"powers"`
	SecretIdentity string `json:"secretIdentity"`
}

type Food struct {
	ID    string `json:"id"`
	Image struct {
		Height int64  `json:"height"`
		URL    string `json:"url"`
		Width  int64  `json:"width"`
	} `json:"image"`
	Name      string `json:"name"`
	Thumbnail struct {
		Height int64  `json:"height"`
		URL    string `json:"url"`
		Width  int64  `json:"width"`
	} `json:"thumbnail"`
	Type string `json:"type"`
}

func getSuperPower(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)

	var request1 SuperPower

	if err = json.Unmarshal(body, &request1); err != nil {
		fmt.Println("Failed decoding json message")
	}

	Age := request1.Age
	Name := request1.Name
	SecretId := request1.SecretIdentity
	Task1 := request1.Powers.Task1
	Task2 := request1.Powers.Task2
	Task3 := request1.Powers.Task3
	Status := request1.Errors.Status
	Pointer := request1.Errors.Source.Pointer
	Title := request1.Errors.Title
	Detail := request1.Errors.Detail

	stmt, err := db.Prepare("INSERT INTO superpower (Name,Age,SecretIdentity,Task1,Task2,Task3,Status,Pointer,Title,Detail) VALUES(?,?,?,?,?,?,?,?,?,?)")
	_, err = stmt.Exec(Name, Age, SecretId, Task1, Task2, Task3, Status, Pointer, Title, Detail)

	if err != nil {
		fmt.Fprintf(w, "Data Duplicate")
	} else {
		fmt.Fprintf(w, "Data Created")
	}
}

func getFood(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)

	var request2 Food

	if err = json.Unmarshal(body, &request2); err != nil {
		fmt.Println("Failed decoding json message")
	}

	ID := request2.ID
	Type := request2.Type
	Name := request2.Name
	ImageURL := request2.Image.URL
	ImageWidth := request2.Image.Width
	ImageHeight := request2.Image.Height
	ThumbnailURL := request2.Thumbnail.URL
	ThumbnailWidth := request2.Thumbnail.Width
	ThumbnailHeight := request2.Image.Height

	stmt, err := db.Prepare("INSERT INTO food (IDFood,Type,Name,URL_image,Width_image,Height_image,URL_thumbnail,Width_thumbnail,Height_thumbnail) VALUES(?,?,?,?,?,?,?,?,?)")
	_, err = stmt.Exec(ID, Type, Name, ImageURL, ImageWidth, ImageHeight, ThumbnailURL, ThumbnailWidth, ThumbnailHeight)

	if err != nil {
		fmt.Fprintf(w, "Data Duplicate")
	} else {
		fmt.Fprintf(w, "Data Created")
	}
}

func main() {

	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/task_4")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// Init router
	r := mux.NewRouter()

	fmt.Println("Server on :8181")

	// Route handles & endpoints
	r.HandleFunc("/superPower", getSuperPower).Methods("POST")
	r.HandleFunc("/food", getFood).Methods("POST")

	// Start server
	log.Fatal(http.ListenAndServe(":8181", r))

}
