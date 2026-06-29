package models

import (
	"time"

	"com.github/FelipecgPereira/go-jobs/db"
)

type Project struct {
	Id        int64
	Name      string
	Price     float64
	StartDate time.Time
	EndDate   time.Time
	UserId    int64
	createAt  time.Time
	updateAt  time.Time
}

func (input *Project) Save() (int64, error) {
	query := `
		INSERT INTO projects (name, price, start_date, end_date, user_id,create_at)
		VALUES (?,?,?,?,?,?)
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	input.createAt = time.Now()
	result, err := stmt.Exec(input.Name, input.Price, input.StartDate, input.EndDate, input.UserId, input.createAt)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return id, nil
}

func GetProjects(input int64) ([]Project, error) {
	query := `
		SELECT id, name, price, start_date, end_date, user_id From projects where user_id = ?
	`
	rows, err := db.DB.Query(query, input)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var projects []Project

	for rows.Next() {
		var project Project
		err := rows.Scan(&project.Id, &project.Name, &project.Price,
			&project.StartDate, &project.EndDate, &project.UserId)
		if err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}
	return projects, nil
}

func GetProjectById(input int64) (*Project, error) {
	query := `
		SELECT id, name, price, start_date, end_date, user_id From projects where id = ?
	`
	row := db.DB.QueryRow(query, input)
	var project Project
	err := row.Scan(&project.Id, &project.Name, &project.Price,
		&project.StartDate, &project.EndDate, &project.UserId)
	if err != nil {
		return nil, err
	}

	return &project, nil
}

func (input *Project) Update() error {
	query := `
	UPDATE projects 
	SET name= ?,price=?,start_date=? ,end_date=?,update_at=? 
	WHERE id =?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	input.updateAt = time.Now()
	_, err = stmt.Exec(input.Name, input.Price, input.StartDate, input.EndDate, input.updateAt, input.Id)
	if err != nil {
		return err
	}
	return nil
}
