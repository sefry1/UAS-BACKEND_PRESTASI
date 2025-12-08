-- ============================================
-- Database Seeder Script
-- Sistem Pelaporan Prestasi Mahasiswa
-- Version: 1.0
-- ============================================

BEGIN;

-- Enable pgcrypto extension for password hashing
CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- ============================================
-- 1. SEED ROLES
-- ============================================
INSERT INTO roles (id, name, description) VALUES
    ('11111111-1111-1111-1111-111111111111', 'Admin', 'Pengelola sistem'),
    ('22222222-2222-2222-2222-222222222222', 'Mahasiswa', 'Pelapor prestasi'),
    ('33333333-3333-3333-3333-333333333333', 'Dosen Wali', 'Verifikator prestasi mahasiswa bimbingannya')
ON CONFLICT (name) DO NOTHING;

-- ============================================
-- 2. SEED PERMISSIONS
-- ============================================
INSERT INTO permissions (id, name, resource, action, description) VALUES
    ('a0000000-0000-0000-0000-000000000001', 'achievement:create', 'achievement', 'create', 'Create new achievement'),
    ('a0000000-0000-0000-0000-000000000002', 'achievement:read', 'achievement', 'read', 'Read achievement data'),
    ('a0000000-0000-0000-0000-000000000003', 'achievement:update', 'achievement', 'update', 'Update achievement data'),
    ('a0000000-0000-0000-0000-000000000004', 'achievement:delete', 'achievement', 'delete', 'Delete achievement'),
    ('a0000000-0000-0000-0000-000000000005', 'achievement:verify', 'achievement', 'verify', 'Verify achievement submission'),
    ('a0000000-0000-0000-0000-000000000006', 'achievement:reject', 'achievement', 'reject', 'Reject achievement submission'),
    ('a0000000-0000-0000-0000-000000000007', 'achievement:submit', 'achievement', 'submit', 'Submit achievement for verification'),
    ('a0000000-0000-0000-0000-000000000008', 'user:manage', 'user', 'manage', 'Manage user accounts and roles')
ON CONFLICT (name) DO NOTHING;

-- ============================================
-- 3. SEED ROLE_PERMISSIONS MAPPING
-- ============================================

-- Admin - ALL PERMISSIONS
INSERT INTO role_permissions (role_id, permission_id) VALUES
    ('11111111-1111-1111-1111-111111111111', 'a0000000-0000-0000-0000-000000000001'),
    ('11111111-1111-1111-1111-111111111111', 'a0000000-0000-0000-0000-000000000002'),
    ('11111111-1111-1111-1111-111111111111', 'a0000000-0000-0000-0000-000000000003'),
    ('11111111-1111-1111-1111-111111111111', 'a0000000-0000-0000-0000-000000000004'),
    ('11111111-1111-1111-1111-111111111111', 'a0000000-0000-0000-0000-000000000005'),
    ('11111111-1111-1111-1111-111111111111', 'a0000000-0000-0000-0000-000000000006'),
    ('11111111-1111-1111-1111-111111111111', 'a0000000-0000-0000-0000-000000000007'),
    ('11111111-1111-1111-1111-111111111111', 'a0000000-0000-0000-0000-000000000008')
ON CONFLICT DO NOTHING;

-- Mahasiswa - Create, Read, Update, Delete, Submit own achievements
INSERT INTO role_permissions (role_id, permission_id) VALUES
    ('22222222-2222-2222-2222-222222222222', 'a0000000-0000-0000-0000-000000000001'),
    ('22222222-2222-2222-2222-222222222222', 'a0000000-0000-0000-0000-000000000002'),
    ('22222222-2222-2222-2222-222222222222', 'a0000000-0000-0000-0000-000000000003'),
    ('22222222-2222-2222-2222-222222222222', 'a0000000-0000-0000-0000-000000000004'),
    ('22222222-2222-2222-2222-222222222222', 'a0000000-0000-0000-0000-000000000007')
ON CONFLICT DO NOTHING;

-- Dosen Wali - Read, Verify, Reject achievements
INSERT INTO role_permissions (role_id, permission_id) VALUES
    ('33333333-3333-3333-3333-333333333333', 'a0000000-0000-0000-0000-000000000002'),
    ('33333333-3333-3333-3333-333333333333', 'a0000000-0000-0000-0000-000000000005'),
    ('33333333-3333-3333-3333-333333333333', 'a0000000-0000-0000-0000-000000000006')
ON CONFLICT DO NOTHING;

-- ============================================
-- 4. SEED TEST USERS
-- Password for all: 123456 (will be hashed with pgcrypto)
-- ============================================
INSERT INTO users (id, username, email, password_hash, full_name, role_id, is_active) VALUES
    ('a1111111-1111-1111-1111-111111111111', 'admin', 'admin@prestasi.ac.id', crypt('123456', gen_salt('bf', 10)), 'Admin System', '11111111-1111-1111-1111-111111111111', true),
    ('b2222222-2222-2222-2222-222222222222', 'mahasiswa1', 'mahasiswa1@student.ac.id', crypt('123456', gen_salt('bf', 10)), 'Budi Santoso', '22222222-2222-2222-2222-222222222222', true),
    ('b2222222-2222-2222-2222-222222222223', 'mahasiswa2', 'mahasiswa2@student.ac.id', crypt('123456', gen_salt('bf', 10)), 'Siti Nurhaliza', '22222222-2222-2222-2222-222222222222', true),
    ('c3333333-3333-3333-3333-333333333333', 'dosenwali1', 'dosenwali1@lecturer.ac.id', crypt('123456', gen_salt('bf', 10)), 'Dr. Ahmad Wijaya', '33333333-3333-3333-3333-333333333333', true),
    ('c3333333-3333-3333-3333-333333333334', 'dosenwali2', 'dosenwali2@lecturer.ac.id', crypt('123456', gen_salt('bf', 10)), 'Dr. Sri Mulyani', '33333333-3333-3333-3333-333333333333', true)
ON CONFLICT (username) DO NOTHING;

-- ============================================
-- 5. SEED LECTURERS
-- ============================================
INSERT INTO lecturers (id, user_id, lecturer_id, department) VALUES
    ('d3333333-3333-3333-3333-333333333333', 'c3333333-3333-3333-3333-333333333333', 'NIP001', 'Teknik Informatika'),
    ('d3333333-3333-3333-3333-333333333334', 'c3333333-3333-3333-3333-333333333334', 'NIP002', 'Teknik Informatika')
ON CONFLICT (lecturer_id) DO NOTHING;

-- ============================================
-- 6. SEED STUDENTS
-- ============================================
INSERT INTO students (id, user_id, student_id, program_study, academic_year, advisor_id) VALUES
    ('e2222222-2222-2222-2222-222222222222', 'b2222222-2222-2222-2222-222222222222', 'NIM001', 'Teknik Informatika', '2023/2024', 'd3333333-3333-3333-3333-333333333333'),
    ('e2222222-2222-2222-2222-222222222223', 'b2222222-2222-2222-2222-222222222223', 'NIM002', 'Teknik Informatika', '2023/2024', 'd3333333-3333-3333-3333-333333333334')
ON CONFLICT (student_id) DO NOTHING;

COMMIT;

-- ============================================
-- VERIFICATION QUERIES
-- ============================================
-- Run these to verify seeder worked correctly

-- Check roles count (should be 3)
-- SELECT COUNT(*) FROM roles;

-- Check permissions count (should be 8)
-- SELECT COUNT(*) FROM permissions;

-- Check users count (should be 5)
-- SELECT COUNT(*) FROM users;

-- Check students count (should be 2)
-- SELECT COUNT(*) FROM students;

-- Check lecturers count (should be 2)
-- SELECT COUNT(*) FROM lecturers;

-- Check role_permissions mapping (Admin should have 8, Mahasiswa 5, Dosen Wali 3)
-- SELECT r.name, COUNT(rp.permission_id) as permission_count 
-- FROM roles r 
-- LEFT JOIN role_permissions rp ON r.id = rp.role_id 
-- GROUP BY r.name;
