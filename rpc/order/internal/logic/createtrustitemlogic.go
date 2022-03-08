package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"stockexchange/rpc/order/model"
	"time"

	"stockexchange/rpc/order/internal/svc"
	"stockexchange/rpc/order/order"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/producer"

	"github.com/tal-tech/go-zero/core/logx"
)

type CreateTrustItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateTrustItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTrustItemLogic {
	return &CreateTrustItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  委托相关
func (l *CreateTrustItemLogic) CreateTrustItem(in *order.TrustItemRequest) (*order.TrustInfoResponse, error) {
	/*
			新建委托
				1. 获取委托信息
				2. 写入委托表
				2. 商品的价格自己查询 - 访问商品服务 (跨微服务)
		        3. 等待成交
				3. 库存的扣减 - 访问库存服务 (跨微服务)
				4. 订单的基本信息表
				5. 还没有成交的继续保存 直到成交所有
				6. 每成交一次 都写入订单表作为一个订单项
	*/
	trustListener := TrustListener{
		//Code:       0,
		//Detail:     "",
		//Id:         0,
		//User:       0,
		//Stock:      int32(in.Stock),
		//Number:     0,
		//Cost:       0,
		//Direction:  0,
		//DealNumber: 0,
		//DealCost:   0,
		//Status:     "",
		//TrustSn:    "",
		ctx:    l.ctx,
		svcCtx: l.svcCtx,
	}
	p, err := rocketmq.NewTransactionProducer(
		&trustListener,
		producer.WithNameServer([]string{"127.0.0.1:9876"}),
	)
	if err != nil {
		logx.Errorf("RocketMQ生成producer失败:%s", err.Error())
		return nil, err
	}

	if err = p.Start(); err != nil {
		logx.Errorf("RocketMQ启动producer失败:%s", err.Error())
		return nil, err
	}

	trustModel := model.Trust{
		User:       in.User,
		Stock:      in.Stock,
		Number:     int64(in.Number),
		Cost:       float64(in.Cost),
		Direction:  int64(in.Direction),
		Dealnumber: int64(in.Number),
		Dealcost:   float64(in.Cost),
		Status:     "submitted(已报)", // 这里是已报 但是提交数据库用的是deal  所以超时的时候去查deal
		TrustSn:    GenerateTrustSn(int32(in.User), int32(in.Stock)),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		IsDeleted:  0,
	}

	trustModelString, _ := json.Marshal(trustModel)

	_, err = p.SendMessageInTransaction(context.Background(),
		primitive.NewMessage("trust_reback", trustModelString))
	if err != nil {
		fmt.Printf("RocketMQ发送失败: %s\n", err)
		return nil, status.Error(codes.Internal, "RocketMQ发送消息失败")
	}

	if trustListener.Code != codes.OK {
		return nil, status.Error(trustListener.Code, trustListener.Detail)
	}

	rsp, _ := l.svcCtx.TrustModel.FindOne(trustModelId)

	return &order.TrustInfoResponse{
		Id:         rsp.Id,
		User:       rsp.User,
		Stock:      rsp.Stock,
		Number:     int32(rsp.Number),
		Cost:       float32(rsp.Cost),
		Direction:  uint32(rsp.Direction),
		DealNumber: int32(rsp.Dealnumber),
		DealCost:   float32(rsp.Dealcost),
		Status:     rsp.Status,
		TrustSn:    rsp.TrustSn,
	}, nil
}
