package repository

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	. "questionBoxWithGo/models"
	. "questionBoxWithGo/utils"
)

var (
	db *sqlx.DB
)

func init() {
	var err error
	db, err = sqlx.Open(
		"mysql",
		"aratasato:hoge@tcp(127.0.0.1:3306)/geing",
	)
	HandleError(err)
}

// 質問を追加
func CreateQuestion(body string) error {
	fmt.Println(body)
	tx := db.MustBegin()
	tx.MustExec("INSERT INTO qandas (question) VALUES (?)", body)
	return tx.Commit()
}

// 質問回答セットを1件取得
func GetQA(id int) (QA, error) {
	qa := QA{}
	err := db.Get(&qa, "SELECT * FROM qandas WHERE id = ?", id)
	return qa, err
}

//func GetQuestions(page)
