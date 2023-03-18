SELECT COUNT(t.transaction_id) FROM transaction_details t 
JOIN products p 
ON p.id = t.product_id 
WHERE p.product_type_id = 2;