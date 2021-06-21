/* Initialize the table */
SET sql_notes = 0;

DROP TABLE IF EXISTS `gofixtures_test`.`coupons`;
CREATE TABLE IF NOT EXISTS `gofixtures_test`.`coupons` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `coupon_id` int(11) NOT NULL,
  `coupon_name` varchar(255) NOT NULL DEFAULT '',
  `name` varchar(255) NOT NULL DEFAULT '',
  `amount` int(10) unsigned NOT NULL,
  `order_limit` int(2) unsigned NOT NULL,
  `note` varchar(1000) NOT NULL DEFAULT '',
  `create_time` int(11) unsigned NOT NULL,
  `update_time` int(11) unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `gofixtures_test`.`admin_users`;
CREATE TABLE IF NOT EXISTS `gofixtures_test`.`admin_users` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `email` varchar(255) NOT NULL DEFAULT '',
  `nickname` varchar(255) NOT NULL DEFAULT '',
  `premissions` int(10) unsigned NOT NULL,
  `role` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `gofixtures_test`.`users`;
CREATE TABLE IF NOT EXISTS `gofixtures_test`.`users` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `nickname` varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `ctime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `mtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- create initial rows
INSERT INTO `gofixtures_test`.`users` (`nickname`) VALUES ("john"), ("mike"), ("tiny");

SET sql_notes = 1;