-- Create cart tables
CREATE TABLE tb_cart
(
    id INT AUTO_INCREMENT,
    user_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    CONSTRAINT fk_tb_user_cart FOREIGN KEY (user_id) REFERENCES tb_user (id) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT uk_tb_cart_user_id UNIQUE (user_id)
);

CREATE TABLE tb_cart_item
(
    id INT AUTO_INCREMENT,
    cart_id INT NOT NULL,
    product_id VARCHAR(36) NOT NULL,
    quantity INT NOT NULL CHECK (quantity > 0),
    price DECIMAL(10,2) NOT NULL CHECK (price >= 0),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    CONSTRAINT fk_tb_cart_cart_items FOREIGN KEY (cart_id) REFERENCES tb_cart (id) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT fk_tb_product_cart_items FOREIGN KEY (product_id) REFERENCES tb_product (product_id) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT uk_tb_cart_item_cart_product UNIQUE (cart_id, product_id)
);