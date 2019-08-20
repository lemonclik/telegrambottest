package db

import (
	"database/sql"
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DataBase struct {
	DB *sql.DB
}

// InitDB creates tables USERs or SALEs if tables not exists
func InitDB(db *sql.DB) (*DataBase, error) {

	d := DataBase{
		DB: db,
	}
	_, err := db.Exec(CREATE_USERS_IF_NOT_EXISTS)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(CREATE_SALES_IF_NOT_EXISTS)
	if err != nil {
		return nil, err
	}

	return &d, nil
}

// PutUser adds user in database
func (d *DataBase) PutUser(ChatId int64) error {

	_, err := d.DB.Exec("INSERT INTO USERS(id, chat_id, lang)"+
		"VALUES ($1,$2,$3)", int(ChatId), ChatId, "en")
	if err != nil {
		return err
	}

	return nil
}

// GetLanguage returns language for user by UserId
func (d *DataBase) GetLanguage(Userid int) string {

	rows := d.DB.QueryRow("SELECT lang FROM USERS WHERE id = $1 limit 1", Userid)

	var lang string
	if rows == nil {
		fmt.Println("Empty!")
	}

	err := rows.Scan(&lang)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return lang
}

// SetLanguage is setting language for user by UserId
func (d *DataBase) SetLanguage(UserId int, lang string) error {

	_, err := d.DB.Exec("UPDATE USERS SET lang = $1 where id = $2", lang, UserId)
	if err != nil {
		return err
	}

	return nil
}

// GetSales returns all sales for user by UserId
func (d *DataBase) GetSales() {

}

// UpdateSales updates (insert new) sales for user by UserId
func (d *DataBase) UpdateSales() {

}
