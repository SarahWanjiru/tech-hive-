-- Revert users table changes
ALTER TABLE tb_user
    DROP COLUMN id,
    DROP COLUMN name,
    DROP COLUMN email,
    DROP COLUMN role,
    DROP COLUMN created_at,
    ADD COLUMN username VARCHAR(255) PRIMARY KEY;