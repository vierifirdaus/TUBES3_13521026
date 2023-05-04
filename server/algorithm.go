package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
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
	} else {
		// filter out non-similar questions
		similarityThreshold := 60
		var topQuestions []string
		for _, q := range similarQuestions {
			score := levenshteinDistance(processedQuestion, processText(q))
			maxScore := (len(processedQuestion) + len(q)) / 2
			similarity := 100 - ((score * 100) / maxScore)
			if similarity >= similarityThreshold {
				topQuestions = append(topQuestions, q)
			}
		}

		if len(topQuestions) > 0 {
			// sort top questions by similarity score
			sort.Slice(topQuestions, func(i, j int) bool {
				score1 := levenshteinDistance(processedQuestion, processText(topQuestions[i]))
				score2 := levenshteinDistance(processedQuestion, processText(topQuestions[j]))
				return score1 < score2
			})

			// construct response message
			var sb strings.Builder
			sb.WriteString("Pertanyaan Anda tidak ditemukan. Pertanyaan yang mirip:\n")
			maxQuestions := 3
			if len(topQuestions) < 3 {
				maxQuestions = len(topQuestions)
			}
			for i := 0; i < maxQuestions; i++ {
				sb.WriteString(fmt.Sprintf("- %s\n", topQuestions[i]))
			}
			return sb.String(), nil
		}
	}

	return "", fmt.Errorf("Maaf, saya tidak mengerti pertanyaan Anda.")
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
