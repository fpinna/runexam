package types

type Exam struct {
	TestMetadata TestMetadata `json:"TestMetadata"`
	Questions    []Question   `json:"Questions"`
}

type TestMetadata struct {
	TestName          string   `json:"TestName"`
	PassingPercentage float64  `json:"PassingPercentage"`
	ExamDomains       []string `json:"ExamDomains"`
	ExamDescription   string   `json:"ExamDescription"`
}

type Question struct {
	Title         string            `json:"Title"`
	Domain        string            `json:"Domain"`
	Question      string            `json:"Question"`
	Options       map[string]string `json:"Options"`
	Type          string            `json:"Type"`
	CorrectAnswer []string          `json:"CorrectAnswer"`
	Explanation   string            `json:"Explanation"`
}

type AnswerResult struct {
	Question      string
	UserAnswer    string
	CorrectAnswer string
	Explanation   string
	IsCorrect     bool
}

type ResultPage struct {
	Total          int
	Correct        int
	Score          float64
	Passed         bool
	Results        []AnswerResult
	TestName       string
	PassingPercent float64
}
