package models

import (
	"database/sql"
	"fmt"
	"go-postgres-crud/config"
	"log"

	_ "github.com/lib/pq" // postgres golang driver
)

// data schema dari tabel Gasoline
// kita coba dengan jika datanya null
// jika return datanya ada yg null, silahkan pake NullString, contohnya dibawah
// Penulis       config.NullString `json:"penulis"`
type Gasoline struct {
	ID           int64 `json:"id"`
	Jumlah_liter int64 `json:"jumlahLiter"`
	Premium      int64 `json:"premium"`
	Pertalite    int64 `json:"pertalite"`
}

func TambahData(gas Gasoline) int64 {

	// mengkoneksikan ke db postgres
	db := config.CreateConnection()

	// kita tutup koneksinya di akhir proses
	defer db.Close()

	// kita buat insert query
	// mengembalikan nilai id akan mengembalikan id dari data yang dimasukkan ke db
	sqlStatement := `INSERT INTO gasoline (jumlah_liter, premium, pertalite) VALUES ($1, $2, $3) RETURNING id`

	// id yang dimasukkan akan disimpan di id ini
	var id int64

	// Scan function akan menyimpan insert id didalam id id
	err := db.QueryRow(sqlStatement, gas.Jumlah_liter, gas.Premium, gas.Pertalite).Scan(&id)

	if err != nil {
		log.Fatalf("Tidak Bisa mengeksekusi query. %v", err)
	}

	fmt.Printf("Insert data single record %v", id)

	// return insert id
	return id
}

// ambil semua data
func AmbilSemuaData() ([]Gasoline, error) {
	// mengkoneksikan ke db postgres
	db := config.CreateConnection()

	// kita tutup koneksinya di akhir proses
	defer db.Close()

	var gass []Gasoline

	// kita buat select query
	sqlStatement := `SELECT * FROM gasoline`

	// mengeksekusi sql query
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("tidak bisa mengeksekusi query. %v", err)
	}

	// kita tutup eksekusi proses sql qeurynya
	defer rows.Close()

	// kita iterasi mengambil datanya
	for rows.Next() {
		var gas Gasoline

		// kita ambil datanya dan unmarshal ke structnya
		err = rows.Scan(&gas.ID, &gas.Jumlah_liter, &gas.Premium, &gas.Pertalite)

		if err != nil {
			log.Fatalf("tidak bisa mengambil data. %v", err)
		}

		// masukkan kedalam slice gass
		gass = append(gass, gas)

	}

	// return empty gas atau jika error
	return gass, err
}

// mengambil satu data
func AmbilSatuData(id int64) (Gasoline, error) {
	// mengkoneksikan ke db postgres
	db := config.CreateConnection()

	// kita tutup koneksinya di akhir proses
	defer db.Close()

	var gas Gasoline

	// buat sql query
	sqlStatement := `SELECT * FROM gasoline WHERE id=$1`

	// eksekusi sql statement
	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&gas.ID, &gas.Jumlah_liter, &gas.Premium, &gas.Pertalite)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("Tidak ada data yang dicari!")
		return gas, nil
	case nil:
		return gas, nil
	default:
		log.Fatalf("tidak bisa mengambil data. %v", err)
	}

	return gas, err
}

// update user in the DB
func UpdateData(id int64, gas Gasoline) int64 {

	// mengkoneksikan ke db postgres
	db := config.CreateConnection()

	// kita tutup koneksinya di akhir proses
	defer db.Close()

	// kita buat sql query create
	sqlStatement := `UPDATE gasoline SET jumlah_liter=$2, premium=$3, pertalite=$4 WHERE id=$1`

	// eksekusi sql statement
	res, err := db.Exec(sqlStatement, id, gas.Jumlah_liter, gas.Premium, gas.Pertalite)

	if err != nil {
		log.Fatalf("Tidak bisa mengeksekusi query. %v", err)
	}

	// cek berapa banyak row/data yang diupdate
	rowsAffected, err := res.RowsAffected()

	//kita cek
	if err != nil {
		log.Fatalf("Error ketika mengecheck rows/data yang diupdate. %v", err)
	}

	fmt.Printf("Total rows/record yang diupdate %v\n", rowsAffected)

	return rowsAffected
}

func HapusData(id int64) int64 {

	// mengkoneksikan ke db postgres
	db := config.CreateConnection()

	// kita tutup koneksinya di akhir proses
	defer db.Close()

	// buat sql query
	sqlStatement := `DELETE FROM gasoline WHERE id=$1`

	// eksekusi sql statement
	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("tidak bisa mengeksekusi query. %v", err)
	}

	// cek berapa jumlah data/row yang di hapus
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("tidak bisa mencari data. %v", err)
	}

	fmt.Printf("Total data yang terhapus %v", rowsAffected)

	return rowsAffected
}
