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

func CreateQuestion(body string) error {
	fmt.Println(body)
	tx := db.MustBegin()
	tx.MustExec("INSERT INTO qandas (question) VALUES (?)", body)
	return tx.Commit()
}

func GetQA(id string) (QA, error) {
	fmt.Println("get question: ", id)
	qa := QA{}
	err := db.Get(&qa, "SELECT * FROM qandas WHERE id = ?", id)
	return qa, err
}
