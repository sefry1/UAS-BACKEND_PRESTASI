package repository

import (
	"database/sql"
	"prestasi_backend/app/database"
	"prestasi_backend/app/model"
)

type StudentRepository struct {
	DB *sql.DB
}

func NewStudentRepository() *StudentRepository {
	return &StudentRepository{DB: database.PostgresDB}
}

func (r *StudentRepository) FindAll() ([]model.Student, error) {
	rows, err := r.DB.Query(`
        SELECT id, user_id, student_id, program_study, academic_year, advisor_id
        FROM students`)
	if err != nil {
		return nil, err
	}

	var list []model.Student
	for rows.Next() {
		var x model.Student
		rows.Scan(&x.ID, &x.UserID, &x.StudentID, &x.ProgramStudy, &x.AcademicYear, &x.AdvisorID)
		list = append(list, x)
	}
	return list, nil
}

func (r *StudentRepository) FindByID(id string) (*model.Student, error) {
	row := r.DB.QueryRow(`
        SELECT id, user_id, student_id, program_study, academic_year, advisor_id
        FROM students WHERE id=$1`, id)

	var x model.Student
	err := row.Scan(&x.ID, &x.UserID, &x.StudentID, &x.ProgramStudy, &x.AcademicYear, &x.AdvisorID)
	if err != nil {
		return nil, err
	}
	return &x, nil
}

func (r *StudentRepository) FindByAdvisor(advisor string) ([]model.Student, error) {
	rows, err := r.DB.Query(`
        SELECT id, user_id, student_id, program_study, academic_year, advisor_id
        FROM students WHERE advisor_id=$1`, advisor)
	if err != nil {
		return nil, err
	}

	var list []model.Student
	for rows.Next() {
		var x model.Student
		rows.Scan(&x.ID, &x.UserID, &x.StudentID, &x.ProgramStudy, &x.AcademicYear, &x.AdvisorID)
		list = append(list, x)
	}
	return list, nil
}

func (r *StudentRepository) UpdateAdvisor(id, advisor string) error {
	_, err := r.DB.Exec(`
        UPDATE students SET advisor_id=$1 WHERE id=$2`,
		advisor, id)
	return err
}
