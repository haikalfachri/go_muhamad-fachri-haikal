CREATE TRIGGER `delete_qty` 
AFTER DELETE ON `transaction_details`
FOR EACH ROW 
BEGIN
    UPDATE transactions SET total_qty = total_qty - OLD.qty
    WHERE id = OLD.transaction_id;
END

DELETE FROM `transaction_details` WHERE transaction_id = 15;