create TABLE IF NOT EXISTS pets (
  `pet_id` int(11) NOT NULL AUTO_INCREMENT,
  `name` text NOT NULL,
  `created_at` DATETIME,
  `updated_at` DATETIME,
  PRIMARY KEY (`pet_id`)
);
