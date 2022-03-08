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
	useraccountFieldNames          = builder.RawFieldNames(&Useraccount{})
	useraccountRows                = strings.Join(useraccountFieldNames, ",")
	useraccountRowsExpectAutoSet   = strings.Join(stringx.Remove(useraccountFieldNames, "`userid`", "`create_time`", "`update_time`"), ",")
	useraccountRowsWithPlaceHolder = strings.Join(stringx.Remove(useraccountFieldNames, "`userid`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheUseraccountUseridPrefix = "cache:useraccount:userid:"
)

type (
	UseraccountModel interface {
		Insert(data *Useraccount) (sql.Result, error)
		FindOne(userid int64) (*Useraccount, error)
		Update(data *Useraccount) error
		Delete(userid int64) error
	}

	defaultUseraccountModel struct {
		sqlc.CachedConn
		table string
	}

	Useraccount struct {
		Userid        int64           `db:"userid"`        // 用户ID
		Account       float64         `db:"account"`       // 用户钱包
		MarketValue   sql.NullFloat64 `db:"marketValue"`   // 持股市值
		Available     float64         `db:"available"`     // 可用金钱
		ProfitAndLoss float64         `db:"profitAndLoss"` // 盈亏
		CreatedAt     time.Time       `db:"created_at"`    // 创建时间
		UpdatedAt     time.Time       `db:"updated_at"`    // 更新时间
		DeletedAt     sql.NullTime    `db:"deleted_at"`    // 删除时间
		IsDeleted     int64           `db:"isDeleted"`     // 是否删除, 0否1是
	}
)

func NewUseraccountModel(conn sqlx.SqlConn, c cache.CacheConf) UseraccountModel {
	return &defaultUseraccountModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`useraccount`",
	}
}

func (m *defaultUseraccountModel) Insert(data *Useraccount) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, useraccountRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Account, data.MarketValue, data.Available, data.ProfitAndLoss, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.IsDeleted)

	return ret, err
}

func (m *defaultUseraccountModel) FindOne(userid int64) (*Useraccount, error) {
	useraccountUseridKey := fmt.Sprintf("%s%v", cacheUseraccountUseridPrefix, userid)
	var resp Useraccount
	err := m.QueryRow(&resp, useraccountUseridKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `userid` = ? limit 1", useraccountRows, m.table)
		return conn.QueryRow(v, query, userid)
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

func (m *defaultUseraccountModel) Update(data *Useraccount) error {
	useraccountUseridKey := fmt.Sprintf("%s%v", cacheUseraccountUseridPrefix, data.Userid)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `userid` = ?", m.table, useraccountRowsWithPlaceHolder)
		return conn.Exec(query, data.Account, data.MarketValue, data.Available, data.ProfitAndLoss, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.IsDeleted, data.Userid)
	}, useraccountUseridKey)
	return err
}

func (m *defaultUseraccountModel) Delete(userid int64) error {

	useraccountUseridKey := fmt.Sprintf("%s%v", cacheUseraccountUseridPrefix, userid)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `userid` = ?", m.table)
		return conn.Exec(query, userid)
	}, useraccountUseridKey)
	return err
}

func (m *defaultUseraccountModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheUseraccountUseridPrefix, primary)
}

func (m *defaultUseraccountModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `userid` = ? limit 1", useraccountRows, m.table)
	return conn.QueryRow(v, query, primary)
}
