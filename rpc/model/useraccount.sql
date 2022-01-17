DROP TABLE IF EXISTS `useraccount` ;

CREATE TABLE `useraccount`
(
    `userid` int NOT NULL AUTO_INCREMENT COMMENT '用户ID',
    `account` decimal(10, 2) NOT NULL COMMENT '用户钱包',
    `marketValue` decimal(10, 2) DEFAULT NULL COMMENT '持股市值',
    `available` decimal(10, 2) NOT NULL COMMENT '可用金钱',
    `profitAndLoss` decimal(10, 2) NOT NULL COMMENT '盈亏',
    -- 总金日收益 与 股票日收益 计算方法
    -- 总金等于所有股票日相加  -- 股票日收益等于所有当前的价 减去开盘价
    -- 所以不用去可以新建一个收益表
    `created_at` datetime NOT NULL COMMENT '创建时间',
    `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    -- 用户可以更新自己的钱包金额
    `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
    -- 用户可以把自己的账户删除
    `isDeleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除, 0否1是',
    PRIMARY KEY(`userid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT = Dynamic;
-- set global innodb_large_prefix=on;
--  set global innodb_file_format=Barracuda;  加上这两句话
-- source useraccount.sql;

-- goctl model mysql ddl -c -src useraccount.sql -dir .