package service

import "prestasi_backend/app/repository"

var (
    UserRepo            *repository.UserRepository
    RoleRepo            *repository.RoleRepository
    PermissionRepo      *repository.PermissionRepository
    RolePermissionRepo  *repository.RolePermissionRepository
    StudentRepo         *repository.StudentRepository
    LecturerRepo        *repository.LecturerRepository
    AchievementRefRepo  *repository.AchievementReferenceRepository
    AchievementMongoRepo *repository.AchievementMongoRepository
)

func InitService() {
    UserRepo = repository.NewUserRepository()
    RoleRepo = repository.NewRoleRepository()
    PermissionRepo = repository.NewPermissionRepository()
    RolePermissionRepo = repository.NewRolePermissionRepository()
    StudentRepo = repository.NewStudentRepository()
    LecturerRepo = repository.NewLecturerRepository()
    AchievementRefRepo = repository.NewAchievementReferenceRepository()
    AchievementMongoRepo = repository.NewAchievementMongoRepository()
}
