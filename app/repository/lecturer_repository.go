package repository

import (
	"database/sql"
	"prestasi_backend/app/database"
	"prestasi_backend/app/model"
)

type LecturerRepository struct {
	DB *sql.DB
}

func NewLecturerRepository() *LecturerRepository {
	return &LecturerRepository{
		DB: database.PostgresDB,
	}
}

func (r *LecturerRepository) FindAll() ([]model.Lecturer, error) {
	rows, err := r.DB.Query(`
        SELECT id, user_id, lecturer_id, department, created_at
        FROM lecturers
        ORDER BY created_at DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.Lecturer

	for rows.Next() {
		var x model.Lecturer
		err := rows.Scan(&x.ID, &x.UserID, &x.LecturerID, &x.Department, &x.CreatedAt)
		if err != nil {
			return nil, err
		}
		list = append(list, x)
	}

	return list, nil
}

func (r *LecturerRepository) FindByID(id string) (*model.Lecturer, error) {
	row := r.DB.QueryRow(`
        SELECT id, user_id, lecturer_id, department, created_at
        FROM lecturers
        WHERE id = $1`, id)

	var x model.Lecturer
	err := row.Scan(&x.ID, &x.UserID, &x.LecturerID, &x.Department, &x.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &x, nil
}
