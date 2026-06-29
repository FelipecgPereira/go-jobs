package models

import (
	"database/sql"
	"errors"
	"time"

	"com.github/FelipecgPereira/go-jobs/db"
	"com.github/FelipecgPereira/go-jobs/models/enums"
)

type B2b struct {
	Id         int64
	CustomerId int64 `binding:"required`
	UserId     int64
	ProjectId  int64
	Status     enums.Status
	CreateAt   time.Time
	UpdateAt   time.Time
}

func (input *B2b) Save() error {
	if !input.Status.IsValid() {
		input.Status = enums.StatusPending
	}

	query := `
	INSERT INTO b2b (customer_id, user_id, project_id, status, create_at)
	VALUES (?, ?, ?, ?,?)
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if input.CreateAt.IsZero() {
		input.CreateAt = time.Now()
	}

	_, err = stmt.Exec(input.CustomerId, input.UserId, input.ProjectId, string(input.Status), input.CreateAt)
	if err != nil {
		return err
	}

	return nil
}

func GetB2bByID(input int64) (*B2b, error) {
	query := `SELECT id, customer_id, user_id, project_id, status, create_at FROM b2b WHERE id = ?`
	row := db.DB.QueryRow(query, input)
	var statusString string

	var b2b B2b
	err := row.Scan(&b2b.Id, &b2b.CustomerId, &b2b.UserId, &b2b.ProjectId, &statusString, &b2b.CreateAt)
	if err != nil {
		return nil, err
	}

	b2b.Status = enums.Status(statusString)
	return &b2b, nil
}

func (input *B2b) Update() error {
	if input.Id == 0 {
		return errors.New("missing id")
	}

	query := `
	UPDATE b2b
	SET customer_id = ?, project_id = ?, status = ?, update_at = ?
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	input.UpdateAt = time.Now()
	_, err = stmt.Exec(input.CustomerId, input.ProjectId, string(input.Status), input.UpdateAt, input.Id)
	if err != nil {
		return err
	}
	return nil
}

func SumPaymentsByStatusAndDate(status enums.Status, startDate, endDate time.Time, userId int64) (float64, error) {
	query := `
        SELECT COALESCE(SUM(p.price), 0)
        FROM b2b b
        JOIN projects p ON p.id = b.project_id
        WHERE b.status = ?
          AND p.start_date >= ?
          AND p.end_date <= ?
					AND b.user_id = ?
    `

	var total sql.NullFloat64
	err := db.DB.QueryRow(query, string(status), startDate, endDate, userId).Scan(&total)
	if err != nil {
		return 0, err
	}

	return total.Float64, nil
}
