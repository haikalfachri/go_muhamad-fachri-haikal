CREATE TRIGGER `delete_transaction_details` 
AFTER DELETE ON `transactions`
FOR EACH ROW 
BEGIN
    DECLARE t_id INT;
    SET t_id = OLD.id;
    DELETE FROM transaction_details WHERE transaction_id = t_id;
END