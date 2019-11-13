package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Question struct {
	Id        int    `db:"id" json:"id"`
	Body      string `db:"question" json:"body"`
	CreatedAt string `db:"created_at" json:"created_at"`
}

type QAndA struct {
	Id        int     `db:"id" json:"id"`
	Question  string  `db:"question" json:"question"`
	Answer    *string `db:"answer" json:"answer"`
	CreatedAt string  `db:"created_at" json:"created_at"`
}

type Conn struct {
	conn *sqlx.DB
}

var ErrContentNotFound = errors.New("not found")

func NewDB() (*Conn, error) {
	var err error
	db, err := sqlx.Open(
		"mysql",
		"aratasato:hoge@tcp(mysql:3306)/geing",
	)

	return &Conn{db}, errors.Wrap(err, "failed to connect db")
}

// 質問を追加
func (db *Conn) SaveQuestion(body string) error {
	fmt.Println("Save question: " + body)
	tx, err := db.conn.Beginx()
	if err != nil {
		return errors.Wrap(err, "failed to connect db")
	}
	_, err = tx.Exec("INSERT INTO qandas (question) VALUES (?)", body)
	if err != nil {
		return errors.Wrap(err, "failed to add question")
	}
	_ = tx.Commit()
	return nil
}

// 質問回答セットを1件取得
func (db *Conn) GetQA(id int) (*QAndA, error) {
	qa := QAndA{}
	err := db.conn.Get(&qa, "SELECT * FROM qandas WHERE id = ?", id)
	if err == sql.ErrNoRows {
		return nil, ErrContentNotFound
	}

	if err != nil {
		return nil, errors.Wrap(err, "failed to get qa")
	}
	return &qa, err
}

// 質問を20件取得
func (db *Conn) GetQuestions(page int) ([]Question, error) {
	var questions []Question
	err := db.conn.Select(
		&questions,
		`
			SELECT id, question, created_at
			FROM qandas WHERE id > ? * 10
			AND answer IS NOT NULL
			ORDER BY id DESC
			LIMIT 20
		`,
		page,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get question")
	}
	return questions, nil
}

// 回答を追加
func (db *Conn) SaveAnswer(body string, id int) error {
	fmt.Println("Save answer: " + body)
	tx, err := db.conn.Beginx()
	if err != nil {
		return errors.Wrap(err, "failed to connect db")
	}
	_, err = tx.Exec("UPDATE `qandas` t SET t.`answer` = ? WHERE t.`id` = ?", body, id)
	if err != nil {
		return errors.Wrap(err, "failed to add answer")
	}
	_ = tx.Commit()
	return nil
}
