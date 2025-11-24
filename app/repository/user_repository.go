package repository

import (
	"database/sql"
	"prestasi_backend/app/database"
	"prestasi_backend/app/model"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{DB: database.PostgresDB}
}

func (r *UserRepository) FindAll() ([]model.User, error) {
	rows, err := r.DB.Query(`
        SELECT id, username, email, full_name, role_id, is_active, created_at, updated_at 
        FROM users ORDER BY created_at DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []model.User

	for rows.Next() {
		var u model.User
		err := rows.Scan(&u.ID, &u.Username, &u.Email, &u.FullName, &u.RoleID,
			&u.IsActive, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			return nil, err
		}
		result = append(result, u)
	}

	return result, nil
}

func (r *UserRepository) FindByID(id string) (*model.User, error) {
	row := r.DB.QueryRow(`
        SELECT id, username, email, full_name, role_id, is_active, created_at, updated_at
        FROM users WHERE id=$1`, id)

	var u model.User
	err := row.Scan(&u.ID, &u.Username, &u.Email, &u.FullName,
		&u.RoleID, &u.IsActive, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *UserRepository) FindByUsername(username string) (*model.User, error) {
	row := r.DB.QueryRow(`
        SELECT id, username, email, password_hash, full_name, role_id, is_active, created_at, updated_at
        FROM users WHERE username=$1`, username)

	var u model.User
	err := row.Scan(&u.ID, &u.Username, &u.Email, &u.PasswordHash,
		&u.FullName, &u.RoleID, &u.IsActive, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) Create(username, email, password, fullName, roleID string) error {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	_, err := r.DB.Exec(`
        INSERT INTO users (username, email, password_hash, full_name, role_id, created_at, updated_at)
        VALUES ($1,$2,$3,$4,$5,NOW(),NOW())
    `, username, email, string(hash), fullName, roleID)

	return err
}

func (r *UserRepository) Update(id, email, fullName, roleID string) error {
	_, err := r.DB.Exec(`
        UPDATE users 
        SET email=$1, full_name=$2, role_id=$3, updated_at=NOW()
        WHERE id=$4
    `, email, fullName, roleID, id)
	return err
}

func (r *UserRepository) Delete(id string) error {
	_, err := r.DB.Exec(`DELETE FROM users WHERE id=$1`, id)
	return err
}

func (r *UserRepository) UpdateRole(id, roleID string) error {
	_, err := r.DB.Exec(`UPDATE users SET role_id=$1 WHERE id=$2`, roleID, id)
	return err
}
