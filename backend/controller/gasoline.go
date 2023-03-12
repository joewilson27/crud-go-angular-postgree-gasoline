package controller

import (
	"encoding/json" // package untuk enkode dan mendekode json menjadi struct dan sebaliknya
	"fmt"
	"strconv" // package yang digunakan untuk mengubah string menjadi tipe int

	"log"
	"net/http" // digunakan untuk mengakses objek permintaan dan respons dari api

	"go-postgres-crud/models" //models package dimana Gasoline didefinisikan

	"github.com/gorilla/mux" // digunakan untuk mendapatkan parameter dari router
	_ "github.com/lib/pq"    // postgres golang driver
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

type Response struct {
	Status  int               `json:"status"`
	Message string            `json:"message"`
	Data    []models.Gasoline `json:"data"`
}

// TambahData
func TmbhData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// create an empty data of type models.Gasoline
	// kita buat empty gasoline dengan tipe models.Gasoline
	var gas models.Gasoline

	// decode data json request ke data
	err := json.NewDecoder(r.Body).Decode(&gas)

	if err != nil {
		log.Fatalf("Tidak bisa mendecode dari request body.  %v", err)
	}

	// panggil modelsnya lalu insert data
	models.TambahData(gas)
	//insertID := models.TambahData(gas)

	// format response objectnya
	// res := response{
	// 	ID:      insertID,
	// 	Message: "Data telah ditambahkan",
	// }

	fmt.Fprintf(w, "New data was created")

	// kirim response
	// json.NewEncoder(w).Encode(res)
}

// AmbilData mengambil single data dengan parameter id
func AmbilData(w http.ResponseWriter, r *http.Request) {
	// kita set headernya
	w.Header().Set("Content-Type", "application/json")
	//w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// dapatkan iddata dari parameter request, keynya adalah "id"
	params := mux.Vars(r)

	// konversi id dari string ke int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Tidak bisa mengubah dari string ke int.  %v", err)
	}

	// memanggil models ambilsatudata dengan parameter id yg nantinya akan mengambil single data
	gas, err := models.AmbilSatuData(int64(id))

	if err != nil {
		log.Fatalf("Tidak bisa mengambil data data. %v", err)
	}

	// kirim response
	json.NewEncoder(w).Encode(gas)
}

// Ambil semua data
func AmbilSemuaData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// memanggil models AmbilSemuaData
	gass, err := models.AmbilSemuaData()

	if err != nil {
		log.Fatalf("Tidak bisa mengambil data. %v", err)
	}

	// var response Response
	// response.Status = 1
	// response.Message = "Success"
	// response.Data = gass

	// kirim semua response
	json.NewEncoder(w).Encode(gass)
}

func UpdateData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// kita ambil request parameter idnya
	params := mux.Vars(r)

	// konversikan ke int yang sebelumnya adalah string
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Tidak bisa mengubah dari string ke int.  %v", err)
	}

	// buat variable gas dengan type models.Gasoline
	var gas models.Gasoline

	// decode json request ke variable data
	err = json.NewDecoder(r.Body).Decode(&gas)

	if err != nil {
		log.Fatalf("Tidak bisa decode request body.  %v", err)
	}

	// panggil updatedata untuk mengupdate data
	updatedRows := models.UpdateData(int64(id), gas)

	// ini adalah format message berupa string
	msg := fmt.Sprintf("Data telah berhasil diupdate. Jumlah yang diupdate %v rows/record", updatedRows)

	// ini adalah format response message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// kirim berupa response
	json.NewEncoder(w).Encode(res)
}

func HapusData(w http.ResponseWriter, r *http.Request) {

	// kita ambil request parameter idnya
	params := mux.Vars(r)

	// konversikan ke int yang sebelumnya adalah string
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Tidak bisa mengubah dari string ke int.  %v", err)
	}

	// panggil fungsi hapusData , dan convert int ke int64
	deletedRows := models.HapusData(int64(id))

	// ini adalah format message berupa string
	msg := fmt.Sprintf("Data sukses di hapus. Total data yang dihapus %v", deletedRows)

	// ini adalah format reponse message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}
