-- Revert products table to use UUID primary key
ALTER TABLE tb_product
    DROP COLUMN id,
    ADD PRIMARY KEY (product_id);