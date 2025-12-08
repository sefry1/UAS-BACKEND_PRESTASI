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
	rows, err := r.DB.Query(`SELECT * FROM sp_get_all_users()`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []model.User

	for rows.Next() {
		var u model.User
		var roleName string
		err := rows.Scan(&u.ID, &u.Username, &u.Email, &u.FullName, &u.RoleID,
			&roleName, &u.IsActive, &u.CreatedAt)
		if err != nil {
			return nil, err
		}
		u.Role = &model.Role{Name: roleName}
		result = append(result, u)
	}

	return result, nil
}

func (r *UserRepository) FindByID(id string) (*model.User, error) {
	row := r.DB.QueryRow(`SELECT * FROM sp_get_user_by_id($1)`, id)

	var u model.User
	var roleName string
	err := row.Scan(&u.ID, &u.Username, &u.Email, &u.FullName,
		&u.RoleID, &roleName, &u.IsActive, &u.CreatedAt)
	if err != nil {
		return nil, err
	}

	u.Role = &model.Role{Name: roleName}
	return &u, nil
}

func (r *UserRepository) FindByUsername(username string) (*model.User, error) {
	row := r.DB.QueryRow(`SELECT * FROM sp_get_user_by_username($1)`, username)

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

	_, err := r.DB.Exec(`SELECT sp_create_user($1,$2,$3,$4,$5)`,
		username, email, string(hash), fullName, roleID)

	return err
}

func (r *UserRepository) Update(id, email, fullName, roleID string) error {
	_, err := r.DB.Exec(`SELECT sp_update_user($1,$2,$3)`, id, email, fullName)
	return err
}

func (r *UserRepository) Delete(id string) error {
	_, err := r.DB.Exec(`SELECT sp_delete_user($1)`, id)
	return err
}

func (r *UserRepository) UpdateRole(id, roleID string) error {
	_, err := r.DB.Exec(`SELECT sp_update_user_role($1,$2)`, id, roleID)
	return err
}
