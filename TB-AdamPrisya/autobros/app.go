package main

import (
	"autobros/controls/kendaraancontroller"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()
		router.HandleFunc("/", kendaraancontroller.Dashboard).Methods("GET")
		router.HandleFunc("/autobros", kendaraancontroller.Dashboard).Methods("GET")
		router.HandleFunc("/autobros/tables", kendaraancontroller.Tables).Methods("GET")
		router.HandleFunc("/autobros/tambahservis", kendaraancontroller.Forms).Methods("GET")
		router.HandleFunc("/autobros/tambahservis", kendaraancontroller.Forms).Methods("POST")
		router.HandleFunc("/autobros/tables/updateservis", kendaraancontroller.FormAdmin).Methods("GET")
		router.HandleFunc("/autobros/tables/updateservis", kendaraancontroller.FormAdmin).Methods("POST")
	// no method functional
		router.HandleFunc("/autobros/tables/hapusservis", kendaraancontroller.Delete)
		router.HandleFunc("/autobros/tables/detaildata/hapusservis", kendaraancontroller.Delete)
		router.HandleFunc("/autobros/tables/detaildata", kendaraancontroller.Details)
	http.ListenAndServe(":8000", router)
}
