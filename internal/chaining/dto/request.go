package dto

type UserIndicatorInput struct {
	ChainingID     string `json:"chaining_id"`
	QuestionID     string `json:"question_id"`
	QuestionStatus bool   `json:"question_status"`
}
