package types

type Exam struct {
	TestMetadata TestMetadata `json:"TestMetadata"`
	Questions    []Question   `json:"Questions"`
}

type TestMetadata struct {
	TestName          string   `json:"TestName"`
	TestDescription   string   `json:"TestDescription"`
	PassingPercentage float64  `json:"PassingPercentage"`
	TestVersion       string   `json:"TestVersion"`
	TestAuthor        string   `json:"TestAuthor"`
	TestDate          string   `json:"TestDate"`
	TestDuration      int      `json:"TestDuration"`
	ExamDomains       []string `json:"ExamDomains"`
	ExamDescription   string   `json:"ExamDescription"`
	TotalQuestions    int      `json:"TotalQuestions"`
}

type Question struct {
	Title         string            `json:"Title"`
	Domain        string            `json:"Domain"`
	Question      string            `json:"Question"`
	Options       map[string]string `json:"Options,omitempty"`
	Type          string            `json:"Type"`                    // "Single", "Multiple", "True", "False"
	CorrectAnswer []string          `json:"CorrectAnswer,omitempty"` // For Single/Multiple only
	Explanation   string            `json:"Explanation"`
}

type AnswerResult struct {
	Question      string
	UserAnswer    []string
	CorrectAnswer []string // For Single/Multiple
	CorrectTF     string   // For True/False: copy q.Type
	Explanation   string
	IsCorrect     bool
	Type          string
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
