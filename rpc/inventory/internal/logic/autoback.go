package logic

import (
	"context"
	"encoding/json"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/tal-tech/go-zero/core/logx"
	"gorm.io/gorm"
	"stockexchange/rpc/demo/global"
	"stockexchange/rpc/inventory/model"
)

func AutoReback(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
	type TrustInfo struct {
		TrustSn string
	}
	for i := range msgs {
		//既然是归还库存，那么我应该具体的知道每件商品应该归还多少， 但是有一个问题是什么？重复归还的问题
		//所以说这个接口应该确保幂等性， 你不能因为消息的重复发送导致一个订单的库存归还多次， 没有扣减的库存你别归还
		//如果确保这些都没有问题， 新建一张表， 这张表记录了详细的订单扣减细节，以及归还细节
		var trustInfo TrustInfo
		err := json.Unmarshal(msgs[i].Body, &trustInfo)
		if err != nil {
			logx.Errorf("解析json失败： %v\n", msgs[i].Body)
			return consumer.ConsumeSuccess, nil
		}

		//去将inv的库存加回去 将selldetail的status设置为2， 要在事务中进行
		tx := global.DB.Begin()
		var sellDetail model.StockSellDetail
		// 没查询到 直接返回 同时表示消息消费成功
		if result := tx.Model(&model.StockSellDetail{}).Where(&model.StockSellDetail{TrustSn: trustInfo.TrustSn, Status: 1}).First(&sellDetail); result.RowsAffected == 0 {
			return consumer.ConsumeSuccess, nil
		}
		//如果查询到那么归还库存

		//update怎么用
		//先查询一下inventory表在， update语句的 update xx set stocks=stocks+2
		if result := tx.Model(&model.Inventory{}).Where(&model.Inventory{Stock: sellDetail.Detail.Stock}).Update("num", gorm.Expr("total+?", sellDetail.Detail.Num)); result.RowsAffected == 0 {
			tx.Rollback()
			return consumer.ConsumeRetryLater, nil
		}

		if result := tx.Model(&model.StockSellDetail{}).Where(&model.StockSellDetail{TrustSn: sellDetail.TrustSn}).Update("status", 2); result.RowsAffected == 0 {
			tx.Rollback()
			return consumer.ConsumeRetryLater, nil
		}
		tx.Commit()
		return consumer.ConsumeSuccess, nil
	}
	return consumer.ConsumeSuccess, nil
}
