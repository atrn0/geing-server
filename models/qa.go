package models

type AddQuestionsResponse struct {
	Id        int    `json:"id"`
	Body      string `json:"body"`
	CreatedAt string `json:"created_at"`
}

type GetQAsResponse struct {
	Id        int    `json:"id"`
	Question  string `json:"question"`
	Answer    string `json:"answer"`
	CreatedAt string `json:"created_at"`
}

type QA struct {
	Id        int    `db:"id"`
	Question  string `db:"question"`
	Answered  bool   `db:"answered"`
	Answer    string `db:"answer"`
	CreatedAt string `db:"created_at"`
}
