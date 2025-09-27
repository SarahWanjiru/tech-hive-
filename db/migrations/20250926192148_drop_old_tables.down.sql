-- Recreate old tables if needed for rollback
CREATE TABLE tb_user_role
(
    user_role_id varchar(36),
    role         varchar(10),
    username     varchar(100),
    PRIMARY KEY (user_role_id),
    CONSTRAINT fk_tb_user_user_roles FOREIGN KEY (username) REFERENCES tb_user (username) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE tb_transaction
(
    transaction_id varchar(36),
    total_price    bigint,
    PRIMARY KEY (transaction_id)
);

CREATE TABLE tb_transaction_detail
(
    transaction_detail_id varchar(36),
    sub_total_price       bigint,
    price                 bigint,
    quantity              int,
    transaction_id        varchar(36),
    product_id            varchar(36),
    PRIMARY KEY (transaction_detail_id),
    CONSTRAINT fk_tb_product_transaction_details FOREIGN KEY (product_id) REFERENCES tb_product (product_id) ON DELETE SET NULL ON UPDATE CASCADE,
    CONSTRAINT fk_tb_transaction_transaction_details FOREIGN KEY (transaction_id) REFERENCES tb_transaction (transaction_id) ON DELETE CASCADE ON UPDATE CASCADE
);