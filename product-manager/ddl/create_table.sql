CREATE TABLE `yorushika_product`.`product` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `text` VARCHAR(200) NULL,
  `name` VARCHAR(45) NOT NULL,
  `price` INT NOT NULL,
  `img` VARCHAR(200) NOT NULL,
  `created_at` DATETIME NOT NULL,
  `apdated_at` DATETIME NOT NULL,
  PRIMARY KEY (`id`)
);