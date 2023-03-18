SELECT products.*, product_types.name FROM products 
JOIN product_types 
ON products.product_type_id = product_types.id;