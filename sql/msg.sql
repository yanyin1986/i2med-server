CREATE TABLE `messages` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `message` text NOT NULL,
  `commit_time` int(11) NOT NULL,
  PRIMARY KEY (`id`)
);