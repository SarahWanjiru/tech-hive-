-- Remove check constraints for order status
ALTER TABLE tb_order
    DROP CONSTRAINT chk_tb_order_status;