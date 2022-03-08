package model

import (
	"database/sql"
	"fmt"
	"github.com/tal-tech/go-zero/core/stores/builder"
	"strings"
	"time"

	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
)

var (
	orderFieldNames          = builder.RawFieldNames(&Order{})
	orderRows                = strings.Join(orderFieldNames, ",")
	orderRowsExpectAutoSet   = strings.Join(stringx.Remove(orderFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	orderRowsWithPlaceHolder = strings.Join(stringx.Remove(orderFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheOrderIdPrefix = "cache:order:id:"
)

type (
	OrderModel interface {
		Insert(data *Order) (sql.Result, error)
		FindOne(id int64) (*Order, error)
		Update(data *Order) error
		Delete(id int64) error
	}

	defaultOrderModel struct {
		sqlc.CachedConn
		table string
	}

	Order struct {
		Id        int64        `db:"id"`         // ID
		User      int64        `db:"user"`       // 用户id
		Stock     int64        `db:"stock"`      // 股票id
		Number    int64        `db:"number"`     // 订单数量
		Cost      float64      `db:"cost"`       // 订单成本
		Direction int64        `db:"direction"`  // 1买入, 2表示卖出
		Status    string       `db:"status"`     // PAYING(待支付), TRADE_SUCCESS(成功), TRADE_CLOSED(超时关闭), WAIT_BUYER_PAY(交易创建), TRADE_FINISHED(交易结束)
		OrderSn   string       `db:"orderSn"`    // 订单单号
		CreatedAt time.Time    `db:"created_at"` // 创建时间
		UpdatedAt time.Time    `db:"updated_at"` // 更新时间
		DeletedAt sql.NullTime `db:"deleted_at"` // 删除时间
		IsDeleted int64        `db:"isDeleted"`  // 是否删除, 0否1是
	}
)

func NewOrderModel(conn sqlx.SqlConn, c cache.CacheConf) OrderModel {
	return &defaultOrderModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`order`",
	}
}

func (m *defaultOrderModel) Insert(data *Order) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, orderRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.User, data.Stock, data.Number, data.Cost, data.Direction, data.Status, data.OrderSn, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.IsDeleted)

	return ret, err
}

func (m *defaultOrderModel) FindOne(id int64) (*Order, error) {
	orderIdKey := fmt.Sprintf("%s%v", cacheOrderIdPrefix, id)
	var resp Order
	err := m.QueryRow(&resp, orderIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", orderRows, m.table)
		return conn.QueryRow(v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultOrderModel) Update(data *Order) error {
	orderIdKey := fmt.Sprintf("%s%v", cacheOrderIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, orderRowsWithPlaceHolder)
		return conn.Exec(query, data.User, data.Stock, data.Number, data.Cost, data.Direction, data.Status, data.OrderSn, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.IsDeleted, data.Id)
	}, orderIdKey)
	return err
}

func (m *defaultOrderModel) Delete(id int64) error {

	orderIdKey := fmt.Sprintf("%s%v", cacheOrderIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, orderIdKey)
	return err
}

func (m *defaultOrderModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheOrderIdPrefix, primary)
}

func (m *defaultOrderModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", orderRows, m.table)
	return conn.QueryRow(v, query, primary)
}
