package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	PertanyaanReq struct {
		Pertanyaan string `json:"pertanyaan"`
	}
	Pertanyaan struct {
		Pertanyaan string `json:"pertanyaan"`
		Jawaban    string `json:"jawaban"`
	}
	PertayaanHistori struct {
		Pertanyaan string `json:"pertanyaan"`
		ID_histori int    `json:"id_histori"`
		Jenis      string `json:"jenis"`
	}
	HistoriReq struct {
		Nama string `json:"nama"`
	}

	Histori struct {
		Nama string   `json:"nama"`
		Isi  []Respon `json:"isi"`
	}
	UpdateHistori struct {
		NewName    string `json:"new_name"`
		ID_histori int    `json:"ID_histori"`
	}
	Respon struct {
		ID_histori int    `json:"id_histori"`
		Jenis      string `json:"jenis"`
		Isi        string `json:"isi"`
	}

	HistoriReqId struct {
		ID_histori int `json:"id_histori"`
	}

	HistoriId struct {
		ID_histori int    `json:"id"`
		Nama       string `json:"nama"`
	}
)

func connect() (*sql.DB, error) {
	// db, err := sql.Open("mysql", "root:qwerty@tcp(127.0.0.1:3306)"+"/tubes3")
	db, err := sql.Open("mysql", "poe:ocetengkyu@tcp(104.248.157.133:3306)"+"/teman2lemon")

	if err != nil {
		fmt.Println("error")
	}
	return db, err
}

func getAllHistori(c echo.Context) error {
	db, err := connect()
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "error 1")
	}
	defer db.Close()
	rows, err := db.Query("select * from histori")
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "error 2")
	}
	defer rows.Close()
	var isi []HistoriId
	for rows.Next() {
		var respon HistoriId
		err := rows.Scan(&respon.ID_histori, &respon.Nama)
		if err != nil {
			return c.JSON(http.StatusUnprocessableEntity, "error 3")
		}
		isi = append(isi, respon)
	}
	return c.JSON(http.StatusOK, isi)
}

func findAnswer(c echo.Context) error {
	var questHistori PertayaanHistori
	err := c.Bind(&questHistori)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "error 1")
	}

	var quest PertanyaanReq
	quest.Pertanyaan = questHistori.Pertanyaan

	//add respon question
	var statusRespon string
	var respon Respon
	respon.ID_histori = questHistori.ID_histori
	respon.Jenis = "input"
	respon.Isi = quest.Pertanyaan
	statusRespon = addResponReq(respon)
	if statusRespon == "success" {
		fmt.Println("Berhasil add respon")
	} else {
		fmt.Println("Gagal add respon")
	}

	quest.Pertanyaan = strings.ToLower(quest.Pertanyaan)

	var listPertanyaanInput []string
	listPertanyaanInput = strings.Split(quest.Pertanyaan, "\n")
	var resultAnswer string
	if len(listPertanyaanInput) == 1 {
		resultAnswer = parsingAnswer(quest.Pertanyaan, questHistori.ID_histori, questHistori.Jenis, c)
	} else {
		var resultAnswer string
		for i := 0; i < len(listPertanyaanInput); i++ {
			resultAnswer = resultAnswer + "Respon " + strconv.Itoa(i+1) + " " + parsingAnswer(listPertanyaanInput[i], questHistori.ID_histori, questHistori.Jenis, c) + "\n"
		}
	}

	// nulis respon answer
	var respon2 Respon
	respon2.ID_histori = questHistori.ID_histori
	respon2.Jenis = "output"
	respon2.Isi = resultAnswer
	statusRespon = addResponReq(respon2)
	if statusRespon == "success" {
		fmt.Println("Berhasil add respon")
	} else {
		fmt.Println("Gagal add respon")
	}
	//aman
	// return parsingAnswer(quest.Pertanyaan, questHistori.ID_histori, questHistori.Jenis, c)
	return c.JSON(http.StatusOK, resultAnswer)
}
func parsingAnswer(questionInput string, ID_histori int, jenis string, c echo.Context) string {
	if dateCheck(questionInput) {
		var str1 string
		str1 = parsingValidDate(parsingDate(questionInput))
		fmt.Println(str1)
		if isValidDate(str1) {
			return "Hari dari tanggal " + parsingDate(questionInput) + " adalah " + getDay(parsingDate(questionInput))
		} else {
			return "Tanggal " + parsingDate(questionInput) + " tidak valid"
		}
	} else if updateQuestionCheck(questionInput) { // aman
		var question Pertanyaan
		question.Jawaban = parsingUpdateQuestion(questionInput)[1]
		question.Pertanyaan = parsingUpdateQuestion(questionInput)[0]
		var status string

		if questionCheck(question.Pertanyaan) {
			status = addQuestionReq(question)
			if status == "success" {
				fmt.Println("Berhasil update question")
			} else {
				fmt.Println("Gagal update question")
			}
			return "Pertanyaan " + question.Pertanyaan + " sudah ada! Jawaban diupdate ke " + question.Jawaban
		} else {
			status = addQuestionReq(question)
			if status == "success" {
				fmt.Println("Berhasil update question")
			} else {
				fmt.Println("Gagal update question")
			}
			return "Pertanyaan " + question.Pertanyaan + " telah ditambahkan"
		}
	} else if deleteQuestionCheck(questionInput) {
		var question string
		question = parsingDeleteQuestion(questionInput)
		var statusDelete string

		if questionCheck(question) {
			statusDelete = deleteQuestionReq(question)
			if statusDelete == "success" {
				fmt.Println("Berhasil delete question")
			} else {
				fmt.Println("Gagal delete question")
			}
			return "Pertanyaan " + question + " telah dihapus"
		} else {
			return "Tidak ada pertanyaan " + question + " dalam database"
		}
	} else if allMath(filterMath(questionInput)) && filterMath(questionInput) != "" {
		var strOperation string
		strOperation = filterMath(questionInput)
		fmt.Println(strOperation, calculatorCheck(strOperation))
		if calculatorCheck(strOperation) {
			result, e := calculator(strOperation)
			if e == nil {
				return "Hasil dari " + strOperation + " adalah " + strconv.FormatFloat(result, 'f', 2, 64)
			} else {
				return "Input tidak valid"
			}
		} else {
			return "Input tidak valid"
		}
	} else {
		jenisSearching := jenis
		db, err := connect()
		if err != nil {
			return "error database"
		}
		defer db.Close()

		rows, err := db.Query("select Pertanyaan,Jawaban from pertanyaan")
		if err != nil {
			fmt.Println(err.Error())
		}
		defer rows.Close()

		var result []Pertanyaan
		fmt.Println(rows)
		for rows.Next() {
			var each = Pertanyaan{}
			var err = rows.Scan(&each.Pertanyaan, &each.Jawaban)
			if err != nil {
				fmt.Println(err.Error())
			}
			result = append(result, each)
		}
		var pertanyaanList []string
		for _, item := range result {
			pertanyaanList = append(pertanyaanList, item.Pertanyaan)
		}
		var jawabanList []string
		for _, item := range result {
			jawabanList = append(jawabanList, item.Jawaban)
		}
		var answer string
		if jenisSearching == "1" {
			questionSimilar, hasil := findMatch(questionInput, pertanyaanList, jawabanList, "kmp")
			if hasil == nil {
				answer = questionSimilar
			} else {
				answer = "Maaf, saya tidak mengerti pertanyaan anda"
			}
			fmt.Println("answer", answer)
		} else {
			questionSimilar, hasil := findMatch(questionInput, pertanyaanList, jawabanList, "bm")
			if hasil == nil {
				answer = questionSimilar
			} else {
				answer = "Maaf, saya tidak mengerti pertanyaan anda"
			}
			fmt.Println("answer", answer)
		}
		return answer
	}
}

func questionCheck(question string) bool {
	db, err := connect()
	if err != nil {
		fmt.Println("error")
	}
	defer db.Close()

	rows, err := db.Query("select Pertanyaan,Jawaban from pertanyaan where Pertanyaan=? ", question)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer rows.Close()

	var result []Pertanyaan
	fmt.Println(rows)
	if rows.Next() {
		var each = Pertanyaan{}
		var err = rows.Scan(&each.Pertanyaan, &each.Jawaban)
		if err != nil {
			fmt.Println(err.Error())
		}
		result = append(result, each)
	}
	if len(result) > 0 {
		return true
	} else {
		return false
	}
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

func deleteQuestionReq(question string) string {
	db, err := connect()
	if err != nil {
		return "err"
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM pertanyaan WHERE pertanyaan = ?", question)
	if err != nil {
		fmt.Println("gagal hapus")
		return "err"
	}

	return "success"
}

func addQuestionReq(quest Pertanyaan) string {
	db, err := connect()
	if err != nil {
		return "err"
	}
	defer db.Close()

	rows, err := db.Query("select Pertanyaan,Jawaban from pertanyaan where Pertanyaan=? ", quest.Pertanyaan)
	if err != nil {
		fmt.Println(err.Error())
		return "err"
	}
	defer rows.Close()

	var result []Pertanyaan
	fmt.Println(rows)
	if rows.Next() {
		var each = Pertanyaan{}
		var err = rows.Scan(&each.Pertanyaan, &each.Jawaban)

		if err != nil {
			fmt.Println(err.Error())
			return "err"
		}

		result = append(result, each)
	}
	fmt.Println(result)
	if err = rows.Err(); err != nil {
		return "err"
	}

	if result != nil {
		fmt.Println("udah adaaa")
		update, err := db.Exec("UPDATE pertanyaan SET jawaban = ? WHERE pertanyaan = ?", quest.Jawaban, quest.Pertanyaan)
		fmt.Println("update ", update)
		fmt.Println("error", err)
		if err != nil {
			return "err"
		}
		defer rows.Close()
		update.RowsAffected()
		return "success"
	}

	_, err = db.Exec("INSERT INTO pertanyaan (Pertanyaan,Jawaban) VALUES (?,?)", quest.Pertanyaan, quest.Jawaban)
	if err != nil {
		return "err"
	}

	return "success"
}
func addResponReq(respon Respon) string {

	db, err := connect()
	if err != nil {
		return "err"
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO respon (ID_histori,Jenis,Isi) VALUES (?,?,?)", respon.ID_histori, respon.Jenis, respon.Isi)
	if err != nil {
		return "err"
	}

	return "success"
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

func getChatFromId(c echo.Context) error {
	var histori HistoriReqId
	err := c.Bind(&histori)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "error 1")
	}
	db, err := connect()
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "error 1")
	}
	defer db.Close()
	rows, err := db.Query("select h.ID_histori, r.Jenis, r.Isi from histori as h, respon as r where h.ID_histori=r.ID_histori AND h.ID_histori=?", histori.ID_histori)
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
	return c.JSON(http.StatusOK, isi)
}

func showHistori(c echo.Context) error {
	HistoriID := c.QueryParam("Id_histori")
	db, err := connect()
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "error 1")
	}
	defer db.Close()
	rows, err := db.Query("select h.ID_histori, r.Jenis, r.Isi from histori as h, respon as r where h.ID_histori=r.ID_histori AND h.ID_histori=?", HistoriID)
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
	nama, err := db.Query("select Nama from histori where ID_histori=?", HistoriID)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "error 2")
	}
	defer nama.Close()
	var namaHistori string
	for nama.Next() {
		err := nama.Scan(&namaHistori)
		if err != nil {
			return c.JSON(http.StatusUnprocessableEntity, "error 3")
		}
	}
	historiReq := &Histori{
		Nama: namaHistori,
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
	id_histori := c.QueryParam("Id_histori")
	db, err := connect()
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "error 1")
	}
	_, err = db.Exec("DELETE FROM respon WHERE ID_histori=?", id_histori)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "error 3")
	}

	_, err = db.Exec("DELETE FROM histori WHERE ID_histori=?", id_histori)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "error 3")
	}

	return c.JSON(http.StatusCreated, "successs")
}

func updateHistoriName(c echo.Context) error {
	var historiChange UpdateHistori
	err := c.Bind(&historiChange)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "error 1")
	}

	db, err := connect()
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "error 2")
	}
	defer db.Close()

	update, err := db.Exec("UPDATE histori SET Nama = ? WHERE ID_histori = ?", historiChange.NewName, historiChange.ID_histori)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "error 2")
	}
	update.RowsAffected()
	return c.JSON(http.StatusOK, "success")
}

func main() {
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	e.GET("histori", showHistori)
	e.GET("chat", getChatFromId)
	e.GET("allhistori", getAllHistori)
	e.POST("find", findAnswer)
	e.POST("quest", addQuestion)
	e.POST("respon", addRespon)
	e.POST("histori", addHistori)
	e.PUT("histori", updateHistoriName)
	e.DELETE("histori", deleteHistori)
	e.Logger.Fatal(e.Start(":1234"))
}
