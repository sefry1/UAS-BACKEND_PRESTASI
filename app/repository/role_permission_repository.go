package repository

import (
	"database/sql"
	"prestasi_backend/app/database"
)

type RolePermissionRepository struct {
	DB *sql.DB
}

func NewRolePermissionRepository() *RolePermissionRepository {
	return &RolePermissionRepository{DB: database.PostgresDB}
}

func (r *RolePermissionRepository) GetPermissions(roleID string) ([]string, error) {
	rows, err := r.DB.Query(`
        SELECT p.name 
        FROM role_permissions rp
        JOIN permissions p ON rp.permission_id = p.id
        WHERE rp.role_id=$1
    `, roleID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []string
	for rows.Next() {
		var perm string
		rows.Scan(&perm)
		list = append(list, perm)
	}

	return list, nil
}
