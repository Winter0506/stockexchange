CREATE TABLE `user`
(
    `id` int NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `username` varchar(20) NOT NULL UNIQUE COMMENT '用户名',
    `password` varchar(100) NOT NULL COMMENT '密码',
    `email` varchar(50) NOT NULL COMMENT '邮件',
    `gender` varchar(6) DEFAULT 'male' COMMENT 'female表示女, male表示男',
    `role` tinyint(1) DEFAULT '1' COMMENT '1表示普通用户, 2表示管理员',
    `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
    `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
    `deleted_at` datetime(0) NULL DEFAULT NULL COMMENT '删除时间',
    `isDeleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除, 0否1是',
    PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT = Dynamic;
-- set global innodb_large_prefix=on;
--  set global innodb_file_format=Barracuda;  加上这两句话
-- source user.sql;
