package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

type (
	PertanyaanReq struct {
		Pertanyaan string `json:"pertanyaan"`
	}
	Pertanyaan struct {
		Pertanyaan string `json:"pertanyaan"`
		Jawaban    string `json:"jawaban"`
	}
	HistoriReq struct {
		Nama string `json:"nama"`
	}
	Histori struct {
		Nama string   `json:"nama"`
		Isi  []Respon `json:"isi"`
	}
	Respon struct {
		ID_histori int    `json:"id_histori"`
		Jenis      string `json:"jenis"`
		Isi        string `json:"isi"`
	}
)

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", os.Getenv("DATABASE_URL")+"/tubes3")
	if err != nil {
		fmt.Println("error")
	}
	return db, err
}

func findAnswer(c echo.Context) error {
	var quest PertanyaanReq
	err := c.Bind(&quest)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "error 1")
	}

	db, err := connect()
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "error 2")
	}
	defer db.Close()

	row := db.QueryRow("SELECT jawaban FROM pertanyaan WHERE pertanyaan = ?", quest.Pertanyaan)
	var answer string
	err = row.Scan(&answer)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "error 3")
	}

	return c.JSON(http.StatusOK, answer)
}

func addQuestion(c echo.Context) error {
	var quest Pertanyaan
	err := c.Bind(&quest)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "error 1")
	}

	db, err := connect()
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "error 2")
	}
	defer db.Close()

	rows, err := db.Query("select Pertanyaan,Jawaban from pertanyaan where Pertanyaan=? ", quest.Pertanyaan)
	if err != nil {
		fmt.Println(err.Error())
		return c.JSON(http.StatusUnprocessableEntity, "error 2")
	}
	defer rows.Close()

	var result []Pertanyaan
	fmt.Println(rows)
	if rows.Next() {
		var each = Pertanyaan{}
		var err = rows.Scan(&each.Pertanyaan, &each.Jawaban)

		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(http.StatusUnprocessableEntity, "error 3")
		}

		result = append(result, each)
	}
	fmt.Println(result)
	if err = rows.Err(); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "error 3")
	}

	if result != nil {
		update, err := db.Exec("UPDATE Pertanyaan SET Jawaban = ? WHERE Pertanyaan = ?", quest.Jawaban, quest.Pertanyaan)
		if err != nil {
			return c.JSON(http.StatusUnprocessableEntity, "error 2")
		}
		defer rows.Close()
		update.RowsAffected()
		return c.JSON(http.StatusOK, quest)
	}

	_, err = db.Exec("INSERT INTO pertanyaan (Pertanyaan,Jawaban) VALUES (?,?)", quest.Pertanyaan, quest.Jawaban)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "error 3")
	}

	return c.JSON(http.StatusCreated, quest)
}

func addRespon(c echo.Context) error {
	var respon Respon
	err := c.Bind(&respon)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "error 1")
	}

	db, err := connect()
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "error 2")
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO respon (ID_histori,Jenis,Isi) VALUES (?,?,?)", respon.ID_histori, respon.Jenis, respon.Isi)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "error 3")
	}

	return c.JSON(http.StatusCreated, respon)
}

func showHistori(c echo.Context) error {
	var histori HistoriReq
	err := c.Bind(&histori)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "error 1")
	}
	db, err := connect()
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "error 1")
	}
	defer db.Close()
	rows, err := db.Query("select h.ID_histori, r.Jenis, r.Isi from histori as h, respon as r where h.ID_histori=r.ID_histori AND h.Nama=?", histori.Nama)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "error 2")
	}
	defer rows.Close()
	var isi []Respon
	for rows.Next() {
		var respon Respon
		err := rows.Scan(&respon.ID_histori, &respon.Jenis, &respon.Isi)
		if err != nil {
			return c.JSON(http.StatusUnprocessableEntity, "error 3")
		}
		isi = append(isi, respon)
	}
	historiReq := &Histori{
		Nama: histori.Nama,
		Isi:  isi,
	}
	return c.JSON(http.StatusOK, historiReq)
}

func addHistori(c echo.Context) error {
	var histori HistoriReq
	err := c.Bind(&histori)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "error 1")
	}

	db, err := connect()
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "error 2")
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO histori (Nama) VALUES (?)", histori.Nama)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "error 3")
	}

	return c.JSON(http.StatusCreated, histori)
}

func deleteHistori(c echo.Context) error {
	var histori HistoriReq
	err := c.Bind(&histori)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "error 1")
	}

	db, err := connect()
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "error 2")
	}
	defer db.Close()

	id, err := db.Query("select ID_histori from histori where Nama=?", histori.Nama)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "error 2")
	}
	defer id.Close()
	var id_histori int
	for id.Next() {
		err := id.Scan(&id_histori)
		if err != nil {
			return c.JSON(http.StatusUnprocessableEntity, "error 3")
		}
	}
	_, err = db.Exec("DELETE FROM respon WHERE ID_histori=?", id_histori)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "error 3")
	}

	_, err = db.Exec("DELETE FROM histori WHERE Nama=?", histori.Nama)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "error 3")
	}

	return c.JSON(http.StatusCreated, histori)
}

func main() {
	e := echo.New()

	e.GET("find", findAnswer)
	e.POST("quest", addQuestion)
	e.POST("respon", addRespon)
	e.GET("histori", showHistori)
	e.POST("histori", addHistori)
	e.DELETE("histori", deleteHistori)
	e.Start(":1234")
}
