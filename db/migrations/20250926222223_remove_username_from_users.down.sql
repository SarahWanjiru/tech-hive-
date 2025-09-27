-- Reverse: Add back username column and revert constraints
ALTER TABLE tb_user
    -- Drop the unique constraint on email
    DROP INDEX uk_tb_user_email,
    -- Add back the username column
    ADD COLUMN username VARCHAR(100) UNIQUE AFTER id,
    -- Revert name, email, and role to nullable
    MODIFY COLUMN name VARCHAR(150) NULL,
    MODIFY COLUMN email VARCHAR(255) NULL,
    MODIFY COLUMN role VARCHAR(50) NULL DEFAULT 'customer',
    -- Add back unique constraint on username
    ADD CONSTRAINT uk_tb_user_username UNIQUE (username);