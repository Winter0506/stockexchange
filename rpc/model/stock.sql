DROP TABLE IF EXISTS `stock` ;

CREATE TABLE `stock`
(
    `id` int NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `stockname` varchar(20) NOT NULL UNIQUE COMMENT '股票名称',
    `stockcode` varchar(20) NOT NULL UNIQUE COMMENT '股票代码',
    `created_at` datetime NOT NULL COMMENT '创建时间',
    `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
    `isDeleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除, 0否1是',
    PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT = Dynamic;
-- set global innodb_large_prefix=on;
--  set global innodb_file_format=Barracuda;  加上这两句话
-- source stock.sql;

-- 生成model层代码 goctl model mysql ddl -c -src stock.sql -dir .