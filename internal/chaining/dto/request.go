package dto

type UserIndicatorInput struct {
	QuestionID     string `json:"question_id"`
	QuestionStatus bool   `json:"question_status"`
}
