-- Revert products table changes
ALTER TABLE tb_product
    DROP COLUMN description,
    DROP COLUMN image_url,
    DROP COLUMN created_at,
    MODIFY COLUMN price BIGINT,
    MODIFY COLUMN quantity INT;