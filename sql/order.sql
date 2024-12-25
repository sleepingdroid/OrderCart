CREATE TABLE `Users` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255),
  `email` varchar(255) UNIQUE,
  `password` varchar(255),
  `created_at` datetime,
  `updated_at` datetime
);

CREATE TABLE `Orders` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int,
  `status` varchar(255),
  `total` decimal,
  `created_at` datetime,
  `updated_at` datetime
);

CREATE TABLE `OrderItems` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `order_id` int,
  `item_id` int,
  `quantity` int,
  `price` decimal,
  `created_at` datetime,
  `updated_at` datetime
);

CREATE TABLE `Inventory` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255),
  `description` text,
  `stock` int,
  `price` decimal,
  `created_at` datetime,
  `updated_at` datetime
);

CREATE TABLE `Payments` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `order_id` int,
  `payment_method` varchar(255),
  `status` varchar(255),
  `amount` decimal,
  `created_at` datetime,
  `updated_at` datetime
);

ALTER TABLE `Orders` ADD FOREIGN KEY (`user_id`) REFERENCES `Users` (`id`);

ALTER TABLE `OrderItems` ADD FOREIGN KEY (`order_id`) REFERENCES `Orders` (`id`);

ALTER TABLE `Payments` ADD FOREIGN KEY (`order_id`) REFERENCES `Orders` (`id`);
