CREATE TABLE `alta_online_shop_`.`user_payment_method_details` 
(`user_id` INT NOT NULL , 
`payment_method_id` INT NOT NULL 
) ENGINE = InnoDB;

ALTER TABLE `user_payment_method_details` 
ADD FOREIGN KEY (`user_id`) 
REFERENCES `users`(`id`) 
ON DELETE RESTRICT 
ON UPDATE RESTRICT; 

ALTER TABLE `user_payment_method_details` 
ADD FOREIGN KEY (`payment_method_id`) 
REFERENCES `payment_methods`(`id`)
ON DELETE RESTRICT 
ON UPDATE RESTRICT;