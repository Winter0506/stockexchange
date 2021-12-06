CREATE TABLE `user`
(
    `id` int NOT NULL AUTO_INCREMENT COMMENT 'id',
    `username` varchar(255) NOT NULL UNIQUE COMMENT 'username',
    `password` varchar(255) NOT NULL COMMENT 'password',
    PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT = Dynamic;
-- set global innodb_large_prefix=on;
--  set global innodb_file_format=Barracuda;  加上这两句话
-- source user.sql;