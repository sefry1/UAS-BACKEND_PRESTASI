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

// FindAll - Get all lecturers using stored procedure
func (r *LecturerRepository) FindAll() ([]model.Lecturer, error) {
	rows, err := r.DB.Query(`SELECT * FROM sp_get_all_lecturers()`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.Lecturer

	for rows.Next() {
		var l model.Lecturer
		var lecturerName sql.NullString
		err := rows.Scan(&l.ID, &l.UserID, &l.LecturerID, &lecturerName, &l.Department)
		if err != nil {
			return nil, err
		}
		list = append(list, l)
	}

	return list, nil
}

// FindByID - Get lecturer by ID
func (r *LecturerRepository) FindByID(id string) (*model.Lecturer, error) {
	row := r.DB.QueryRow(`SELECT * FROM lecturers WHERE id = $1`, id)

	var l model.Lecturer
	err := row.Scan(&l.ID, &l.UserID, &l.LecturerID, &l.Department, &l.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &l, nil
}

// FindByUserID - Get lecturer by user_id using stored procedure
func (r *LecturerRepository) FindByUserID(userID string) (*model.Lecturer, error) {
	row := r.DB.QueryRow(`SELECT * FROM sp_get_lecturer_by_user_id($1)`, userID)

	var l model.Lecturer
	err := row.Scan(&l.ID, &l.UserID, &l.LecturerID, &l.Department, &l.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &l, nil
}

// GetAdvisees - Get all students advised by this lecturer using stored procedure
func (r *LecturerRepository) GetAdvisees(lecturerID string) ([]model.Student, error) {
	rows, err := r.DB.Query(`SELECT * FROM sp_get_lecturer_advisees($1)`, lecturerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.Student

	for rows.Next() {
		var s model.Student
		err := rows.Scan(&s.ID, &s.UserID, &s.StudentID, &s.ProgramStudy, &s.AcademicYear, &s.AdvisorID, &s.CreatedAt)
		if err != nil {
			return nil, err
		}
		list = append(list, s)
	}

	return list, nil
}
