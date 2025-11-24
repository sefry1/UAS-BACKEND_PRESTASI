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
	return &AchievementReferenceRepository{DB: database.PostgresDB}
}

func (r *AchievementReferenceRepository) FindByID(id string) (*model.AchievementReference, error) {

	row := r.DB.QueryRow(`
        SELECT id, student_id, mongo_achievement_id, status,
               submitted_at, verified_at, verified_by,
               rejection_note, created_at, updated_at
        FROM achievement_references
        WHERE id=$1
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

func (r *AchievementReferenceRepository) FindByStudent(studentID string) ([]model.AchievementReference, error) {
	rows, err := r.DB.Query(`
        SELECT id, student_id, mongo_achievement_id, status,
               submitted_at, verified_at, verified_by,
               rejection_note, created_at, updated_at
        FROM achievement_references
        WHERE student_id=$1`, studentID)
	if err != nil {
		return nil, err
	}

	var list []model.AchievementReference
	for rows.Next() {
		var a model.AchievementReference
		rows.Scan(
			&a.ID, &a.StudentID, &a.MongoAchievementID, &a.Status,
			&a.SubmittedAt, &a.VerifiedAt, &a.VerifiedBy,
			&a.RejectionNote, &a.CreatedAt, &a.UpdatedAt,
		)
		list = append(list, a)
	}

	return list, nil
}

func (r *AchievementReferenceRepository) Submit(id string) error {
	_, err := r.DB.Exec(`
        UPDATE achievement_references
        SET status='submitted', submitted_at=$1, updated_at=$1
        WHERE id=$2
    `, time.Now(), id)
	return err
}

func (r *AchievementReferenceRepository) Verify(id, verifier string) error {
	_, err := r.DB.Exec(`
        UPDATE achievement_references
        SET status='verified', verified_by=$1, verified_at=$2
        WHERE id=$3
    `, verifier, time.Now(), id)
	return err
}

func (r *AchievementReferenceRepository) Reject(id, verifier, note string) error {
	_, err := r.DB.Exec(`
        UPDATE achievement_references
        SET status='rejected', verified_by=$1, verified_at=$2, rejection_note=$3
        WHERE id=$4
    `, verifier, time.Now(), note, id)
	return err
}
