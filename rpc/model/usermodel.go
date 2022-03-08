package model

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder" // builderx 要改成build
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userFieldNames          = builder.RawFieldNames(&User{})
	userRows                = strings.Join(userFieldNames, ",")
	userRowsExpectAutoSet   = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	userRowsWithPlaceHolder = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheUserIdPrefix       = "cache:user:id:"
	cacheUserEmailPrefix    = "cache:user:email:"
	cacheUserUsernamePrefix = "cache:user:username:"
)

type (
	UserModel interface {
		Insert(data *User) (sql.Result, error)
		FindOne(id int64) (*User, error)
		FindOneByEmail(email string) (*User, error)
		FindOneByUsername(username string) (*User, error)
		Update(data *User) error
		Delete(id int64) error
		// 新增接口
		FindAll() (*[]User, error)
	}

	defaultUserModel struct {
		sqlc.CachedConn
		table string
	}

	User struct {
		Id        int64        `db:"id"`         // ID
		Username  string       `db:"username"`   // 用户名
		Password  string       `db:"password"`   // 密码
		Email     string       `db:"email"`      // 邮件
		Gender    string       `db:"gender"`     // female表示女, male表示男
		Role      int64        `db:"role"`       // 1表示普通用户, 2表示管理员
		CreatedAt sql.NullTime `db:"created_at"` // 创建时间
		UpdatedAt sql.NullTime `db:"updated_at"` // 更新时间
		DeletedAt sql.NullTime `db:"deleted_at"` // 删除时间
		IsDeleted int64        `db:"isDeleted"`  // 是否删除, 0否1是
	}
)

func NewUserModel(conn sqlx.SqlConn, c cache.CacheConf) UserModel {
	return &defaultUserModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`user`",
	}
}

// 这是查询所有
func (m *defaultUserModel) FindAll() (*[]User, error) {
	// 写成*[]User不行 会发生内存错误
	var resp []User
	Key := fmt.Sprintf("%s", cacheUserIdPrefix)
	err := m.QueryRow(&resp, Key, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s", userRows, m.table)
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

func (m *defaultUserModel) Insert(data *User) (sql.Result, error) {
	userEmailKey := fmt.Sprintf("%s%v", cacheUserEmailPrefix, data.Email)
	userUsernameKey := fmt.Sprintf("%s%v", cacheUserUsernamePrefix, data.Username)
	userIdKey := fmt.Sprintf("%s%v", cacheUserIdPrefix, data.Id)
	ret, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, userRowsExpectAutoSet)
		fmt.Println(query)
		return conn.Exec(query, data.Username, data.Password, data.Email, data.Gender, data.Role, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.IsDeleted)
	}, userIdKey, userEmailKey, userUsernameKey)
	return ret, err
}

func (m *defaultUserModel) FindOne(id int64) (*User, error) {
	userIdKey := fmt.Sprintf("%s%v", cacheUserIdPrefix, id)
	var resp User
	err := m.QueryRow(&resp, userIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userRows, m.table)
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

func (m *defaultUserModel) FindOneByEmail(email string) (*User, error) {
	userEmailKey := fmt.Sprintf("%s%v", cacheUserEmailPrefix, email)
	var resp User
	err := m.QueryRowIndex(&resp, userEmailKey, m.formatPrimary, func(conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `email` = ? limit 1", userRows, m.table)
		if err := conn.QueryRow(&resp, query, email); err != nil {
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

func (m *defaultUserModel) FindOneByUsername(username string) (*User, error) {
	userUsernameKey := fmt.Sprintf("%s%v", cacheUserUsernamePrefix, username)
	var resp User
	err := m.QueryRowIndex(&resp, userUsernameKey, m.formatPrimary, func(conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `username` = ? limit 1", userRows, m.table)
		if err := conn.QueryRow(&resp, query, username); err != nil {
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

func (m *defaultUserModel) Update(data *User) error {
	userIdKey := fmt.Sprintf("%s%v", cacheUserIdPrefix, data.Id)
	userEmailKey := fmt.Sprintf("%s%v", cacheUserEmailPrefix, data.Email)
	userUsernameKey := fmt.Sprintf("%s%v", cacheUserUsernamePrefix, data.Username)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userRowsWithPlaceHolder)
		return conn.Exec(query, data.Username, data.Password, data.Email, data.Gender, data.Role, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.IsDeleted, data.Id)
	}, userIdKey, userEmailKey, userUsernameKey)
	return err
}

// 普通业务应该为软删除 管理员才可以调用Delete方法
func (m *defaultUserModel) Delete(id int64) error {
	data, err := m.FindOne(id)
	if err != nil {
		return err
	}

	userIdKey := fmt.Sprintf("%s%v", cacheUserIdPrefix, id)
	userEmailKey := fmt.Sprintf("%s%v", cacheUserEmailPrefix, data.Email)
	userUsernameKey := fmt.Sprintf("%s%v", cacheUserUsernamePrefix, data.Username)
	_, err = m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, userUsernameKey, userIdKey, userEmailKey)
	return err
}

func (m *defaultUserModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheUserIdPrefix, primary)
}

func (m *defaultUserModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userRows, m.table)
	return conn.QueryRow(v, query, primary)
}
