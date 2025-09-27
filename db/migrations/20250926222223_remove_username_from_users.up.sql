-- Remove username column and update constraints
ALTER TABLE tb_user
    -- Drop the unique constraint on username
    DROP INDEX uk_tb_user_username,
    -- Drop the username column
    DROP COLUMN username,
    -- Make name, email, and role NOT NULL
    MODIFY COLUMN name VARCHAR(150) NOT NULL,
    MODIFY COLUMN email VARCHAR(255) NOT NULL,
    MODIFY COLUMN role VARCHAR(50) NOT NULL DEFAULT 'customer',
    -- Add unique constraint on email instead of username
    ADD CONSTRAINT uk_tb_user_email UNIQUE (email);