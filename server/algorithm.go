package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

func findMatch(question string, questions []string, answers []string, algo string) (string, error) {
	processedQuestion := processText(question)

	// choose algorithm
	var matchFunc func(string, string) bool
	switch algo {
	case "kmp":
		matchFunc = KMP
	case "bm":
		matchFunc = BM
	default:
		return "", fmt.Errorf("Algoritma tidak valid")
	}

	// check for exact match using chosen algorithm
	for i, q := range questions {
		if matchFunc(processedQuestion, processText(q)) {
			return answers[i], nil
		}
	}

	// check for approximate match using Levenshtein
	bestScore := -1
	bestMatch := ""
	var similarQuestions []string
	for i, q := range questions {
		score := levenshteinDistance(processedQuestion, processText(q))
		maxScore := (len(processedQuestion) + len(q)) / 2
		// score threshold for a good match
		if score <= maxScore/10 && score > bestScore*9/10 {
			bestScore = score
			bestMatch = answers[i]
		} else if score >= maxScore/9 {
			similarQuestions = append(similarQuestions, q)
		}
	}
	if bestMatch != "" {
		return bestMatch, nil
	} else if len(similarQuestions) >= 3 {
		// sort similar questions by Levenshtein distance
		sort.Slice(similarQuestions, func(i, j int) bool {
			return levenshteinDistance(processedQuestion, processText(similarQuestions[i])) < levenshteinDistance(processedQuestion, processText(similarQuestions[j]))
		})
		// select the 3 most similar questions
		selectedQuestions := similarQuestions[:3]
		// construct response message
		var sb strings.Builder
		sb.WriteString("Pertanyaan Anda tidak ditemukan. Pertanyaan yang mirip:\n")
		for _, q := range selectedQuestions {
			sb.WriteString(fmt.Sprintf("- %s\n", q))
		}
		return sb.String(), nil
	}

	return "", fmt.Errorf("Maaf, saya tidak mengerti pertanyaan Anda.")
}

func levenshteinDistance(s, t string) int {
	m := len(s)
	n := len(t)

	if m == 0 {
		return n
	} else if n == 0 {
		return m
	}

	// initialize matrix
	d := make([][]int, m+1)
	for i := range d {
		d[i] = make([]int, n+1)
		d[i][0] = i
	}
	for j := 1; j <= n; j++ {
		d[0][j] = j
	}

	// calculate distance
	for j := 1; j <= n; j++ {
		for i := 1; i <= m; i++ {
			var cost int
			if s[i-1] == t[j-1] {
				cost = 0
			} else {
				cost = 1
			}

			d[i][j] = min(min(d[i-1][j]+1, d[i][j-1]+1), d[i-1][j-1]+cost)
		}
	}

	return d[m][n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func KMP(text string, pattern string) bool {
	if len(pattern) == 0 {
		return true
	}

	if len(text) == 0 {
		return false
	}

	// Compute prefix table
	prefix := make([]int, len(pattern))
	j := 0
	for i := 1; i < len(pattern); i++ {
		for j > 0 && pattern[j] != pattern[i] {
			j = prefix[j-1]
		}
		if pattern[j] == pattern[i] {
			j++
		}
		prefix[i] = j
	}

	// Perform string matching
	j = 0
	for i := 0; i < len(text); i++ {
		for j > 0 && pattern[j] != text[i] {
			j = prefix[j-1]
		}
		if pattern[j] == text[i] {
			j++
		}
		if j == len(pattern) {
			return true
		}
	}

	return false
}

func BM(text string, pattern string) bool {
	n := len(text)
	m := len(pattern)

	if m == 0 {
		return true
	}

	if n < m {
		return false
	}

	// create bad character table
	bc := make(map[byte]int)
	for i := 0; i < m-1; i++ {
		bc[pattern[i]] = m - i - 1
	}

	// search for pattern
	i := m - 1
	j := m - 1
	for i < n {
		if text[i] == pattern[j] {
			if j == 0 {
				return true
			}
			i--
			j--
		} else {
			i += max(bc[text[i]], m-j)
			j = m - 1
		}
	}

	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func processText(text string) string {
	// Remove leading and trailing whitespaces
	text = strings.TrimSpace(text)
	// Convert to lowercase
	text = strings.ToLower(text)
	// Replace multiple whitespaces with a single whitespace
	text = regexp.MustCompile(`\s+`).ReplaceAllString(text, " ")
	// Remove non-alphanumeric and non-whitespace characters
	text = regexp.MustCompile(`[^[:alnum:]\s]`).ReplaceAllString(text, "")

	return text
}

func deleteQuestionCheck(question string) bool {
	pattern := regexp.MustCompile(`^hapus pertanyaan\s(.+)$`)
	return pattern.MatchString(question)
}

func parsingDeleteQuestion(question string) string {
	if deleteQuestionCheck(question) {
		pattern := regexp.MustCompile(`^hapus pertanyaan\s(.+)$`)
		return pattern.FindStringSubmatch(question)[1]
	} else {
		return ""
	}
}

func updateQuestionCheck(question string) bool {
	pattern := regexp.MustCompile(`^tambah pertanyaan\s(.+)\sdengan jawaban\s(.+)$`)
	return pattern.MatchString(question)
}

func parsingUpdateQuestion(question string) []string {
	if updateQuestionCheck(question) {
		pattern := regexp.MustCompile(`^tambah pertanyaan\s(.+)\sdengan jawaban\s(.+)$`)
		array := []string{pattern.FindStringSubmatch(question)[1], pattern.FindStringSubmatch(question)[2]}
		return array
	} else {
		array := []string{"", ""}
		return array
	}
}

func dateCheck(date string) bool {
	var pattern, _ = regexp.Compile(`([0-9]|[0-9][0-9])\/([0-9]|[0-9][0-9])\/([0-9]|[0-9][0-9]|[0-9][0-9][0-9]|[0-9][0-9][0-9][0-9])`)
	return pattern.MatchString(date)
}

func parsingDate(date string) string {
	if dateCheck(date) {
		pattern, _ := regexp.Compile(`([0-9]|[0-9][0-9])\/([0-9]|[0-9][0-9])\/([0-9]|[0-9][0-9]|[0-9][0-9][0-9]|[0-9][0-9][0-9][0-9])`)
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
	dateParts := strings.Split(dateString, "/")
	day, _ := strconv.Atoi(dateParts[0])
	month, _ := strconv.Atoi(dateParts[1])
	year, _ := strconv.Atoi(dateParts[2])
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

	if year%10 == year {
		yearString = "000" + strconv.Itoa(year)
	} else if year%100 == year {
		yearString = "0" + strconv.Itoa(year)
	} else if year%1000 == year {
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
	var daysOfWeek = []string{"Minggu", "Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu"}
	return daysOfWeek[h]
}

func calculatorCheck(calculator string) bool {
	pattern := regexp.MustCompile(`^(?:\d+|\(\s*(?:(?:\d+|[+\-*/^()])\s*)+\))(?:\s*[+\-*/^]\s*(?:\d+|\(\s*(?:(?:\d+|[+\-*/^()])\s*)+\)))*$`)
	return pattern.MatchString(calculator)
}

func parsingCalculator(calculator string) string {
	if calculatorCheck(calculator) {
		pattern := regexp.MustCompile(`^(?:\d+|\(\s*(?:(?:\d+|[+\-*/^()])\s*)+\))(?:\s*[+\-*/^]\s*(?:\d+|\(\s*(?:(?:\d+|[+\-*/^()])\s*)+\)))*$`)
		return pattern.FindStringSubmatch(calculator)[0]
	} else {
		return ""
	}
}



func deleteQuestionCheck(question string) bool {
	pattern := regexp.MustCompile(`^hapus pertanyaan\s(.+)$`)
	return pattern.MatchString(question)
}

func parsingDeleteQuestion(question string) string {
	if deleteQuestionCheck(question) {
		pattern := regexp.MustCompile(`^hapus pertanyaan\s(.+)$`)
		return pattern.FindStringSubmatch(question)[1]
	} else {
		return ""
	}
}

func updateQuestionCheck(question string) bool {
	pattern := regexp.MustCompile(`^tambah pertanyaan\s(.+)\sdengan jawaban\s(.+)$`)
	return pattern.MatchString(question)
}

func parsingUpdateQuestion(question string) []string {
	if updateQuestionCheck(question) {
		pattern := regexp.MustCompile(`^tambah pertanyaan\s(.+)\sdengan jawaban\s(.+)$`)
		array := []string{pattern.FindStringSubmatch(question)[1], pattern.FindStringSubmatch(question)[2]}
		return array
	} else {
		array := []string{"", ""}
		return array
	}
}

func dateCheck(date string) bool {
	var pattern, _ = regexp.Compile(`([0-9]|[0-9][0-9])\/([0-9]|[0-9][0-9])\/([0-9]|[0-9][0-9]|[0-9][0-9][0-9]|[0-9][0-9][0-9][0-9])`)
	return pattern.MatchString(date)
}

func parsingDate(date string) string {
	if dateCheck(date) {
		pattern, _ := regexp.Compile(`([0-9]|[0-9][0-9])\/([0-9]|[0-9][0-9])\/([0-9]|[0-9][0-9]|[0-9][0-9][0-9]|[0-9][0-9][0-9][0-9])`)
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
	var daysOfWeek = []string{"Minggu", "Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu"}
	return daysOfWeek[h]
}

func calculatorCheck(calculator string) bool {
	pattern := regexp.MustCompile(`^(?:\d+|\(\s*(?:(?:\d+|[+\-*/^()])\s*)+\))(?:\s*[+\-*/^]\s*(?:\d+|\(\s*(?:(?:\d+|[+\-*/^()])\s*)+\)))*$`)
	return pattern.MatchString(calculator)
}

func parsingCalculator(calculator string) string {
	if calculatorCheck(calculator) {
		pattern := regexp.MustCompile(`^(?:\d+|\(\s*(?:(?:\d+|[+\-*/^()])\s*)+\))(?:\s*[+\-*/^]\s*(?:\d+|\(\s*(?:(?:\d+|[+\-*/^()])\s*)+\)))*$`)
		return pattern.FindStringSubmatch(calculator)[0]
	} else {
		return ""
	}
}

// func main() {
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Println("Selamat datang di ChatGPT!")
// 	for {
// 		fmt.Print("Anda: ")
// 		input, _ := reader.ReadString('\n')
// 		input = strings.TrimSpace(input)
// 		if input == "" {
// 			continue
// 		}
// 		answer := findAnswerBM(input, questions, answers)
// 		fmt.Println("ChatGPT:", answer)
// 	}
// }
