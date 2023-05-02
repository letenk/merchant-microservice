package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// middleware merchant
func merchantMiddle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		author := r.Header.Get("Authorization")

		if author != "merchant" {
			w.Write([]byte("Forbidden"))
			return
		}

		next.ServeHTTP(w, r)
	}
}

// middleware super admin
func superMiddle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		author := r.Header.Get("Authorization")

		if author != "su-admin" {
			w.Write([]byte("Forbidden"))
			return
		}

		next.ServeHTTP(w, r)
	}
}

type Merchant struct {
	Nama          string `json:"nama"`
	NamaToko      string `json:"nama_toko"`
	JumlahProduct string `json:"jumlah_product"`
}

func getMerchant(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://localhost:8000/api/get-merchant")
	if err != nil {
		log.Fatal("Err: ", err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Err: ", err)
	}

	merch := &Merchant{}
	json.Unmarshal(data, merch)

	json.NewEncoder(w).Encode(merch)
}

type Store struct {
	Nama          string `json:"nama"`
	JumlahProduct string `json:"jumlah_product"`
}

func getAllStore(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://localhost:8081/get-all-store")
	if err != nil {
		log.Fatal("Err: ", err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Err: ", err)
	}

	store := &[]Store{}
	json.Unmarshal(data, store)

	json.NewEncoder(w).Encode(store)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/merchants", merchantMiddle(getMerchant))
	mux.HandleFunc("/store", superMiddle(getAllStore))

	fmt.Println("Api Gateway is running")

	http.ListenAndServe(":8080", mux)
}
