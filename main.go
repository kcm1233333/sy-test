package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

const (
	host     = "peanut.db.elephantsql.com"
	port     = 5432
	user     = "zkhnlyuj"
	password = "0uAnAFE-OGiQrTFJdhxPjdbz38Sp2xyb"
	dbname   = "zkhnlyuj"
)

type Pengguna struct {
	kodepengguna   string
	namapengguna   string
	alamatpengguna string
	emailpengguna  string
	katasandi      string
	kodeotp        string
}
type Body struct {
	body *gorm.DB
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully Inserted!")
	}
}

type Response struct {
	status string `json:"Status"`
}

func Registration(w http.ResponseWriter, r *http.Request) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	keyVal := make(map[string]string)
	json.Unmarshal(b, &keyVal)
	kodepengguna := keyVal["kodepengguna"]
	namapengguna := keyVal["namapengguna"]
	alamatpengguna := keyVal["alamatpengguna"]
	emailpengguna := keyVal["emailpengguna"]
	katasandi := keyVal["katasandi"]
	kodeotp := keyVal["kodeotp"]

	insertStmt := `insert into "pengguna" values($1, $2, $3,$4,$5,$6)`
	_, e := db.Exec(insertStmt, kodepengguna, namapengguna, alamatpengguna, emailpengguna, katasandi, kodeotp)
	CheckError(e)
	data := []struct {
		Status string
	}{
		{"Successfully Inserted!"},
	}
	output, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

func Login(w http.ResponseWriter, r *http.Request) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	keyVal := make(map[string]string)
	json.Unmarshal(b, &keyVal)
	kodepengguna := keyVal["kodepengguna"]
	katasandi := keyVal["katasandi"]

	logStmt := `select "namapengguna" from "pengguna" where kodepengguna=$1 and katasandi=$2`
	login, e := db.Query(logStmt, kodepengguna, katasandi)
	CheckError(e)
	//defer rows.Close()
	for login.Next() {
		var namapengguna string
		login.Scan(&namapengguna)

		data := []struct {
			NamaPengguna string
		}{
			{namapengguna},
		}
		output, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Header().Set("content-type", "application/json")
		w.Write(output)
	}
}
func EntryItems(w http.ResponseWriter, r *http.Request) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	keyVal := make(map[string]string)
	keyValInt := make(map[string]int)
	json.Unmarshal(b, &keyVal)
	json.Unmarshal(b, &keyValInt)
	kodebarang := keyVal["kodebarang"]
	namabarang := keyVal["namabarang"]
	kategoribarang := keyVal["kategoribarang"]
	hargabarang := keyValInt["hargabarang"]

	insertStmt := `insert into "barang" values($1, $2, $3,$4)`
	_, e := db.Exec(insertStmt, kodebarang, namabarang, kategoribarang, hargabarang)
	CheckError(e)
	data := []struct {
		Status string
	}{
		{"Successfully Inserted Item!"},
	}
	output, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}
func ShowItemsPerCategory(w http.ResponseWriter, r *http.Request) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	keyVal := make(map[string]string)

	json.Unmarshal(b, &keyVal)

	kategoribarang := keyVal["kategoribarang"]

	itemStmt := `select "namabarang", "hargabarang" from "barang" where kategoribarang=$1`
	items, e := db.Query(itemStmt, kategoribarang)
	CheckError(e)

	for items.Next() {
		var namabarang string
		var hargabarang int
		items.Scan(&namabarang, &hargabarang)

		data := []struct {
			NamaBarang  string
			HargaBarang int
		}{
			{namabarang, hargabarang},
		}
		output, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Header().Set("content-type", "application/json")
		w.Write(output)
	}
}
func AddCart(w http.ResponseWriter, r *http.Request) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	keyVal := make(map[string]string)
	json.Unmarshal(b, &keyVal)
	kodebarang := keyVal["kodebarang"]
	kodepengguna := keyVal["kodepengguna"]
	kodekeranjang := keyVal["kodekeranjang"]

	insertCartStmt := `insert into "keranjang" values($1, $2, $3)`
	_, e := db.Exec(insertCartStmt, kodekeranjang, kodebarang, kodepengguna)
	CheckError(e)
	data := []struct {
		Status string
	}{
		{"Successfully Inserted Item into Cart!"},
	}
	output, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}
func DeleteCart(w http.ResponseWriter, r *http.Request) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	keyVal := make(map[string]string)
	json.Unmarshal(b, &keyVal)
	kodebarang := keyVal["kodebarang"]

	removeCartStmt := `delete from "keranjang" where kodebarang=$1`
	_, e := db.Exec(removeCartStmt, kodebarang)
	CheckError(e)
	data := []struct {
		Status string
	}{
		{"Successfully Removed Item in Cart!"},
	}
	output, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}
func Uang(w http.ResponseWriter, r *http.Request) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	keyVal := make(map[string]string)
	json.Unmarshal(b, &keyVal)
	kodebayar := keyVal["kodebayar"]
	metodebayar := keyVal["metodebayar"]
	namapembayar := keyVal["namapembayar"]
	namabank := keyVal["namabank"]
	nomorrek := keyVal["nomorrek"]

	uangStmt := `insert into "uang" values($1, $2, $3,$4,$5)`
	_, e := db.Exec(uangStmt, kodebayar, metodebayar, namapembayar, namabank, nomorrek)
	CheckError(e)
	data := []struct {
		Status string
	}{
		{"Successfully Inserted Money!"},
	}
	output, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}
func Payment(w http.ResponseWriter, r *http.Request) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	keyVal := make(map[string]string)
	json.Unmarshal(b, &keyVal)
	kodetransaksi := keyVal["kodetransaksi"]
	kodekeranjang := keyVal["kodekeranjang"]
	kodebayar := keyVal["kodebayar"]
	bayarStmt := `insert into "transaksi" values($1, $2, $3)`
	_, e := db.Exec(bayarStmt, kodetransaksi, kodekeranjang, kodebayar)
	CheckError(e)
	data := []struct {
		Status string
	}{
		{"Successfully Paid!"},
	}
	output, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}
func ShopList(w http.ResponseWriter, r *http.Request) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	keyVal := make(map[string]string)
	json.Unmarshal(b, &keyVal)
	kodepengguna := keyVal["kodepengguna"]

	lookStmt := `select  namabarang,hargabarang from keranjang as k, barang as b where k.kodebarang=b.kodebarang and k.kodepengguna=$1`
	look, e := db.Query(lookStmt, kodepengguna)
	CheckError(e)

	for look.Next() {
		var namabarang string
		var hargabarang int
		//answer := &Answer{}
		look.Scan(&namabarang, &hargabarang)

		data := []struct {
			NamaBarang  string
			HargaBarang int
		}{
			{namabarang, hargabarang},
		}
		output, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Header().Set("content-type", "application/json")
		w.Write(output)
	}
}
func main() {
	log.Println("Please do something on it !")
	http.HandleFunc("/registration", Registration)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/entryitems", EntryItems)
	http.HandleFunc("/category", ShowItemsPerCategory)
	http.HandleFunc("/addcart", AddCart)
	http.HandleFunc("/deletecart", DeleteCart)
	http.HandleFunc("/money", Uang)
	http.HandleFunc("/payment", Payment)
	http.HandleFunc("/shoplist", ShopList)
	http.ListenAndServe(":5051", nil)
}
