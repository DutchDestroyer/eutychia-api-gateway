package database

// GenericQuestionDAO model of the generic questions
type GenericQuestionDAO struct {
	Question     string
	QuestionType string
	Answers      []string
}
