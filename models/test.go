package models

// GenericTestOverview overview data of a generic test
type GenericTestOverview struct {
	ID   string
	Name string
	Type string
}

type GenericTestData struct {
	ID             string
	Name           string
	Title          string
	Type           string
	Description    string
	DisplayAnswers bool
	FinalRemark    string
	Questions      []GenericQuestion
}

type SubmittedAnswers struct {
	QuestionNumber int32
	Answer         string
	TimeToAnswer   float32
}
