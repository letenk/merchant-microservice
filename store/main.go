package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getDetailToko(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"nama":           "Jabutech store",
		"jumlah_product": "10",
	})
}

func getAllToko(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]map[string]string{
		{
			"nama":           "Jabutech store",
			"jumlah_product": "10",
		},
		{
			"nama":           "Aisyah store",
			"jumlah_product": "30",
		},
	})
}

func main() {
	var mux = http.NewServeMux()

	mux.HandleFunc("/get-detail-store", getDetailToko)
	mux.HandleFunc("/get-all-store", getAllToko)

	fmt.Println("Server is running.")

	http.ListenAndServe(":8081", mux)
}
