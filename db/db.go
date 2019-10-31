package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Questions struct {
	Id        int    `db:"id"`
	Question  string `db:"question"`
	CreatedAt string `db:"created_at"`
}

type QAndA struct {
	Id        int     `db:"id"`
	Question  string  `db:"question"`
	Answered  bool    `db:"answered"`
	Answer    *string `db:"answer"`
	CreatedAt string  `db:"created_at"`
}

type Conn struct {
	conn *sqlx.DB
}

var ErrContentNotFound = errors.New("not found")

func NewDB() (*Conn, error) {
	var err error
	db, err := sqlx.Open(
		"mysql",
		"aratasato:hoge@tcp(127.0.0.1:3306)/geing",
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
func (db *Conn) GetQuestions(page int) (*[]Questions, error) {
	var questions []Questions
	err := db.conn.Select(&questions, "SELECT id, question, created_at FROM qandas WHERE id > ? * 10 LIMIT 20", page)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get question")
	}
	return &questions, nil
}
