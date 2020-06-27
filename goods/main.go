package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var err error

// Book struct (Model)
type Good struct {
	ID          string `json:"id"`
	MerchantID  string `json:"merchant_id"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Cancellable string `json:"cancellable"`
	CheckAvail  string `json:"check_avail"`
	Currency    string `json:"currency"`
	PriceStart  string `json:"price_start"`
	PriceEnd    string `json:"price_end"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	CreatedBy   string `json:"created_by"`
	ModifiedAt  string `json:"modified_at"`
	ModifiedBy  string `json:"modified_by"`
}

// Get all orders

func getGoods(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var goods []Good

	sql := `SELECT
				id,
				IFNULL(merchant_id,''),
				IFNULL(code,'') Code,
				IFNULL(description,'') Description,
				IFNULL(category,'') Category,
				IFNULL(cancellable,'') Cancellable,
				IFNULL(check_avail,'') CheckAvail,
				IFNULL(currency,'') Currency ,
				IFNULL(price_start,'') PriceStart,
				IFNULL(price_end,'') PriceAnd,
				IFNULL(status,'') Status,
				IFNULL(created_at,'') CreatedAt,
				IFNULL(created_by,'') CreatedBy,
				IFNULL(modified_at,'') ModifiedAt,
				IFNULL(modified_by,'') ModifiedBy
			FROM goods`

	result, err := db.Query(sql)

	defer result.Close()

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		var good Good
		err := result.Scan(&good.ID, &good.MerchantID, &good.Code, &good.Description, &good.Category,
			&good.Cancellable, &good.CheckAvail, &good.Currency, &good.PriceStart, &good.PriceEnd,
			&good.Status, &good.CreatedAt, &good.CreatedBy, &good.ModifiedAt, &good.ModifiedBy)

		if err != nil {
			panic(err.Error())
		}
		goods = append(goods, good)
	}

	json.NewEncoder(w).Encode(goods)
}

func createGood(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		ID := r.FormValue("id")
		MerchantID := r.FormValue("merchantID")
		Code := r.FormValue("code")
		Description := r.FormValue("description")
		Category := r.FormValue("category")
		Cancellable := r.FormValue("cancellable")
		CheckAvail := r.FormValue("checkAvail")
		Currency := r.FormValue("currency")
		PriceStart := r.FormValue("priceStart")
		PriceEnd := r.FormValue("priceEnd")
		Status := r.FormValue("status")
		CreatedAt := r.FormValue("createdAt")
		CreatedBy := r.FormValue("createdBy")
		ModifiedAt := r.FormValue("modifiedAt")
		ModifiedBy := r.FormValue("modifiedBy")

		stmt, err := db.Prepare("INSERT INTO goods (id,merchant_id,code,description,category,cancellable,check_avail,currency,price_start,price_end,status,created_at,created_by,modified_at,modified_by) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
		_, err = stmt.Exec(ID, MerchantID, Code, Description, Category, Cancellable, CheckAvail, Currency,
			PriceStart, PriceEnd, Status, CreatedAt, CreatedBy, ModifiedAt, ModifiedBy)

		if err != nil {
			fmt.Fprintf(w, "Data Duplicate")
		} else {
			fmt.Fprintf(w, "Data Created")
		}
	}
}

func getGood(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var goods []Good
	params := mux.Vars(r)

	sql := `SELECT
				id,
				IFNULL(merchant_id,''),
				IFNULL(code,'') Code,
				IFNULL(description,'') Description,
				IFNULL(category,'') Category,
				IFNULL(cancellable,'') Cancellable,
				IFNULL(check_avail,'') CheckAvail,
				IFNULL(currency,'') Currency ,
				IFNULL(price_start,'') PriceStart,
				IFNULL(price_end,'') PriceAnd,
				IFNULL(status,'') Status,
				IFNULL(created_at,'') CreatedAt,
				IFNULL(created_by,'') CreatedBy,
				IFNULL(modified_at,'') ModifiedAt,
				IFNULL(modified_by,'') ModifiedBy
			FROM goods WHERE id = ?`

	result, err := db.Query(sql, params["id"])

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var good Good

	for result.Next() {

		err := result.Scan(&good.ID, &good.MerchantID, &good.Code, &good.Description, &good.Category,
			&good.Cancellable, &good.CheckAvail, &good.Currency, &good.PriceStart, &good.PriceEnd,
			&good.Status, &good.CreatedAt, &good.CreatedBy, &good.ModifiedAt, &good.ModifiedBy)

		if err != nil {
			panic(err.Error())
		}

		goods = append(goods, good)
	}

	json.NewEncoder(w).Encode(goods)
}

func updateGood(w http.ResponseWriter, r *http.Request) {

	if r.Method == "PUT" {

		params := mux.Vars(r)

		newMerchantID := r.FormValue("merchantID")
		newCode := r.FormValue("code")
		newCancellable := r.FormValue("cancellable")
		newCheckavail := r.FormValue("checkAvail")
		newPricestart := r.FormValue("priceStart")
		newPriceend := r.FormValue("priceEnd")

		stmt, err := db.Prepare("UPDATE goods SET merchant_id = ?,code = ?, cancellable=?, check_avail=?, price_start=?, price_end=? WHERE id = ?")
		_, err = stmt.Exec(newMerchantID, newCode, newCancellable, newCheckavail, newPricestart, newPriceend, params["id"])

		if err != nil {
			panic(err.Error())
		}

		fmt.Fprintf(w, "Goods with id = %s was updated", params["id"])
	}
}

func deleteGood(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	stmt, err := db.Prepare("DELETE FROM goods WHERE id = ?")

	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(params["id"])

	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "Good with ID = %s was deleted", params["id"])
}

func getPost(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var goods []Good

	ID := r.FormValue("id")
	MerchantID := r.FormValue("merchantID")

	sql := `SELECT
				id,
				IFNULL(merchant_id,''),
				IFNULL(code,'') Code,
				IFNULL(description,'') Description,
				IFNULL(category,'') Category,
				IFNULL(cancellable,'') Cancellable,
				IFNULL(check_avail,'') CheckAvail,
				IFNULL(currency,'') Currency ,
				IFNULL(price_start,'') PriceStart,
				IFNULL(price_end,'') PriceAnd,
				IFNULL(status,'') Status,
				IFNULL(created_at,'') CreatedAt,
				IFNULL(created_by,'') CreatedBy,
				IFNULL(modified_at,'') ModifiedAt,
				IFNULL(modified_by,'') ModifiedBy
			FROM goods WHERE id = ? AND merchant_id=?`

	result, err := db.Query(sql, ID, MerchantID)

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var good Good

	for result.Next() {

		err := result.Scan(&good.ID, &good.MerchantID, &good.Code, &good.Description, &good.Category,
			&good.Cancellable, &good.CheckAvail, &good.Currency, &good.PriceStart, &good.PriceEnd,
			&good.Status, &good.CreatedAt, &good.CreatedBy, &good.ModifiedAt, &good.ModifiedBy)

		if err != nil {
			panic(err.Error())
		}

		goods = append(goods, good)
	}

	json.NewEncoder(w).Encode(goods)

}

// Main function
func main() {

	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/db_testing")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	fmt.Println("Listen Server:8484")
	// Init router
	r := mux.NewRouter()

	// Route handles & endpoints
	r.HandleFunc("/goods", getGoods).Methods("GET")
	r.HandleFunc("/goods", createGood).Methods("POST")
	r.HandleFunc("/goods/{id}", getGood).Methods("GET")
	r.HandleFunc("/goods/{id}", updateGood).Methods("PUT")
	r.HandleFunc("/goods/{id}", deleteGood).Methods("DELETE")
	//New
	r.HandleFunc("/getgood", getPost).Methods("POST")

	// Start server
	log.Fatal(http.ListenAndServe(":8484", r))
}
