package server

import (
	"encoding/json"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"runexam/types"
	"runexam/utils"
	"strconv"
	"time"
)

var (
	questionsData    types.Exam
	lastResult       types.ResultPage
	currentQuestions []types.Question
)

func StartServer(jsonPath, listen string, port int) {
	data, err := os.ReadFile(jsonPath)
	if err != nil {
		log.Fatalf("Error reading JSON: %v", err)
	}

	if err := json.Unmarshal(data, &questionsData); err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", showExam)
	mux.HandleFunc("/submit", handleSubmit)
	mux.HandleFunc("/pdf", handlePDF)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	addr := listen + ":" + strconv.Itoa(port)
	log.Printf("Server listening on http://%s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}

func getFuncMap() template.FuncMap {
	return template.FuncMap{
		"add1": func(i int) int { return i + 1 },
	}
}

// Função para copiar e embaralhar um slice de questões
func shuffleQuestions(questions []types.Question) []types.Question {
	shuffled := make([]types.Question, len(questions))
	copy(shuffled, questions)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(shuffled), func(i, j int) { shuffled[i], shuffled[j] = shuffled[j], shuffled[i] })
	return shuffled
}

func showExam(w http.ResponseWriter, r *http.Request) {
	currentQuestions = shuffleQuestions(questionsData.Questions)
	examCopy := questionsData
	examCopy.Questions = currentQuestions

	tmpl := template.Must(template.New("").Funcs(getFuncMap()).ParseFiles("templates/exam.html"))
	tmpl.ExecuteTemplate(w, "exam.html", examCopy)
}

func handleSubmit(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error processing answers", http.StatusBadRequest)
		return
	}

	correct := 0
	var results []types.AnswerResult

	for i, q := range currentQuestions {
		ans := r.Form.Get("q" + strconv.Itoa(i))
		correctAns := q.CorrectAnswer[0]

		results = append(results, types.AnswerResult{
			Question:      q.Question,
			UserAnswer:    ans,
			CorrectAnswer: correctAns,
			Explanation:   q.Explanation,
			IsCorrect:     ans == correctAns,
		})

		if ans == correctAns {
			correct++
		}
	}

	totalQuestions := len(currentQuestions)
	score := float64(correct) / float64(totalQuestions) * 100
	passingPercent := questionsData.TestMetadata.PassingPercentage
	passed := score >= passingPercent

	lastResult = types.ResultPage{
		Total:          totalQuestions,
		Correct:        correct,
		Score:          score,
		Passed:         passed,
		Results:        results,
		TestName:       questionsData.TestMetadata.TestName,
		PassingPercent: passingPercent,
	}

	tmpl := template.Must(template.New("").Funcs(getFuncMap()).ParseFiles("templates/result.html"))
	tmpl.ExecuteTemplate(w, "result.html", lastResult)
}

func handlePDF(w http.ResponseWriter, r *http.Request) {
	pdfBytes := utils.GeneratePDF(lastResult)
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=\"resultado.pdf\"")
	w.Write(pdfBytes)
}
