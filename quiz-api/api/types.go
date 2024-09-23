package api

type Option struct {
	ID        string
	Description string
}

type Question struct {
	ID string
	Description string
	Options     []Option
}

type Answer struct {
	QuestionID string `json:"question_id"`
	OptionID   string `json:"option_id"`
}

const UnauthorizedErrorMessage = "user unauthorized to perform method call"
