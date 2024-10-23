package main

import (
	"final-project-backend/database"
	"fmt"
)

func main() {
	db := database.ConnectToDb()
	fmt.Println("Koneksi ke database berhasil", db)
}
