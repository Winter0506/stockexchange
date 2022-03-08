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
	holdpositionFieldNames          = builder.RawFieldNames(&Holdposition{})
	holdpositionRows                = strings.Join(holdpositionFieldNames, ",")
	holdpositionRowsExpectAutoSet   = strings.Join(stringx.Remove(holdpositionFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	holdpositionRowsWithPlaceHolder = strings.Join(stringx.Remove(holdpositionFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheHoldpositionIdPrefix     = "cache:holdposition:id:"
	cacheHoldpositionUseridPrefix = "cache:holdposition:userid:"
)

type (
	HoldpositionModel interface {
		Insert(data *Holdposition) (sql.Result, error)
		FindOne(id int64) (*Holdposition, error)
		Update(data *Holdposition) error
		Delete(id int64) error
		FindHoldPositionByUser(user int64) (*[]Holdposition, error)
	}

	defaultHoldpositionModel struct {
		sqlc.CachedConn
		table string
	}

	Holdposition struct {
		Id        int64        `db:"id"`         // ID
		User      int64        `db:"user"`       // 用户id
		Stock     int64        `db:"stock"`      // 股票id
		StockName string       `db:"stockName"`  // 股票名
		Number    int64        `db:"number"`     // 持仓数量
		Cost      float64      `db:"cost"`       // 持仓成本
		CreatedAt time.Time    `db:"created_at"` // 创建时间
		UpdatedAt time.Time    `db:"updated_at"` // 更新时间
		DeletedAt sql.NullTime `db:"deleted_at"` // 删除时间
		IsDeleted int64        `db:"isDeleted"`  // 是否删除, 0否1是
	}
)

func NewHoldpositionModel(conn sqlx.SqlConn, c cache.CacheConf) HoldpositionModel {
	return &defaultHoldpositionModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`holdposition`",
	}
}

//{"@timestamp":"2022-02-24T22:28:03.763+08","level":"info","duration":"1.0ms","content":"sql query: select `id`,`user`,
//`stock`,`stockName`,`number`,`cost`,`created_
//at`,`updated_at`,`deleted_at`,`isDeleted` from `holdposition` where `user` = '1'"}
//unsupported unmarshal type

func (m *defaultHoldpositionModel) FindHoldPositionByUser(user int64) (*[]Holdposition, error) {
	userIdKey := fmt.Sprintf("%d", user)
	var resp []Holdposition
	err := m.QueryRowIndex(&resp, userIdKey, m.formatPrimary, func(conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `user` = ?", holdpositionRows, m.table)
		if err := conn.QueryRow(&resp, query, user); err != nil {
			return nil, err
		}
		return &resp, nil
	}, m.queryPrimary)
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

func (m *defaultHoldpositionModel) Insert(data *Holdposition) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, holdpositionRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.User, data.Stock, data.StockName, data.Number, data.Cost, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.IsDeleted)

	return ret, err
}

func (m *defaultHoldpositionModel) FindOne(id int64) (*Holdposition, error) {
	holdpositionIdKey := fmt.Sprintf("%s%v", cacheHoldpositionIdPrefix, id)
	var resp Holdposition
	err := m.QueryRow(&resp, holdpositionIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", holdpositionRows, m.table)
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

func (m *defaultHoldpositionModel) Update(data *Holdposition) error {
	holdpositionIdKey := fmt.Sprintf("%s%v", cacheHoldpositionIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, holdpositionRowsWithPlaceHolder)
		return conn.Exec(query, data.User, data.Stock, data.StockName, data.Number, data.Cost, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.IsDeleted, data.Id)
	}, holdpositionIdKey)
	return err
}

func (m *defaultHoldpositionModel) Delete(id int64) error {

	holdpositionIdKey := fmt.Sprintf("%s%v", cacheHoldpositionIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, holdpositionIdKey)
	return err
}

func (m *defaultHoldpositionModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheHoldpositionIdPrefix, primary)
}

func (m *defaultHoldpositionModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", holdpositionRows, m.table)
	return conn.QueryRow(v, query, primary)
}
