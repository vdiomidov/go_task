
CREATE DATABASE IF NOT EXISTS `adv_db` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci */;
USE `adv_db`;

CREATE TABLE IF NOT EXISTS `adv` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

DELETE FROM `adv`;

INSERT INTO `adv` (`id`, `name`) VALUES
	(1, 'a1'),
	(2, 'a2'),
	(3, 'a3'),
	(4, 'a4'),
	(5, 'a5');

CREATE TABLE IF NOT EXISTS `sessions` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `active` tinyint(1) DEFAULT 1,
  `created_at` datetime NOT NULL DEFAULT current_timestamp(),
  `price` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


DELETE FROM `sessions`;

INSERT INTO `sessions` (`id`, `user_id`, `active`, `created_at`, `price`) VALUES
	(2, '111', 0, '2020-06-03 23:23:13', 10),
	(3, '222', 0, '2020-06-04 00:23:19', 20),
	(4, '111', 0, '2020-06-04 11:00:16', 11),
	(5, '111', 0, '2020-06-04 12:22:19', 30),
	(6, '222', 0, '2020-06-04 10:34:09', 4),
	(7, '222', 1, '2020-06-04 11:34:40', 4),
	(8, '111', 1, '2020-06-04 13:06:58', 50);

CREATE TABLE IF NOT EXISTS `showing` (
  `session_id` int(11) NOT NULL,
  `adv_id` int(11) NOT NULL,
  KEY `FK_showing_sessions` (`session_id`),
  KEY `FK_showing_adv` (`adv_id`),
  CONSTRAINT `FK_showing_adv` FOREIGN KEY (`adv_id`) REFERENCES `adv` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `FK_showing_sessions` FOREIGN KEY (`session_id`) REFERENCES `sessions` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;