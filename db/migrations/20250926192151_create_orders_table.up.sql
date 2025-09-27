-- Create orders table
CREATE TABLE tb_order
(
    id INT AUTO_INCREMENT,
    user_id INT NOT NULL,
    total DECIMAL(10,2) NOT NULL,
    status VARCHAR(50) DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    CONSTRAINT fk_tb_user_orders FOREIGN KEY (user_id) REFERENCES tb_user (id) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT chk_tb_order_status CHECK (status IN ('pending', 'confirmed', 'processing', 'shipped', 'delivered', 'cancelled'))
);