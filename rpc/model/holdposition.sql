DROP TABLE IF EXISTS `holdposition` ;

CREATE TABLE `holdposition`
(
    `id` int NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `user` int NOT NULL COMMENT '用户id',
    `stock` int NOT NULL COMMENT '股票id',
    `number` int NOT NULL COMMENT '持仓数量',
    `cost` decimal(10, 4) NOT NULL COMMENT '持仓成本',
    -- 当前市值 当前盈亏 都不用往数据库里面存
    `created_at` datetime NOT NULL COMMENT '创建时间',
    `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
    `isDeleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除, 0否1是',
    PRIMARY KEY(`id`),
    INDEX `user`(`user`) USING BTREE,
    INDEX `stock`(`stock`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT = Dynamic;
-- set global innodb_large_prefix=on;
--  set global innodb_file_format=Barracuda;  加上这两句话
-- source holdposition.sql;