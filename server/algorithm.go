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
		for i, q := range questions {
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

func findAnswerBM(question string, questions []string, answers []string) (string, bool) {
	// cari pertanyaan yang paling mirip dengan pertanyaan yang diberikan
	mostSimilarQuestion, found := findMostSimilarKMP(question, questions)
	if !found {
		// jika tidak ditemukan pertanyaan yang cocok, cari tiga pertanyaan yang paling mirip
		var topQuestions []string
		var topSimilarities []float64
		for i, q := range questions {
			similarity := float64(BM(question, q)) / float64(len(q))
			if similarity > 0.9 {
				if len(topQuestions) < 3 {
					topQuestions = append(topQuestions, q)
					topSimilarities = append(topSimilarities, similarity)
				} else {
					// ganti pertanyaan dengan kesamaan terendah di antara tiga pertanyaan
					minIndex := 0
					for j := 1; j < 3; j++ {
						if topSimilarities[j] < topSimilarities[minIndex] {
							minIndex = j
						}
					}
					if similarity > topSimilarities[minIndex] {
						topQuestions[minIndex] = q
						topSimilarities[minIndex] = similarity
					}
				}
			}
		}
		if len(topQuestions) > 0 {
			var options string
			for i, q := range topQuestions {
				options += fmt.Sprintf("%d) %s\n", i+1, q)
			}
			return fmt.Sprintf("Pertanyaan tidak ditemukan. Apakah yang Anda maksud adalah:\n%s", options), false
		}
		return "Maaf, saya tidak dapat menemukan jawaban untuk pertanyaan Anda.", false
	}
	// cari jawaban yang cocok dengan pertanyaan yang ditemukan
	for i, q := range questions {
		if q == mostSimilarQuestion {
			return answers[i], true
		}
	}
	return "", false
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