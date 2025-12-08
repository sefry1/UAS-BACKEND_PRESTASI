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

// FindAll - Get all students using stored procedure
func (r *StudentRepository) FindAll() ([]model.Student, error) {
	rows, err := r.DB.Query(`SELECT * FROM sp_get_all_students()`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.Student
	for rows.Next() {
		var s model.Student
		var studentName, advisorName sql.NullString
		err := rows.Scan(&s.ID, &s.UserID, &s.StudentID, &studentName, &s.ProgramStudy, &s.AcademicYear, &s.AdvisorID, &advisorName)
		if err != nil {
			return nil, err
		}
		list = append(list, s)
	}
	return list, nil
}

// FindByID - Get student by ID
func (r *StudentRepository) FindByID(id string) (*model.Student, error) {
	row := r.DB.QueryRow(`SELECT * FROM students WHERE id=$1`, id)

	var s model.Student
	err := row.Scan(&s.ID, &s.UserID, &s.StudentID, &s.ProgramStudy, &s.AcademicYear, &s.AdvisorID, &s.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

// FindByUserID - Get student by user_id (IMPORTANT for achievement creation)
func (r *StudentRepository) FindByUserID(userID string) (*model.Student, error) {
	row := r.DB.QueryRow(`SELECT * FROM sp_get_student_by_user_id($1)`, userID)

	var s model.Student
	var advisorName sql.NullString // advisor_name from stored procedure
	err := row.Scan(&s.ID, &s.UserID, &s.StudentID, &s.ProgramStudy, &s.AcademicYear, &s.AdvisorID, &advisorName, &s.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

// FindByAdvisor - Get students by advisor ID
func (r *StudentRepository) FindByAdvisor(advisorID string) ([]model.Student, error) {
	rows, err := r.DB.Query(`
        SELECT id, user_id, student_id, program_study, academic_year, advisor_id, created_at
        FROM students WHERE advisor_id=$1`, advisorID)
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

// UpdateAdvisor - Set student's advisor using stored procedure
func (r *StudentRepository) UpdateAdvisor(studentID, advisorID string) error {
	_, err := r.DB.Exec(`SELECT sp_set_student_advisor($1, $2)`, studentID, advisorID)
	return err
}
