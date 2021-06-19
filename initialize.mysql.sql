/* Initialize the table */
SET sql_notes = 0;
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