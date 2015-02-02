CREATE TABLE `friends` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `friend_id` int(11) NOT NULL,
  `status` int(11) NOT NULL,
  `create_time` int(11) NOT NULL,
  `modify_time` int(11) NOT NULL,
  PRIMARY KEY (`id`)
);