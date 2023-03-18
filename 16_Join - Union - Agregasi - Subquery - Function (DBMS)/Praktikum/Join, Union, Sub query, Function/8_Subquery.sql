SELECT * FROM products
WHERE id NOT IN (
  SELECT product_id
  FROM transaction_details
);