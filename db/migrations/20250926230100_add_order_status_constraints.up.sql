-- Add check constraints for order status to match ERD
ALTER TABLE tb_order
    ADD CONSTRAINT chk_tb_order_status CHECK (status IN ('pending', 'confirmed', 'processing', 'shipped', 'delivered', 'cancelled'));