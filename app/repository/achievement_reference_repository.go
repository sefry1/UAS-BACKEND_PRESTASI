package repository

import (
	"database/sql"
	"prestasi_backend/app/database"
	"prestasi_backend/app/model"
)

type AchievementReferenceRepository struct {
	DB *sql.DB
}

func NewAchievementReferenceRepository() *AchievementReferenceRepository {
	return &AchievementReferenceRepository{
		DB: database.PostgresDB,
	}
}

//
// ========================================================
// FIND ALL (dipakai REPORT SERVICE / ADMIN DASHBOARD)
// ========================================================
func (r *AchievementReferenceRepository) FindAll() ([]model.AchievementReference, error) {

	rows, err := r.DB.Query(`SELECT * FROM sp_get_all_achievements()`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.AchievementReference

	for rows.Next() {
		var a model.AchievementReference
		var studentName string
		err := rows.Scan(
			&a.ID, &a.StudentID, &studentName, &a.MongoAchievementID, &a.Status,
			&a.SubmittedAt, &a.VerifiedAt, &a.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, a)
	}

	return list, nil
}

//
// ========================================================
// FIND BY ID
// ========================================================
func (r *AchievementReferenceRepository) FindByID(id string) (*model.AchievementReference, error) {

	row := r.DB.QueryRow(`SELECT * FROM sp_get_achievement_by_id($1)`, id)

	var a model.AchievementReference
	var verifiedByName sql.NullString
	err := row.Scan(
		&a.ID, &a.StudentID, &a.MongoAchievementID, &a.Status,
		&a.SubmittedAt, &a.VerifiedAt, &a.VerifiedBy,
		&verifiedByName, &a.RejectionNote, &a.CreatedAt, &a.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &a, nil
}

//
// ========================================================
// FIND BY USER ID (dipakai Achievement List mahasiswa login)
// ========================================================
func (r *AchievementReferenceRepository) FindByUserID(userID string) ([]model.AchievementReference, error) {

	// First, get student_id from user_id
	var studentID string
	err := r.DB.QueryRow(`SELECT id FROM students WHERE user_id = $1`, userID).Scan(&studentID)
	if err != nil {
		return nil, err
	}

	rows, err := r.DB.Query(`SELECT * FROM sp_get_achievements_by_student($1)`, studentID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.AchievementReference

	for rows.Next() {
		var a model.AchievementReference
		err := rows.Scan(
			&a.ID, &a.MongoAchievementID, &a.Status,
			&a.SubmittedAt, &a.VerifiedAt, &a.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		a.StudentID = studentID
		list = append(list, a)
	}

	return list, nil
}

//
// ========================================================
// FIND BY STUDENT ID (dipakai Student Detail & Report Student)
// ========================================================
func (r *AchievementReferenceRepository) FindByStudentID(studentID string) ([]model.AchievementReference, error) {

	rows, err := r.DB.Query(`SELECT * FROM sp_get_achievements_by_student($1)`, studentID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.AchievementReference

	for rows.Next() {
		var a model.AchievementReference
		err := rows.Scan(
			&a.ID, &a.MongoAchievementID, &a.Status,
			&a.SubmittedAt, &a.VerifiedAt, &a.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		a.StudentID = studentID
		list = append(list, a)
	}

	return list, nil
}

//
// ========================================================
// FIND BY ADVISOR (untuk dosen wali)
// ========================================================
func (r *AchievementReferenceRepository) FindByAdvisor(lecturerID string) ([]model.AchievementReference, error) {

	rows, err := r.DB.Query(`SELECT * FROM sp_get_achievements_by_advisor($1)`, lecturerID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.AchievementReference

	for rows.Next() {
		var a model.AchievementReference
		var studentName string
		err := rows.Scan(
			&a.ID, &a.StudentID, &studentName, &a.MongoAchievementID,
			&a.Status, &a.SubmittedAt, &a.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, a)
	}

	return list, nil
}

//
// ========================================================
// CREATE (dari mahasiswa create achievement)
// ========================================================
func (r *AchievementReferenceRepository) Create(studentID, mongoAchievementID string) (string, error) {

	var achievementID string
	err := r.DB.QueryRow(`SELECT sp_create_achievement_reference($1, $2)`,
		studentID, mongoAchievementID).Scan(&achievementID)

	if err != nil {
		return "", err
	}

	return achievementID, nil
}

//
// ========================================================
// SUBMIT (mahasiswa submit for verification)
// ========================================================
func (r *AchievementReferenceRepository) Submit(id string) error {

	var success bool
	err := r.DB.QueryRow(`SELECT sp_submit_achievement($1)`, id).Scan(&success)

	if err != nil {
		return err
	}

	return nil
}

//
// ========================================================
// VERIFY (dosen wali verify)
// ========================================================
func (r *AchievementReferenceRepository) Verify(id, verifier string) error {

	var success bool
	err := r.DB.QueryRow(`SELECT sp_verify_achievement($1, $2)`, id, verifier).Scan(&success)

	if err != nil {
		return err
	}

	return nil
}

//
// ========================================================
// REJECT (dosen wali reject)
// ========================================================
func (r *AchievementReferenceRepository) Reject(id, verifier, note string) error {

	var success bool
	err := r.DB.QueryRow(`SELECT sp_reject_achievement($1, $2, $3)`,
		id, verifier, note).Scan(&success)

	if err != nil {
		return err
	}

	return nil
}

//
// ========================================================
// DELETE (only draft)
// ========================================================
func (r *AchievementReferenceRepository) Delete(id string) error {

	var success bool
	err := r.DB.QueryRow(`SELECT sp_delete_achievement($1)`, id).Scan(&success)

	if err != nil {
		return err
	}

	return nil
}
