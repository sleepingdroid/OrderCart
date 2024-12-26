CREATE TABLE `Users` (
  `id` INTEGER PRIMARY KEY AUTO_INCREMENT,
  `role` ENUM ('admin', 'vendor', 'customer') NOT NULL,
  `email` VARCHAR(255) UNIQUE,
  `username` VARCHAR(255) UNIQUE,
  `password` VARCHAR(255) NOT NULL,
  `addresses` VARCHAR(255),
  `created_at` DATETIME DEFAULT (CURRENT_TIMESTAMP),
  `updated_at` DATETIME DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE `Orders` (
  `id` INTEGER PRIMARY KEY AUTO_INCREMENT,
  `user_id` INTEGER NOT NULL,
  `status` ENUM ('created', 'pending', 'confirmed', 'deleted') NOT NULL,
  `subtotal` DECIMAL(10,2) DEFAULT 0,
  `texes` DECIMAL(10,2) DEFAULT 0,
  `total` DECIMAL(10,2) DEFAULT 0,
  `billing_address` VARCHAR(255),
  `shipping_address` VARCHAR(255),
  `created_at` DATETIME DEFAULT (CURRENT_TIMESTAMP),
  `updated_at` DATETIME DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE `OrderItems` (
  `order_id` INTEGER NOT NULL,
  `item_id` INTEGER NOT NULL,
  `quantity` INTEGER NOT NULL,
  `price` DECIMAL(10,2) NOT NULL,
  `created_at` DATETIME DEFAULT (CURRENT_TIMESTAMP),
  `updated_at` DATETIME DEFAULT (CURRENT_TIMESTAMP)
);

CREATE TABLE `Inventory` (
  `id` INTEGER PRIMARY KEY AUTO_INCREMENT,
  `vendor_id` INTEGER NOT NULL,
  `name` VARCHAR(255) NOT NULL,
  `description` TEXT,
  `stock` INTEGER NOT NULL,
  `price` DECIMAL(10,2) NOT NULL,
  `created_at` DATETIME DEFAULT (CURRENT_TIMESTAMP),
  `updated_at` DATETIME DEFAULT (CURRENT_TIMESTAMP)
);

ALTER TABLE `Orders` ADD FOREIGN KEY (`user_id`) REFERENCES `Users` (`id`);

ALTER TABLE `Inventory` ADD FOREIGN KEY (`vendor_id`) REFERENCES `Users` (`id`);

ALTER TABLE `Orders` ADD FOREIGN KEY (`billing_address`) REFERENCES `Users` (`addresses`);

ALTER TABLE `Orders` ADD FOREIGN KEY (`shipping_address`) REFERENCES `Users` (`addresses`);

ALTER TABLE `OrderItems` ADD FOREIGN KEY (`order_id`) REFERENCES `Orders` (`id`);

ALTER TABLE `OrderItems` ADD FOREIGN KEY (`item_id`) REFERENCES `Inventory` (`id`);
