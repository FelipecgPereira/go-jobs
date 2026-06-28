package models

import (
	"time"

	"com.github/FelipecgPereira/go-jobs/db"
	"com.github/FelipecgPereira/go-jobs/utils"
)

type User struct {
	Id       int64
	Name     string `binding:"required`
	Email    string `binding:"required`
	Password string `binding:"required`
	createAt time.Time
	updateAt time.Time
}

func (input *User) Save() error {
	query := `
	INSERT INTO users (email, password, name, create_at)
	VALUES (?,?,?,?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	input.userCreateAt()

	hashedPassword, err := utils.HashPassword(input.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(input.Email, hashedPassword, input.Name, input.createAt)

	if err != nil {
		return err
	}

	_, err = result.LastInsertId()

	if err != nil {
		return err
	}

	return nil

}

func (input *User) userCreateAt() {
	input.createAt = time.Now()
}
