CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` text NOT NULL,
  `age` int(11) DEFAULT '0',
  `hospital_id` int(11) DEFAULT NULL,
  `title` tinytext,
  `department` tinytext,
  `province` tinytext,
  `city` tinytext,
  PRIMARY KEY (`id`)
);
