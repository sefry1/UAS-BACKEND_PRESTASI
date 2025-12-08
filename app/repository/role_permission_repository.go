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

// GetPermissions - Get all permissions for a role using stored procedure
func (r *RolePermissionRepository) GetPermissions(roleID string) ([]string, error) {
	rows, err := r.DB.Query(`SELECT * FROM sp_get_role_permissions($1)`, roleID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []string
	for rows.Next() {
		var perm string
		err := rows.Scan(&perm)
		if err != nil {
			return nil, err
		}
		list = append(list, perm)
	}

	return list, nil
}

// HasPermission - Check if user has specific permission using stored procedure
func (r *RolePermissionRepository) HasPermission(userID, permission string) (bool, error) {
	var hasPermission bool
	err := r.DB.QueryRow(`SELECT sp_user_has_permission($1, $2)`, userID, permission).Scan(&hasPermission)
	if err != nil {
		return false, err
	}
	return hasPermission, nil
}
