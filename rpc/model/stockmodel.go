package model

import (
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	stockFieldNames          = builder.RawFieldNames(&Stock{})
	stockRows                = strings.Join(stockFieldNames, ",")
	stockRowsExpectAutoSet   = strings.Join(stringx.Remove(stockFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	stockRowsWithPlaceHolder = strings.Join(stringx.Remove(stockFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheStockIdPrefix        = "cache:stock:id:"
	cacheStockStockcodePrefix = "cache:stock:stockcode:"
	cacheStockStocknamePrefix = "cache:stock:stockname:"
)

type (
	StockModel interface {
		Insert(data *Stock) (sql.Result, error)
		FindOne(id int64) (*Stock, error)
		FindOneByStockcode(stockcode string) (*Stock, error)
		FindOneByStockname(stockname string) (*Stock, error)
		Update(data *Stock) error
		Delete(id int64) error
		// 新增接口
		FindAll() (*[]Stock, error)
	}

	defaultStockModel struct {
		sqlc.CachedConn
		table string
	}

	Stock struct {
		Id        int32        `db:"id"`         // ID
		Stockname string       `db:"stockname"`  // 股票名称
		Stockcode string       `db:"stockcode"`  // 股票代码
		CreatedAt time.Time    `db:"created_at"` // 创建时间  // 为什么和user 时间格式不一样
		UpdatedAt time.Time    `db:"updated_at"` // 更新时间
		DeletedAt sql.NullTime `db:"deleted_at"` // 删除时间
		IsDeleted int64        `db:"isDeleted"`  // 是否删除, 0否1是
	}
)

// 查询所有stock
func (m *defaultStockModel) FindAll() (*[]Stock, error) {
	var resp []Stock
	Key := fmt.Sprintf("%s", cacheStockIdPrefix)
	err := m.QueryRow(&resp, Key, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s", stockRows, m.table)
		return conn.QueryRows(v, query)
	})
	fmt.Println(err)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func NewStockModel(conn sqlx.SqlConn, c cache.CacheConf) StockModel {
	return &defaultStockModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`stock`",
	}
}

func (m *defaultStockModel) Insert(data *Stock) (sql.Result, error) {
	stockIdKey := fmt.Sprintf("%s%v", cacheStockIdPrefix, data.Id)
	stockStockcodeKey := fmt.Sprintf("%s%v", cacheStockStockcodePrefix, data.Stockcode)
	stockStocknameKey := fmt.Sprintf("%s%v", cacheStockStocknamePrefix, data.Stockname)
	ret, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, stockRowsExpectAutoSet)
		return conn.Exec(query, data.Stockname, data.Stockcode, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.IsDeleted)
	}, stockIdKey, stockStockcodeKey, stockStocknameKey)
	return ret, err
}

func (m *defaultStockModel) FindOne(id int64) (*Stock, error) {
	stockIdKey := fmt.Sprintf("%s%v", cacheStockIdPrefix, id)
	var resp Stock
	err := m.QueryRow(&resp, stockIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", stockRows, m.table)
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

func (m *defaultStockModel) FindOneByStockcode(stockcode string) (*Stock, error) {
	stockStockcodeKey := fmt.Sprintf("%s%v", cacheStockStockcodePrefix, stockcode)
	var resp Stock
	err := m.QueryRowIndex(&resp, stockStockcodeKey, m.formatPrimary, func(conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `stockcode` = ? limit 1", stockRows, m.table)
		if err := conn.QueryRow(&resp, query, stockcode); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultStockModel) FindOneByStockname(stockname string) (*Stock, error) {
	stockStocknameKey := fmt.Sprintf("%s%v", cacheStockStocknamePrefix, stockname)
	var resp Stock
	err := m.QueryRowIndex(&resp, stockStocknameKey, m.formatPrimary, func(conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `stockname` = ? limit 1", stockRows, m.table)
		if err := conn.QueryRow(&resp, query, stockname); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultStockModel) Update(data *Stock) error {
	stockIdKey := fmt.Sprintf("%s%v", cacheStockIdPrefix, data.Id)
	stockStockcodeKey := fmt.Sprintf("%s%v", cacheStockStockcodePrefix, data.Stockcode)
	stockStocknameKey := fmt.Sprintf("%s%v", cacheStockStocknamePrefix, data.Stockname)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, stockRowsWithPlaceHolder)
		return conn.Exec(query, data.Stockname, data.Stockcode, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.IsDeleted, data.Id)
	}, stockIdKey, stockStockcodeKey, stockStocknameKey)
	return err
}

func (m *defaultStockModel) Delete(id int64) error {
	data, err := m.FindOne(id)
	if err != nil {
		return err
	}

	stockIdKey := fmt.Sprintf("%s%v", cacheStockIdPrefix, id)
	stockStockcodeKey := fmt.Sprintf("%s%v", cacheStockStockcodePrefix, data.Stockcode)
	stockStocknameKey := fmt.Sprintf("%s%v", cacheStockStocknamePrefix, data.Stockname)
	_, err = m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, stockIdKey, stockStockcodeKey, stockStocknameKey)
	return err
}

func (m *defaultStockModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheStockIdPrefix, primary)
}

func (m *defaultStockModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", stockRows, m.table)
	return conn.QueryRow(v, query, primary)
}
