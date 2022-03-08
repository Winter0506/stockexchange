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

DROP TABLE IF EXISTS `holdposition` ;

CREATE TABLE `holdposition`
(
    `id` int NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `user` int NOT NULL COMMENT '用户id',
    `stock` int NOT NULL COMMENT '股票id',
    `stockName` int NOT NULL COMMENT '股票名',
    `number` int NOT NULL COMMENT '持仓数量',
    `cost` decimal(10, 4) NOT NULL COMMENT '持仓成本',
    -- 当前市值 当前盈亏 都不用往数据库里面存
    `created_at` datetime NOT NULL COMMENT '创建时间',
    `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
    `isDeleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除, 0否1是',
    PRIMARY KEY(`id`),
    INDEX `user`(`user`) USING BTREE,
    INDEX `stock`(`stock`) USING BTREE,
    INDEX `stockName`(`stockName`) USING BTREE
    -- KEY `user_stock` (`user`,`stock`) -- 这个地方应该添加联合索引, 先不添加走走逻辑试一下
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT = Dynamic;
-- set global innodb_large_prefix=on;
--  set global innodb_file_format=Barracuda;  加上这两句话
-- source holdposition.sql;

DROP TABLE IF EXISTS `trust` ;

CREATE TABLE `trust`
(
    `id` int NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `user` int NOT NULL COMMENT '用户id',
    `stock` int NOT NULL COMMENT '股票id',
    `number` int NOT NULL COMMENT '委托数量',
    `cost` decimal(10, 4) NOT NULL COMMENT '委托成本',
    `direction` tinyint(1) NOT NULL COMMENT '1买入, 2表示卖出',
    `dealnumber` int NOT NULL COMMENT '成交数量',
    `dealcost` decimal(10, 4) NOT NULL COMMENT '成交均价',
    `status` varchar(20) NOT NULL COMMENT 'submitted(已报), deal(成交), partial(部分成交), undo(撤销), closed(超时关闭)',
    `trustSn` varchar(30) NOT NULL COMMENT '委托单号',
    `created_at` datetime NOT NULL COMMENT '创建时间',
    `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
    `isDeleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除, 0否1是',
    PRIMARY KEY(`id`),
    INDEX `user`(`user`) USING BTREE,
    INDEX `stock`(`stock`) USING BTREE,
    INDEX `trustSn`(`trustSn`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT = Dynamic;
-- set global innodb_large_prefix=on;
--  set global innodb_file_format=Barracuda;  加上这两句话
-- source stock.sql;

-- 生成model层代码 goctl model mysql ddl -c -src stock.sql -dir .
-- 还应该有成交数量 成交金额 成交均价 这里省略 否则逻辑过于复杂

DROP TABLE IF EXISTS `order` ;

CREATE TABLE `order`
(
    `id` int NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `user` int NOT NULL COMMENT '用户id',
    `stock` int NOT NULL COMMENT '股票id',
    `number` int NOT NULL COMMENT '订单数量',
    `cost` decimal(10, 4) NOT NULL COMMENT '订单成本',
    `direction` tinyint(1) NOT NULL COMMENT '1买入, 2表示卖出',
    -- 需要和库存服务交互的
    `status` varchar(20) NOT NULL COMMENT 'PAYING(待支付), TRADE_SUCCESS(成功), TRADE_CLOSED(超时关闭), WAIT_BUYER_PAY(交易创建), TRADE_FINISHED(交易结束)',
    `orderSn` varchar(30) NOT NULL COMMENT '订单单号',
    `created_at` datetime NOT NULL COMMENT '创建时间',
    `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
    `isDeleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除, 0否1是',
    PRIMARY KEY(`id`),
    INDEX `user`(`user`) USING BTREE,
    INDEX `stock`(`stock`) USING BTREE,
    INDEX `orderSn`(`orderSn`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT = Dynamic;
-- set global innodb_large_prefix=on;
--  set global innodb_file_format=Barracuda;  加上这两句话
-- source stock.sql;

-- 生成model层代码 goctl model mysql ddl -c -src order.sql -dir .