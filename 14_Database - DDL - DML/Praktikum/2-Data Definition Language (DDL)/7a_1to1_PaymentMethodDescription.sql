CREATE TABLE `alta_online_shop_`.`payment_description` 
(`id` INT NOT NULL AUTO_INCREMENT , 
`description` TEXT NOT NULL , 
`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP , 
`updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP , 
PRIMARY KEY (`id`)
) ENGINE = InnoDB;

ALTER TABLE `payment_description` 
ADD FOREIGN KEY (`id`) 
REFERENCES `payment_methods`(`id`) 
ON DELETE RESTRICT 
ON UPDATE RESTRICT;

ALTER TABLE `payment_description` 
ADD FOREIGN KEY (`id`) R
EFERENCES `payment_methods`(`id`) 
ON DELETE RESTRICT 
ON UPDATE RESTRICT;