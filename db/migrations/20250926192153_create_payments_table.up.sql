-- Create payments table
CREATE TABLE tb_payment
(
    id INT AUTO_INCREMENT,
    order_id INT NOT NULL,
    transaction_id VARCHAR(255),
    status VARCHAR(50) DEFAULT 'pending',
    paid_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    CONSTRAINT fk_tb_order_payments FOREIGN KEY (order_id) REFERENCES tb_order (id) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT chk_tb_payment_status CHECK (status IN ('pending', 'success', 'failed', 'cancelled'))
);