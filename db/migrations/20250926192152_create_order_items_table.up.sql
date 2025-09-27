-- Create order_items table
CREATE TABLE tb_order_item
(
    id INT AUTO_INCREMENT,
    order_id INT NOT NULL,
    product_id VARCHAR(36) NOT NULL,
    quantity INT NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    CONSTRAINT fk_tb_order_order_items FOREIGN KEY (order_id) REFERENCES tb_order (id) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT fk_tb_product_order_items FOREIGN KEY (product_id) REFERENCES tb_product (product_id) ON DELETE CASCADE ON UPDATE CASCADE
);