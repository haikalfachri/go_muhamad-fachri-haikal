SELECT t.*, p.name AS product_name, u.name AS user_name
FROM transactions t 
INNER JOIN transaction_details td ON t.id = td.transaction_id 
INNER JOIN products p ON td.product_id = p.id 
INNER JOIN users u ON t.user_id = u.id;
