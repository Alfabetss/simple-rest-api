CREATE TABLE `experience` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `company` varchar(50) COLLATE utf8mb4_general_ci NOT NULL,
  `talent_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `talent_id` (`talent_id`),
  CONSTRAINT `experience_ibfk_1` FOREIGN KEY (`talent_id`) REFERENCES `talent` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci