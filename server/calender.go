package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func dateCheck(date string) bool {
	var pattern, _ = regexp.Compile(`([0-9]|[0-9][0-9])\/([0-9]|[0-9][0-9])\/\d{1,10}`)
	return pattern.MatchString(date)
}

func parsingDate(date string) string {
	if dateCheck(date) {
		pattern, _ := regexp.Compile(`([0-9]|[0-9][0-9])\/([0-9]|[0-9][0-9])\/\d{1,10}`)
		return pattern.FindStringSubmatch(date)[0]
	} else {
		return ""
	}
}

func isValidDate(dateString string) bool {
	_, err := time.Parse("02/01/2006", dateString)
	return err == nil
}

func parsingValidDate(dateString string) string {
	fmt.Println(dateString)
	dateParts := strings.Split(dateString, "/")
	day, _ := strconv.Atoi(dateParts[0])
	month, _ := strconv.Atoi(dateParts[1])
	year, _ := strconv.Atoi(dateParts[2])
	fmt.Println(day, month, year)
	var dayString, monthString, yearString string
	if day%10 == day {
		dayString = "0" + strconv.Itoa(day)
	} else {
		dayString = strconv.Itoa(day)
	}

	if month%10 == month {
		monthString = "0" + strconv.Itoa(month)
	} else {
		monthString = strconv.Itoa(month)
	}

	if year < 10 {
		yearString = "000" + strconv.Itoa(year)
	} else if year < 100 {
		yearString = "0" + strconv.Itoa(year)
	} else if year < 1000 {
		yearString = "0" + strconv.Itoa(year)
	} else {
		yearString = strconv.Itoa(year)
	}

	return dayString + "/" + monthString + "/" + yearString
}

func getDay(date string) string {
	if dateCheck(date) == false {
		return "Format tanggal salah"
	}
	// Mengubah bulan Januari dan Februari menjadi bulan ke-13 dan ke-14
	// dan mengurangi tahun sebanyak 1 untuk perhitungan
	dateParts := strings.Split(date, "/")
	day, _ := strconv.Atoi(dateParts[0])
	month, _ := strconv.Atoi(dateParts[1])
	year, _ := strconv.Atoi(dateParts[2])

	if month == 1 || month == 2 {
		month += 12
		year -= 1
	}
	// Menghitung hari dalam minggu menggunakan rumus Zeller's congruence
	// Rumus: h = (q + ((13*(m+1))/5) + K + (K/4) + (J/4) - 2*J) mod 7
	// K = tahun % 100, J = tahun / 100
	var q = day
	var m = month
	var K = year % 100
	var J = year / 100

	var h = (q + ((13 * (m + 1)) / 5) + K + (K / 4) + (J / 4) - 2*J) % 7

	// Menentukan nama hari berdasarkan nilai h
	var daysOfWeek = []string{"Sabtu", "Minggu", "Senin", "Selasa", "Rabu", "Kamis", "Jumat"}
	return daysOfWeek[h]
}
