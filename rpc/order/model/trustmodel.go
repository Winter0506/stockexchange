package model

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/tal-tech/go-zero/core/stores/builder"
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
)

var (
	trustFieldNames          = builder.RawFieldNames(&Trust{})
	trustRows                = strings.Join(trustFieldNames, ",")
	trustRowsExpectAutoSet   = strings.Join(stringx.Remove(trustFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	trustRowsWithPlaceHolder = strings.Join(stringx.Remove(trustFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheTrustIdPrefix = "cache:trust:id:"
)

type (
	TrustModel interface {
		Insert(data *Trust) (sql.Result, error)
		FindOne(id int64) (*Trust, error)
		Update(data *Trust) error
		Delete(id int64) error
		FindTrustByUser(user int64) (*[]Trust, error)
		FindTrustSn(trustSn string) (*Trust, error)
	}

	defaultTrustModel struct {
		sqlc.CachedConn
		table string
	}

	Trust struct {
		Id         int64        `db:"id" gorm:"type:int64"`                           // ID
		User       int64        `db:"user" gorm:"type:int64"`                         // 用户id
		Stock      int64        `db:"stock" gorm:"type:int64"`                        // 股票id
		Number     int64        `db:"number" gorm:"type:int64"`                       // 委托数量
		Cost       float64      `db:"cost" gorm:"type:float64"`                       // 委托成本
		Direction  int64        `db:"direction" gorm:"type:int64"`                    // 1买入, 2表示卖出
		Dealnumber int64        `db:"dealnumber" gorm:"type:int64"`                   // 成交数量
		Dealcost   float64      `db:"dealcost" gorm:"type:float64"`                   // 成交均价
		Status     string       `db:"status" gorm:"type:varchar(20)"`                 // submitted(已报), deal(成交), partial(部分成交), undo(撤销), closed(超时关闭)
		TrustSn    string       `db:"trustSn" gorm:"type:varchar(30);column:trustSn"` // 委托单号
		CreatedAt  time.Time    `db:"created_at"`                                     // 创建时间
		UpdatedAt  time.Time    `db:"updated_at"`                                     // 更新时间
		DeletedAt  sql.NullTime `db:"deleted_at"`                                     // 删除时间
		IsDeleted  int64        `db:"isDeleted"`                                      // 是否删除, 0否1是
	}
)

func (m *defaultTrustModel) FindTrustSn(trustSn string) (*Trust, error) {
	var resp Trust
	err := m.QueryRow(&resp, trustSn, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `trustSn` = ? limit 1", trustRows, m.table)
		if err := conn.QueryRow(&resp, query, trustSn); err != nil {
			return err
		}
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

func NewTrustModel(conn sqlx.SqlConn, c cache.CacheConf) TrustModel {
	return &defaultTrustModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`trust`",
	}
}

func (m *defaultTrustModel) FindTrustByUser(user int64) (*[]Trust, error) {
	userIdKey := fmt.Sprintf("%d", user)
	var resp []Trust
	err := m.QueryRow(&resp, userIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `user` = ? limit 1", trustRows, m.table)
		if err := conn.QueryRow(&resp, query, userIdKey); err != nil {
			return err
		}
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

func (m *defaultTrustModel) Insert(data *Trust) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, trustRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.User, data.Stock, data.Number, data.Cost, data.Direction, data.Dealnumber, data.Dealcost, data.Status, data.TrustSn, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.IsDeleted)

	return ret, err
}

func (m *defaultTrustModel) FindOne(id int64) (*Trust, error) {
	trustIdKey := fmt.Sprintf("%s%v", cacheTrustIdPrefix, id)
	var resp Trust
	err := m.QueryRow(&resp, trustIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", trustRows, m.table)
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

func (m *defaultTrustModel) Update(data *Trust) error {
	trustIdKey := fmt.Sprintf("%s%v", cacheTrustIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, trustRowsWithPlaceHolder)
		return conn.Exec(query, data.User, data.Stock, data.Number, data.Cost, data.Direction, data.Dealnumber, data.Dealcost, data.Status, data.TrustSn, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.IsDeleted, data.Id)
	}, trustIdKey)
	return err
}

func (m *defaultTrustModel) Delete(id int64) error {

	trustIdKey := fmt.Sprintf("%s%v", cacheTrustIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, trustIdKey)
	return err
}

func (m *defaultTrustModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheTrustIdPrefix, primary)
}

func (m *defaultTrustModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", trustRows, m.table)
	return conn.QueryRow(v, query, primary)
}
