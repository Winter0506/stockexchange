package logic

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/tal-tech/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"math/rand"
	"os"
	"stockexchange/rpc/inventory/inventory"
	"stockexchange/rpc/order/internal/svc"
	"stockexchange/rpc/order/model"
	"stockexchange/rpc/stock/stock"
	"time"
)

/*
	Id:                   0,
	User:                 0,
	Stock:                0,
	Number:               0,
	Cost:                 0,
	Direction:            0,
	DealNumber:           0,
	DealCost:             0,
	Status:               "",
	TrustSn:              "",
*/
type TrustListener struct {
	Code   codes.Code
	Detail string
	//Id         int32
	//User       int32
	//Stock      int32
	//Number     int32
	//Cost       float32
	//Direction  uint32
	//DealNumber int32
	//DealCost   float32
	//Status     string
	//TrustSn    string
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

var trustModelId int64

func (t *TrustListener) ExecuteLocalTransaction(message *primitive.Message) primitive.LocalTransactionState {

	var trustModel model.Trust
	_ = json.Unmarshal(message.Body, &trustModel)

	// 跨微服务调用股票微服务
	_, err := t.svcCtx.Stock.GetStockById(t.ctx, &stock.IdRequest{
		Id: trustModel.Stock,
	})

	if err != nil {
		t.Code = codes.Internal
		t.Detail = "查询股票详情失败"
		return primitive.RollbackMessageState
	}

	// 跨服务调用库存微服务进行库存扣减
	if _, err = t.svcCtx.Inventory.Sell(t.ctx, &inventory.SellInfo{StockId: int32(trustModel.Stock), Num: int32(trustModel.Number), TrustSn: trustModel.TrustSn}); err != nil {
		t.Code = codes.ResourceExhausted
		t.Detail = "扣减库存失败"
		return primitive.RollbackMessageState
	}

	// 生成委托表 和 订单表
	rspInsertTrust, err := t.svcCtx.TrustModel.Insert(&model.Trust{
		User:       trustModel.User,
		Stock:      trustModel.Stock,
		Number:     trustModel.Number,
		Cost:       trustModel.Cost,
		Direction:  trustModel.Direction,
		Dealnumber: trustModel.Dealnumber,
		Dealcost:   trustModel.Dealcost,
		Status:     "deal",
		TrustSn:    trustModel.TrustSn,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		DeletedAt:  sql.NullTime{},
		IsDeleted:  0,
	})
	if err != nil {
		t.Code = codes.Internal
		t.Detail = "创建股票委托失败"
		return primitive.CommitMessageState
	}

	trustModelId, _ = rspInsertTrust.LastInsertId()

	_, err = t.svcCtx.OrderModel.Insert(&model.Order{
		User:      trustModel.User,
		Stock:     trustModel.Stock,
		Number:    trustModel.Number,
		Cost:      trustModel.Cost,
		Direction: trustModel.Direction,
		Status:    "TRADE_SUCCESS",
		OrderSn:   trustModel.TrustSn,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: sql.NullTime{},
		IsDeleted: 0,
	})
	if err != nil {
		t.Code = codes.Internal
		t.Detail = "创建股票交易订单失败"
		return primitive.CommitMessageState
	}

	// 更新持仓表
	rspStock, err := t.svcCtx.Stock.GetStockById(t.ctx, &stock.IdRequest{
		Id: trustModel.Stock,
	})
	_, err = t.svcCtx.HoldPositionModel.Insert(&model.Holdposition{
		User:      trustModel.User,
		Stock:     trustModel.Stock,
		StockName: rspStock.StockName,
		Number:    trustModel.Number,
		Cost:      trustModel.Cost,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: sql.NullTime{},
		IsDeleted: 0,
	})
	if err != nil {
		t.Code = codes.Internal
		t.Detail = "创建股票交易订单失败"
		return primitive.CommitMessageState
	}

	rsp, _ := t.svcCtx.UserAccountModel.FindOne(trustModel.User)
	err = t.svcCtx.UserAccountModel.Update(&model.Useraccount{
		Userid:        trustModel.User,
		Account:       rsp.Account,
		MarketValue:   sql.NullFloat64{Float64: rsp.MarketValue.Float64 + trustModel.Dealcost*float64(trustModel.Dealnumber), Valid: true},
		Available:     rsp.Available - trustModel.Dealcost*float64(trustModel.Dealnumber),
		ProfitAndLoss: 0,
		CreatedAt:     rsp.CreatedAt,
		UpdatedAt:     time.Now(),
		DeletedAt:     sql.NullTime{},
		IsDeleted:     0,
	})
	if err != nil {
		t.Code = codes.Internal
		t.Detail = "创建股票交易订单失败"
		return primitive.CommitMessageState
	}

	//发送延时消息
	p, err := rocketmq.NewProducer(producer.WithNameServer([]string{"127.0.0.1:9876"}))
	if err != nil {
		panic("生成producer失败")
	}
	if err = p.Start(); err != nil {
		panic("启动producer失败")
	}

	//
	message = primitive.NewMessage("order_timeout", message.Body)
	message.WithDelayTimeLevel(3)
	_, err = p.SendSync(context.Background(), message)
	if err != nil {
		logx.Errorf("发送延时消息失败: %v\n", err)
		t.Code = codes.Internal
		t.Detail = "发送延时消息失败"
		return primitive.CommitMessageState
	}

	//if err = p.Shutdown(); err != nil {panic("关闭producer失败")}

	t.Code = codes.OK
	return primitive.RollbackMessageState
}

func (t *TrustListener) CheckLocalTransaction(ext *primitive.MessageExt) primitive.LocalTransactionState {
	var trustModel model.Trust
	_ = json.Unmarshal(ext.Body, &trustModel)

	//怎么检查之前的逻辑是否完成
	if result, _ := t.svcCtx.TrustModel.FindTrustSn(trustModel.TrustSn); result == nil {
		return primitive.CommitMessageState
	}

	return primitive.RollbackMessageState
}

func GenerateTrustSn(user int32, stock int32) string {
	//订单号的生成规则
	/*
		年月日时分秒+用户id+2位随机数
	*/
	now := time.Now()
	rand.Seed(time.Now().UnixNano())
	trustSn := fmt.Sprintf("%d%d%d%d%d%d%d%d%d",
		now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Nanosecond(),
		user, stock, rand.Intn(90)+10,
	)
	return trustSn
}

var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root", "123456", "127.0.0.1", 3306, "stockexchange")
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // 禁用彩色打印
		},
	)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
}

func OrderTimeout(ctx context.Context, ext ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
	InitDB()

	for i := range ext {
		var trustInfo model.Trust
		_ = json.Unmarshal(ext[i].Body, &trustInfo)

		fmt.Printf("获取到订单超时消息: %v\n", time.Now())
		//查询委托的状态，如果已支付什么都不做，如果未支付，归还库存

		// 这里trustInfo是可以得到trust的委托号的
		fmt.Println(trustInfo)
		if result := DB.Model(model.Trust{}).Where(model.Trust{TrustSn: trustInfo.TrustSn}).First(&trustInfo); result.RowsAffected == 0 {
			return consumer.ConsumeSuccess, nil
		}

		if trustInfo.Status != "deal" {
			tx := DB.Begin()
			trustInfo.Status = "closed"
			tx.Save(trustInfo)

			p, err := rocketmq.NewProducer(producer.WithNameServer([]string{"127.0.0.1:9876"}))
			if err != nil {
				panic("生成producer失败")
			}

			if err = p.Start(); err != nil {
				panic("启动producer失败")
			}

			_, err = p.SendSync(context.Background(), primitive.NewMessage("order_reback", ext[i].Body))
			if err != nil {
				tx.Rollback()
				fmt.Printf("发送失败: %s\n", err)
				// 稍后重试!
				return consumer.ConsumeRetryLater, nil
			}

			//if err = p.Shutdown(); err != nil {panic("关闭producer失败")}
			return consumer.ConsumeSuccess, nil
		}
	}
	return consumer.ConsumeSuccess, nil
}
