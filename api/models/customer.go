package models

import (
	"time"

	"com.github/FelipecgPereira/go-jobs/db"
)

type Customer struct {
	Id       int64
	Name     string `binding:"required`
	Email    string `binding:"required`
	Phone    string
	UserID   int64
	createAt time.Time
	updateAt time.Time
}

func (input *Customer) Save() error {
	query := `
		INSERT INTO customers (name, email, phone, user_id,create_at) 
		VALUES(?,?,?,?,?)
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	input.createAt = time.Now()

	result, err := stmt.Exec(input.Name, input.Email, input.Phone, input.UserID, input.createAt)

	if err != nil {
		return err
	}

	_, err = result.LastInsertId()

	if err != nil {
		return err
	}

	return nil
}

func GetCustomer(input int64) ([]Customer, error) {
	query := `
		SELECT id, name, email, phone, user_id From customers where user_id = ?
	`
	rows, err := db.DB.Query(query, input)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var customers []Customer

	for rows.Next() {
		var customer Customer
		err := rows.Scan(&customer.Id, &customer.Name, &customer.Email, &customer.Phone, &customer.UserID)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	return customers, nil

}

func (input *Customer) Update() error {
	query := `
	UPDATE customers 
	SET name= ?,email=? ,phone=? ,update_at=? 
	WHERE id =?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	input.updateAt = time.Now()
	_, err = stmt.Exec(input.Name, input.Email, input.Phone, input.updateAt, input.Id)
	if err != nil {
		return err
	}
	return nil
}

func GetCustomerById(input int64) (*Customer, error) {
	query := `
		SELECT id, name, email, phone,user_id From customers where id = ?
	`
	row := db.DB.QueryRow(query, input)
	var customer Customer
	err := row.Scan(&customer.Id, &customer.Name, &customer.Email, &customer.Phone, &customer.UserID)

	if err != nil {
		return nil, err
	}

	return &customer, nil
}
