// 交易相关
// 订单相关  用户金额相关表 用户id 本金 现金 总资产 金额 字段 累计盈亏 浮动盈亏

// 持仓表  市值 数量 现价 成本 盈亏 盈亏率  收藏
// 操作记录表  买入 卖出 价格点 日期 成交额

// 库存相关 拓宽思路



```
go run user.go -f etc/user.yaml
go run stock.go -f etc/stock.yaml
go run operation.go -f etc/operation.yaml
go run stockexchange.go -f etc/stockexchange.yaml
```

##### 1.用户相关
搜索用户名 未实现

##### 2.股票相关
修改 更新 删除 股票信息 未实现

##### 3.通知相关
涉及到所收藏股票 买点和卖点的计算及提醒
将来开发 使用python开发通知微服务

##### 4.用户操作相关 
笔记用于记录用户买入卖出的相关操作 这个地方需要用到订单号 获利 买入价 卖出价 日期 等等
所以需要 库存服务 和 订单服务都做完以后再完成 

对于股票的收藏 是需要在这个微服务里面的