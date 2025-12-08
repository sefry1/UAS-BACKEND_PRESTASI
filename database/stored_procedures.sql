-- ============================================
-- Stored Procedures for Prestasi System
-- Version: 1.0
-- ============================================

-- ============================================
-- 1. USER MANAGEMENT PROCEDURES
-- ============================================

-- Get user by username (for login)
CREATE OR REPLACE FUNCTION sp_get_user_by_username(p_username VARCHAR)
RETURNS TABLE (
    id UUID,
    username VARCHAR,
    email VARCHAR,
    password_hash VARCHAR,
    full_name VARCHAR,
    role_id UUID,
    is_active BOOLEAN,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
) AS $$
BEGIN
    RETURN QUERY
    SELECT u.id, u.username, u.email, u.password_hash, u.full_name, 
           u.role_id, u.is_active, u.created_at, u.updated_at
    FROM users u
    WHERE u.username = p_username AND u.is_active = true;
END;
$$ LANGUAGE plpgsql;

-- Get user by ID
CREATE OR REPLACE FUNCTION sp_get_user_by_id(p_user_id UUID)
RETURNS TABLE (
    id UUID,
    username VARCHAR,
    email VARCHAR,
    full_name VARCHAR,
    role_id UUID,
    role_name VARCHAR,
    is_active BOOLEAN,
    created_at TIMESTAMP
) AS $$
BEGIN
    RETURN QUERY
    SELECT u.id, u.username, u.email, u.full_name, u.role_id, 
           r.name as role_name, u.is_active, u.created_at
    FROM users u
    JOIN roles r ON u.role_id = r.id
    WHERE u.id = p_user_id;
END;
$$ LANGUAGE plpgsql;

-- Create new user
CREATE OR REPLACE FUNCTION sp_create_user(
    p_username VARCHAR,
    p_email VARCHAR,
    p_password_hash VARCHAR,
    p_full_name VARCHAR,
    p_role_id UUID
)
RETURNS UUID AS $$
DECLARE
    v_user_id UUID;
BEGIN
    INSERT INTO users (username, email, password_hash, full_name, role_id)
    VALUES (p_username, p_email, p_password_hash, p_full_name, p_role_id)
    RETURNING id INTO v_user_id;
    
    RETURN v_user_id;
END;
$$ LANGUAGE plpgsql;

-- Update user
CREATE OR REPLACE FUNCTION sp_update_user(
    p_user_id UUID,
    p_email VARCHAR,
    p_full_name VARCHAR
)
RETURNS BOOLEAN AS $$
BEGIN
    UPDATE users
    SET email = p_email,
        full_name = p_full_name,
        updated_at = NOW()
    WHERE id = p_user_id;
    
    RETURN FOUND;
END;
$$ LANGUAGE plpgsql;

-- Update user role
CREATE OR REPLACE FUNCTION sp_update_user_role(
    p_user_id UUID,
    p_role_id UUID
)
RETURNS BOOLEAN AS $$
BEGIN
    UPDATE users
    SET role_id = p_role_id,
        updated_at = NOW()
    WHERE id = p_user_id;
    
    RETURN FOUND;
END;
$$ LANGUAGE plpgsql;

-- Delete user (soft delete by setting is_active = false)
CREATE OR REPLACE FUNCTION sp_delete_user(p_user_id UUID)
RETURNS BOOLEAN AS $$
BEGIN
    UPDATE users
    SET is_active = false,
        updated_at = NOW()
    WHERE id = p_user_id;
    
    RETURN FOUND;
END;
$$ LANGUAGE plpgsql;

-- Get all users
CREATE OR REPLACE FUNCTION sp_get_all_users()
RETURNS TABLE (
    id UUID,
    username VARCHAR,
    email VARCHAR,
    full_name VARCHAR,
    role_id UUID,
    role_name VARCHAR,
    is_active BOOLEAN,
    created_at TIMESTAMP
) AS $$
BEGIN
    RETURN QUERY
    SELECT u.id, u.username, u.email, u.full_name, u.role_id,
           r.name as role_name, u.is_active, u.created_at
    FROM users u
    JOIN roles r ON u.role_id = r.id
    ORDER BY u.created_at DESC;
END;
$$ LANGUAGE plpgsql;

-- ============================================
-- 2. ROLE & PERMISSION PROCEDURES
-- ============================================

-- Get permissions for a role
CREATE OR REPLACE FUNCTION sp_get_role_permissions(p_role_id UUID)
RETURNS TABLE (
    permission_name VARCHAR
) AS $$
BEGIN
    RETURN QUERY
    SELECT p.name
    FROM permissions p
    JOIN role_permissions rp ON p.id = rp.permission_id
    WHERE rp.role_id = p_role_id;
END;
$$ LANGUAGE plpgsql;

-- Check if user has permission
CREATE OR REPLACE FUNCTION sp_user_has_permission(
    p_user_id UUID,
    p_permission_name VARCHAR
)
RETURNS BOOLEAN AS $$
DECLARE
    v_has_permission BOOLEAN;
BEGIN
    SELECT EXISTS (
        SELECT 1
        FROM users u
        JOIN role_permissions rp ON u.role_id = rp.role_id
        JOIN permissions p ON rp.permission_id = p.id
        WHERE u.id = p_user_id AND p.name = p_permission_name
    ) INTO v_has_permission;
    
    RETURN v_has_permission;
END;
$$ LANGUAGE plpgsql;

-- ============================================
-- 3. STUDENT MANAGEMENT PROCEDURES
-- ============================================

-- Create student profile
CREATE OR REPLACE FUNCTION sp_create_student(
    p_user_id UUID,
    p_student_id VARCHAR,
    p_program_study VARCHAR,
    p_academic_year VARCHAR,
    p_advisor_id UUID
)
RETURNS UUID AS $$
DECLARE
    v_student_id UUID;
BEGIN
    INSERT INTO students (user_id, student_id, program_study, academic_year, advisor_id)
    VALUES (p_user_id, p_student_id, p_program_study, p_academic_year, p_advisor_id)
    RETURNING id INTO v_student_id;
    
    RETURN v_student_id;
END;
$$ LANGUAGE plpgsql;

-- Get student by user_id
CREATE OR REPLACE FUNCTION sp_get_student_by_user_id(p_user_id UUID)
RETURNS TABLE (
    id UUID,
    user_id UUID,
    student_id VARCHAR,
    program_study VARCHAR,
    academic_year VARCHAR,
    advisor_id UUID,
    advisor_name VARCHAR,
    created_at TIMESTAMP
) AS $$
BEGIN
    RETURN QUERY
    SELECT s.id, s.user_id, s.student_id, s.program_study, s.academic_year,
           s.advisor_id, u.full_name as advisor_name, s.created_at
    FROM students s
    LEFT JOIN lecturers l ON s.advisor_id = l.id
    LEFT JOIN users u ON l.user_id = u.id
    WHERE s.user_id = p_user_id;
END;
$$ LANGUAGE plpgsql;

-- Get all students
CREATE OR REPLACE FUNCTION sp_get_all_students()
RETURNS TABLE (
    id UUID,
    user_id UUID,
    student_id VARCHAR,
    student_name VARCHAR,
    program_study VARCHAR,
    academic_year VARCHAR,
    advisor_id UUID,
    advisor_name VARCHAR
) AS $$
BEGIN
    RETURN QUERY
    SELECT s.id, s.user_id, s.student_id, u.full_name as student_name,
           s.program_study, s.academic_year, s.advisor_id, 
           lec_user.full_name as advisor_name
    FROM students s
    JOIN users u ON s.user_id = u.id
    LEFT JOIN lecturers l ON s.advisor_id = l.id
    LEFT JOIN users lec_user ON l.user_id = lec_user.id
    ORDER BY s.student_id;
END;
$$ LANGUAGE plpgsql;

-- Set student advisor
CREATE OR REPLACE FUNCTION sp_set_student_advisor(
    p_student_id UUID,
    p_advisor_id UUID
)
RETURNS BOOLEAN AS $$
BEGIN
    UPDATE students
    SET advisor_id = p_advisor_id
    WHERE id = p_student_id;
    
    RETURN FOUND;
END;
$$ LANGUAGE plpgsql;

-- ============================================
-- 4. LECTURER MANAGEMENT PROCEDURES
-- ============================================

-- Create lecturer profile
CREATE OR REPLACE FUNCTION sp_create_lecturer(
    p_user_id UUID,
    p_lecturer_id VARCHAR,
    p_department VARCHAR
)
RETURNS UUID AS $$
DECLARE
    v_lecturer_id UUID;
BEGIN
    INSERT INTO lecturers (user_id, lecturer_id, department)
    VALUES (p_user_id, p_lecturer_id, p_department)
    RETURNING id INTO v_lecturer_id;
    
    RETURN v_lecturer_id;
END;
$$ LANGUAGE plpgsql;

-- Get lecturer by user_id
CREATE OR REPLACE FUNCTION sp_get_lecturer_by_user_id(p_user_id UUID)
RETURNS TABLE (
    id UUID,
    user_id UUID,
    lecturer_id VARCHAR,
    department VARCHAR,
    created_at TIMESTAMP
) AS $$
BEGIN
    RETURN QUERY
    SELECT l.id, l.user_id, l.lecturer_id, l.department, l.created_at
    FROM lecturers l
    WHERE l.user_id = p_user_id;
END;
$$ LANGUAGE plpgsql;

-- Get all lecturers
CREATE OR REPLACE FUNCTION sp_get_all_lecturers()
RETURNS TABLE (
    id UUID,
    user_id UUID,
    lecturer_id VARCHAR,
    lecturer_name VARCHAR,
    department VARCHAR
) AS $$
BEGIN
    RETURN QUERY
    SELECT l.id, l.user_id, l.lecturer_id, u.full_name as lecturer_name, l.department
    FROM lecturers l
    JOIN users u ON l.user_id = u.id
    ORDER BY l.lecturer_id;
END;
$$ LANGUAGE plpgsql;

-- Get lecturer's advisees
CREATE OR REPLACE FUNCTION sp_get_lecturer_advisees(p_lecturer_id UUID)
RETURNS TABLE (
    student_id UUID,
    student_nim VARCHAR,
    student_name VARCHAR,
    program_study VARCHAR,
    academic_year VARCHAR,
    total_achievements BIGINT
) AS $$
BEGIN
    RETURN QUERY
    SELECT s.id, s.student_id, u.full_name as student_name,
           s.program_study, s.academic_year,
           COUNT(ar.id) as total_achievements
    FROM students s
    JOIN users u ON s.user_id = u.id
    LEFT JOIN achievement_references ar ON s.id = ar.student_id
    WHERE s.advisor_id = p_lecturer_id
    GROUP BY s.id, s.student_id, u.full_name, s.program_study, s.academic_year
    ORDER BY s.student_id;
END;
$$ LANGUAGE plpgsql;

-- ============================================
-- 5. ACHIEVEMENT REFERENCE PROCEDURES
-- ============================================

-- Create achievement reference
CREATE OR REPLACE FUNCTION sp_create_achievement_reference(
    p_student_id UUID,
    p_mongo_achievement_id VARCHAR
)
RETURNS UUID AS $$
DECLARE
    v_achievement_id UUID;
BEGIN
    INSERT INTO achievement_references (student_id, mongo_achievement_id, status)
    VALUES (p_student_id, p_mongo_achievement_id, 'draft')
    RETURNING id INTO v_achievement_id;
    
    RETURN v_achievement_id;
END;
$$ LANGUAGE plpgsql;

-- Get achievement reference by ID
CREATE OR REPLACE FUNCTION sp_get_achievement_by_id(p_achievement_id UUID)
RETURNS TABLE (
    id UUID,
    student_id UUID,
    mongo_achievement_id VARCHAR,
    status VARCHAR,
    submitted_at TIMESTAMP,
    verified_at TIMESTAMP,
    verified_by UUID,
    verified_by_name VARCHAR,
    rejection_note TEXT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
) AS $$
BEGIN
    RETURN QUERY
    SELECT ar.id, ar.student_id, ar.mongo_achievement_id, ar.status,
           ar.submitted_at, ar.verified_at, ar.verified_by,
           u.full_name as verified_by_name,
           ar.rejection_note, ar.created_at, ar.updated_at
    FROM achievement_references ar
    LEFT JOIN users u ON ar.verified_by = u.id
    WHERE ar.id = p_achievement_id;
END;
$$ LANGUAGE plpgsql;

-- Get achievements by student_id
CREATE OR REPLACE FUNCTION sp_get_achievements_by_student(p_student_id UUID)
RETURNS TABLE (
    id UUID,
    mongo_achievement_id VARCHAR,
    status VARCHAR,
    submitted_at TIMESTAMP,
    verified_at TIMESTAMP,
    created_at TIMESTAMP
) AS $$
BEGIN
    RETURN QUERY
    SELECT ar.id, ar.mongo_achievement_id, ar.status,
           ar.submitted_at, ar.verified_at, ar.created_at
    FROM achievement_references ar
    WHERE ar.student_id = p_student_id
    ORDER BY ar.created_at DESC;
END;
$$ LANGUAGE plpgsql;

-- Get achievements by advisor (for Dosen Wali)
CREATE OR REPLACE FUNCTION sp_get_achievements_by_advisor(p_lecturer_id UUID)
RETURNS TABLE (
    id UUID,
    student_id UUID,
    student_name VARCHAR,
    mongo_achievement_id VARCHAR,
    status VARCHAR,
    submitted_at TIMESTAMP,
    created_at TIMESTAMP
) AS $$
BEGIN
    RETURN QUERY
    SELECT ar.id, s.id as student_id, u.full_name as student_name,
           ar.mongo_achievement_id, ar.status, ar.submitted_at, ar.created_at
    FROM achievement_references ar
    JOIN students s ON ar.student_id = s.id
    JOIN users u ON s.user_id = u.id
    WHERE s.advisor_id = p_lecturer_id
    ORDER BY ar.created_at DESC;
END;
$$ LANGUAGE plpgsql;

-- Get all achievements (for Admin)
CREATE OR REPLACE FUNCTION sp_get_all_achievements()
RETURNS TABLE (
    id UUID,
    student_id UUID,
    student_name VARCHAR,
    mongo_achievement_id VARCHAR,
    status VARCHAR,
    submitted_at TIMESTAMP,
    verified_at TIMESTAMP,
    created_at TIMESTAMP
) AS $$
BEGIN
    RETURN QUERY
    SELECT ar.id, s.id as student_id, u.full_name as student_name,
           ar.mongo_achievement_id, ar.status, ar.submitted_at,
           ar.verified_at, ar.created_at
    FROM achievement_references ar
    JOIN students s ON ar.student_id = s.id
    JOIN users u ON s.user_id = u.id
    ORDER BY ar.created_at DESC;
END;
$$ LANGUAGE plpgsql;

-- Submit achievement for verification
CREATE OR REPLACE FUNCTION sp_submit_achievement(p_achievement_id UUID)
RETURNS BOOLEAN AS $$
BEGIN
    UPDATE achievement_references
    SET status = 'submitted',
        submitted_at = NOW(),
        updated_at = NOW()
    WHERE id = p_achievement_id AND status = 'draft';
    
    RETURN FOUND;
END;
$$ LANGUAGE plpgsql;

-- Verify achievement
CREATE OR REPLACE FUNCTION sp_verify_achievement(
    p_achievement_id UUID,
    p_verifier_id UUID
)
RETURNS BOOLEAN AS $$
BEGIN
    UPDATE achievement_references
    SET status = 'verified',
        verified_at = NOW(),
        verified_by = p_verifier_id,
        updated_at = NOW()
    WHERE id = p_achievement_id AND status = 'submitted';
    
    RETURN FOUND;
END;
$$ LANGUAGE plpgsql;

-- Reject achievement
CREATE OR REPLACE FUNCTION sp_reject_achievement(
    p_achievement_id UUID,
    p_verifier_id UUID,
    p_rejection_note TEXT
)
RETURNS BOOLEAN AS $$
BEGIN
    UPDATE achievement_references
    SET status = 'rejected',
        verified_by = p_verifier_id,
        rejection_note = p_rejection_note,
        updated_at = NOW()
    WHERE id = p_achievement_id AND status = 'submitted';
    
    RETURN FOUND;
END;
$$ LANGUAGE plpgsql;

-- Delete achievement (only if draft)
CREATE OR REPLACE FUNCTION sp_delete_achievement(p_achievement_id UUID)
RETURNS BOOLEAN AS $$
BEGIN
    DELETE FROM achievement_references
    WHERE id = p_achievement_id AND status = 'draft';
    
    RETURN FOUND;
END;
$$ LANGUAGE plpgsql;

-- ============================================
-- 6. REPORTING & STATISTICS PROCEDURES
-- ============================================

-- Get achievement statistics
CREATE OR REPLACE FUNCTION sp_get_achievement_statistics()
RETURNS TABLE (
    total_achievements BIGINT,
    total_draft BIGINT,
    total_submitted BIGINT,
    total_verified BIGINT,
    total_rejected BIGINT
) AS $$
BEGIN
    RETURN QUERY
    SELECT
        COUNT(*) as total_achievements,
        COUNT(*) FILTER (WHERE status = 'draft') as total_draft,
        COUNT(*) FILTER (WHERE status = 'submitted') as total_submitted,
        COUNT(*) FILTER (WHERE status = 'verified') as total_verified,
        COUNT(*) FILTER (WHERE status = 'rejected') as total_rejected
    FROM achievement_references;
END;
$$ LANGUAGE plpgsql;

-- Get student achievement report
CREATE OR REPLACE FUNCTION sp_get_student_report(p_student_id UUID)
RETURNS TABLE (
    student_name VARCHAR,
    student_nim VARCHAR,
    program_study VARCHAR,
    total_achievements BIGINT,
    total_verified BIGINT,
    total_rejected BIGINT,
    total_pending BIGINT
) AS $$
BEGIN
    RETURN QUERY
    SELECT
        u.full_name as student_name,
        s.student_id as student_nim,
        s.program_study,
        COUNT(ar.id) as total_achievements,
        COUNT(*) FILTER (WHERE ar.status = 'verified') as total_verified,
        COUNT(*) FILTER (WHERE ar.status = 'rejected') as total_rejected,
        COUNT(*) FILTER (WHERE ar.status IN ('draft', 'submitted')) as total_pending
    FROM students s
    JOIN users u ON s.user_id = u.id
    LEFT JOIN achievement_references ar ON s.id = ar.student_id
    WHERE s.id = p_student_id
    GROUP BY u.full_name, s.student_id, s.program_study;
END;
$$ LANGUAGE plpgsql;

COMMIT;
