package repository

import (
	"database/sql"
	"prestasi_backend/app/database"
	"prestasi_backend/app/model"
)

type RoleRepository struct {
	DB *sql.DB
}

func NewRoleRepository() *RoleRepository {
	return &RoleRepository{DB: database.PostgresDB}
}

func (r *RoleRepository) FindAll() ([]model.Role, error) {
	rows, err := r.DB.Query(`SELECT id, name, description, created_at FROM roles ORDER BY created_at`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.Role
	for rows.Next() {
		var x model.Role
		rows.Scan(&x.ID, &x.Name, &x.Description, &x.CreatedAt)
		list = append(list, x)
	}
	return list, nil
}

func (r *RoleRepository) FindByID(id string) (*model.Role, error) {
	row := r.DB.QueryRow(`SELECT id, name, description, created_at FROM roles WHERE id=$1`, id)

	var x model.Role
	err := row.Scan(&x.ID, &x.Name, &x.Description, &x.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &x, nil
}
