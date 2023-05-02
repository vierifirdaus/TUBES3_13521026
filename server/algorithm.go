package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// mencari pertanyaan paling mirip menggunakan algoritma KMP
func findMostSimilarKMP(question string, questions []string) (string, bool) {
	var mostSimilarQuestion string
	var similarity float64
	for _, q := range questions {
		match := KMP(question, q)
		if match >= 0 {
			currentSimilarity := float64(len(q)) / float64(len(question))
			if currentSimilarity > similarity {
				similarity = currentSimilarity
				mostSimilarQuestion = q
			}
		}
	}
	if similarity >= 0.9 {
		return mostSimilarQuestion, true
	} else {
		return "", false
	}
}

func findAnswerKMP(question string, questions []string, answers []string) string {
	mostSimilarQuestion, ok := findMostSimilarKMP(question, questions)
	if ok {
		for i, q := range questions {
			if q == mostSimilarQuestion {
				return answers[i]
			}
		}
	} else {
		var mostSimilarQuestions []string
		var similarities []float64
		for _, q := range questions {
			similarity := float64(KMP(question, q)) / float64(len(q))
			if similarity >= 0.9 {
				mostSimilarQuestions = append(mostSimilarQuestions, q)
				similarities = append(similarities, similarity)
			}
		}
		if len(mostSimilarQuestions) > 0 {
			// sort questions by similarity
			sort.Slice(mostSimilarQuestions, func(i, j int) bool {
				return similarities[i] > similarities[j]
			})
			// select top 3 most similar questions
			if len(mostSimilarQuestions) > 3 {
				mostSimilarQuestions = mostSimilarQuestions[:3]
			}
			return "Pertanyaan Anda tidak ditemukan dalam database. Mungkin pertanyaan berikut mirip dengan yang Anda maksud: " + strings.Join(mostSimilarQuestions, ", ")
		}
	}
	return "Maaf, saya tidak mengerti pertanyaan Anda."
}

// Knuth-Morris-Pratt algorithm
func KMP(text, pattern string) int {
	if len(pattern) == 0 {
		return 0
	}
	prefix := prefix(pattern)
	i := 0
	j := 0
	for i < len(text) {
		if text[i] == pattern[j] {
			i++
			j++
			if j == len(pattern) {
				return i - j
			}
		} else {
			if j == 0 {
				i++
			} else {
				j = prefix[j-1]
			}
		}
	}
	return -1
}

func prefix(pattern string) []int {
	prefix := make([]int, len(pattern))
	i := 1
	j := 0
	for i < len(pattern) {
		if pattern[i] == pattern[j] {
			prefix[i] = j + 1
			i++
			j++
		} else {
			if j == 0 {
				prefix[i] = 0
				i++
			} else {
				j = prefix[j-1]
			}
		}
	}
	return prefix
}

// mencari pertanyaan paling mirip menggunakan algoritma BM
func findMostSimilarBM(question string, questions []string) (string, bool) {
	var mostSimilarQuestion string
	var similarity float64
	for _, q := range questions {
		match := BM(question, q)
		if match >= 0 {
			currentSimilarity := float64(len(q)) / float64(len(question))
			if currentSimilarity > similarity {
				similarity = currentSimilarity
				mostSimilarQuestion = q
			}
		}
	}
	if similarity >= 0.9 {
		return mostSimilarQuestion, true
	} else {
		return "", false
	}
}

func findAnswerBM(question string, questions []string, answers []string) string {
	// cari pertanyaan yang paling mirip dengan pertanyaan yang diberikan
	mostSimilarQuestion, found := findMostSimilarKMP(question, questions)
	if found {
		// cari jawaban yang cocok dengan pertanyaan yang ditemukan
		for i, q := range questions {
			if q == mostSimilarQuestion {
				return answers[i]
			}
		}
	}
	// jika tidak ditemukan pertanyaan yang cocok, cari satu pertanyaan yang paling mirip
	var topQuestion string
	var topSimilarity float64
	for _, q := range questions {
		similarity := float64(BM(question, q)) / float64(len(q))
		if similarity > topSimilarity {
			topQuestion = q
			topSimilarity = similarity
		}
	}
	if topSimilarity > 0.9 {
		return fmt.Sprintf("Maaf, pertanyaan Anda tidak ditemukan. Apakah yang Anda maksud adalah:\n%s", topQuestion)
	}
	return "Maaf, saya tidak dapat menemukan jawaban untuk pertanyaan Anda."
}

// Boyer-Moore algorithm
func BM(text, pattern string) int {
	if len(pattern) == 0 {
		return 0
	}
	last := last(pattern)
	i := len(pattern) - 1
	j := len(pattern) - 1
	for i < len(text) {
		if text[i] == pattern[j] {
			if j == 0 {
				return i
			} else {
				i--
				j--
			}
		} else {
			lo := last[int(text[i])]
			i = i + len(pattern) - min(j, 1+lo)
			j = len(pattern) - 1
		}
	}
	return -1
}

func last(pattern string) []int {
	last := make([]int, 256)
	for i := 0; i < 256; i++ {
		last[i] = -1
	}
	for i := 0; i < len(pattern); i++ {
		last[int(pattern[i])] = i
	}
	return last
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

var questions = []string{
	"Apa saya bisa dapat IP4?",
	"Kapan chatgpt ini dibuat?",
	"Siapa yang membuat chatgpt ini?",
}

var answers = []string{
	"Tidak ada yang tidak mungkin",
	"Chatbot ini dibuat pada tahun 2023.",
	"Vieri, Fajar, dan Copa membuat chatbot ini.",
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
	pattern := regexp.MustCompile(`tambah pertanyaan ([^ ]+) dengan jawaban ([^ ]+)`)
	return pattern.MatchString(question)
}

func parsingUpdateQuestion(question string) []string {
	if updateQuestionCheck(question) {
		pattern := regexp.MustCompile(`tambah pertanyaan ([^ ]+) dengan jawaban ([^ ]+)`)
		array := []string{pattern.FindStringSubmatch(question)[1], pattern.FindStringSubmatch(question)[2]}
		return array
	} else {
		array := []string{"", ""}
		return array
	}
}

func dateCheck(date string) bool {
	var pattern, _ = regexp.Compile(`(0[1-9]|[1-2][0-9]|3[0-1])\/(0[1-9]|1[0-2])\/\d{4}`)
	return pattern.MatchString(date)
}

func parsingDate(date string) string {
	if dateCheck(date) {
		pattern, _ := regexp.Compile(`(0[1-9]|[1-2][0-9]|3[0-1])\/(0[1-9]|1[0-2])\/\d{4}`)
		return pattern.FindStringSubmatch(date)[0]
	} else {
		return ""
	}
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
