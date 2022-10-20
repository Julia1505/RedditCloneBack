SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
                         `id` varchar(200) NOT NULL,
                         `username` varchar(200) NOT NULL,
                         `passwordhash` varchar(200) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
