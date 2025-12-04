package repository

import (
	"database/sql"
	"prestasi_backend/app/database"
	"prestasi_backend/app/model"
)

type PermissionRepository struct {
	DB *sql.DB
}

func NewPermissionRepository() *PermissionRepository {
	return &PermissionRepository{DB: database.PostgresDB}
}

func (r *PermissionRepository) GetAll() ([]model.Permission, error) {
	rows, err := r.DB.Query(`
        SELECT id, name, resource, action, description, created_at
        FROM permissions`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.Permission
	for rows.Next() {
		var p model.Permission
		rows.Scan(&p.ID, &p.Name, &p.Resource, &p.Action, &p.Description, &p.CreatedAt)
		list = append(list, p)
	}
	return list, nil
}
