-- Update products table to match ERD design
ALTER TABLE tb_product
    ADD COLUMN description TEXT AFTER name,
    ADD COLUMN image_url TEXT AFTER quantity,
    ADD COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP AFTER image_url,
    MODIFY COLUMN price DECIMAL(10,2),
    MODIFY COLUMN quantity INT DEFAULT 0;