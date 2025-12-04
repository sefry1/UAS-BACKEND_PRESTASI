package repository

import (
	"database/sql"
	"prestasi_backend/app/database"
	"prestasi_backend/app/model"
	"time"
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

	rows, err := r.DB.Query(`
        SELECT id, student_id, mongo_achievement_id, status,
               submitted_at, verified_at, verified_by,
               rejection_note, created_at, updated_at
        FROM achievement_references
        ORDER BY created_at DESC
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.AchievementReference

	for rows.Next() {
		var a model.AchievementReference
		err := rows.Scan(
			&a.ID, &a.StudentID, &a.MongoAchievementID, &a.Status,
			&a.SubmittedAt, &a.VerifiedAt, &a.VerifiedBy,
			&a.RejectionNote, &a.CreatedAt, &a.UpdatedAt,
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

	row := r.DB.QueryRow(`
        SELECT id, student_id, mongo_achievement_id, status,
               submitted_at, verified_at, verified_by,
               rejection_note, created_at, updated_at
        FROM achievement_references
        WHERE id = $1
    `, id)

	var a model.AchievementReference
	err := row.Scan(
		&a.ID, &a.StudentID, &a.MongoAchievementID, &a.Status,
		&a.SubmittedAt, &a.VerifiedAt, &a.VerifiedBy,
		&a.RejectionNote, &a.CreatedAt, &a.UpdatedAt,
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

	rows, err := r.DB.Query(`
        SELECT ar.id, ar.student_id, ar.mongo_achievement_id, ar.status,
               ar.submitted_at, ar.verified_at, ar.verified_by,
               ar.rejection_note, ar.created_at, ar.updated_at
        FROM achievement_references ar
        JOIN students s ON ar.student_id = s.id
        WHERE s.user_id = $1
        ORDER BY ar.created_at DESC
    `, userID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.AchievementReference

	for rows.Next() {
		var a model.AchievementReference
		err := rows.Scan(
			&a.ID, &a.StudentID, &a.MongoAchievementID, &a.Status,
			&a.SubmittedAt, &a.VerifiedAt, &a.VerifiedBy,
			&a.RejectionNote, &a.CreatedAt, &a.UpdatedAt,
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
// FIND BY STUDENT ID (dipakai Student Detail & Report Student)
// ========================================================
func (r *AchievementReferenceRepository) FindByStudentID(studentID string) ([]model.AchievementReference, error) {

	rows, err := r.DB.Query(`
        SELECT id, student_id, mongo_achievement_id, status,
               submitted_at, verified_at, verified_by,
               rejection_note, created_at, updated_at
        FROM achievement_references
        WHERE student_id = $1
        ORDER BY created_at DESC
    `, studentID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.AchievementReference

	for rows.Next() {
		var a model.AchievementReference
		err := rows.Scan(
			&a.ID, &a.StudentID, &a.MongoAchievementID, &a.Status,
			&a.SubmittedAt, &a.VerifiedAt, &a.VerifiedBy,
			&a.RejectionNote, &a.CreatedAt, &a.UpdatedAt,
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
// UPDATE STATUS (submit)
// ========================================================
func (r *AchievementReferenceRepository) UpdateStatus(id, status string, submittedAt time.Time, reason *string) error {

	_, err := r.DB.Exec(`
        UPDATE achievement_references
        SET status=$1, submitted_at=$2, rejection_note=$3,
            updated_at=NOW()
        WHERE id=$4
    `, status, submittedAt, reason, id)

	return err
}

//
// ========================================================
// VERIFY
// ========================================================
func (r *AchievementReferenceRepository) Verify(id, verifier string) error {

	_, err := r.DB.Exec(`
        UPDATE achievement_references
        SET status='approved',
            verified_at=NOW(),
            verified_by=$1,
            updated_at=NOW()
        WHERE id=$2
    `, verifier, id)

	return err
}

//
// ========================================================
// REJECT
// ========================================================
func (r *AchievementReferenceRepository) Reject(id, verifier, note string) error {

	_, err := r.DB.Exec(`
        UPDATE achievement_references
        SET status='rejected',
            verified_by=$1,
            rejection_note=$2,
            verified_at=NOW(),
            updated_at=NOW()
        WHERE id=$3
    `, verifier, note, id)

	return err
}
