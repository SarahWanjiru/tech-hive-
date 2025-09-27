-- Update products table to use INT primary key instead of UUID to match ERD
ALTER TABLE tb_product
    ADD COLUMN id INT AUTO_INCREMENT PRIMARY KEY FIRST,
    DROP PRIMARY KEY,
    ADD CONSTRAINT fk_tb_product_id UNIQUE (product_id);