// Code generated by goctl. DO NOT EDIT.
// versions:
//  goctl version: 1.7.2

package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	sysAuthFieldNames          = builder.RawFieldNames(&SysAuth{})
	sysAuthRows                = strings.Join(sysAuthFieldNames, ",")
	sysAuthRowsExpectAutoSet   = strings.Join(stringx.Remove(sysAuthFieldNames, "`create_at`", "`create_time`", "`update_at`", "`update_time`"), ",")
	sysAuthRowsWithPlaceHolder = strings.Join(stringx.Remove(sysAuthFieldNames, "`id`", "`create_at`", "`create_time`", "`update_at`", "`update_time`"), "=?,") + "=?"

	cacheSysAuthIdPrefix = "cache:sysAuth:id:"
)

type (
	sysAuthModel interface {
		Insert(ctx context.Context, data *SysAuth) (sql.Result, error)
		FindOne(ctx context.Context, id string) (*SysAuth, error)
		Update(ctx context.Context, data *SysAuth) error
		Delete(ctx context.Context, id string) error
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
		FindList(ctx context.Context, rowBuilder squirrel.SelectBuilder, current, pageSize int64) ([]*SysAuth, error)
		RowBuilder() squirrel.SelectBuilder
		CountBuilder(field string) squirrel.SelectBuilder
		SumBuilder(field string) squirrel.SelectBuilder
		TransCtx(ctx context.Context, fn func(ctx context.Context, sqlx sqlx.Session) error) error
	}

	defaultSysAuthModel struct {
		sqlc.CachedConn
		table string
	}

	SysAuth struct {
		Id          string         `db:"id"`           // 第三方用户ID
		CreatedAt   time.Time      `db:"created_at"`   // 创建时间
		UpdatedAt   sql.NullTime   `db:"updated_at"`   // 更新时间
		DeletedAt   sql.NullTime   `db:"deleted_at"`   // 删除时间
		CreatedName string         `db:"created_name"` // 创建人
		UpdatedName sql.NullString `db:"updated_name"` // 更新人
		DeletedName sql.NullString `db:"deleted_name"` // 删除人
		NickName    string         `db:"nick_name"`    // 机构名
		AuthToken   string         `db:"auth_token"`   // 令牌
		State       int64          `db:"state"`        // 状态 1:正常 2:停用 3:封禁
	}
)

func newSysAuthModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultSysAuthModel {
	return &defaultSysAuthModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`sys_auth`",
	}
}

func (m *defaultSysAuthModel) Delete(ctx context.Context, id string) error {
	sysAuthIdKey := fmt.Sprintf("%s%v", cacheSysAuthIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, sysAuthIdKey)
	return err
}

func (m *defaultSysAuthModel) FindOne(ctx context.Context, id string) (*SysAuth, error) {
	sysAuthIdKey := fmt.Sprintf("%s%v", cacheSysAuthIdPrefix, id)
	var resp SysAuth
	err := m.QueryRowCtx(ctx, &resp, sysAuthIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysAuthRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, sqlx.ErrNotFound
	default:
		return nil, err
	}

}

func (m *defaultSysAuthModel) Insert(ctx context.Context, data *SysAuth) (sql.Result, error) {
	sysAuthIdKey := fmt.Sprintf("%s%v", cacheSysAuthIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, sysAuthRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Id, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.CreatedName, data.UpdatedName, data.DeletedName, data.NickName, data.AuthToken, data.State)
	}, sysAuthIdKey)
	return ret, err
}

func (m *defaultSysAuthModel) Update(ctx context.Context, data *SysAuth) error {
	sysAuthIdKey := fmt.Sprintf("%s%v", cacheSysAuthIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysAuthRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.CreatedName, data.UpdatedName, data.DeletedName, data.NickName, data.AuthToken, data.State, data.Id)
	}, sysAuthIdKey)
	return err
}

func (m *defaultSysAuthModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(sysAuthRows).From(m.table)
}

func (m *defaultSysAuthModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(m.table)
}

func (m *defaultSysAuthModel) SumBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("IFNULL(SUM(" + field + "),0)").From(m.table)
}

func (m *defaultSysAuthModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {

	query, values, err := countBuilder.ToSql()
	if err != nil {
		return 0, err
	}

	var resp int64
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return 0, nil
	default:
		return 0, err
	}
}

func (m *defaultSysAuthModel) FindList(ctx context.Context, rowBuilder squirrel.SelectBuilder, current, pageSize int64) ([]*SysAuth, error) {

	if current < 1 {
		current = 1
	}
	offset := (current - 1) * pageSize

	query, values, err := rowBuilder.Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*SysAuth
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, nil
	default:
		return nil, err
	}
}

func (m *defaultSysAuthModel) TransCtx(ctx context.Context, fn func(ctx context.Context, sqlx sqlx.Session) error) error {
	return m.Transact(func(s sqlx.Session) error {
		return fn(ctx, s)
	})
}

func (m *defaultSysAuthModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheSysAuthIdPrefix, primary)
}

func (m *defaultSysAuthModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysAuthRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultSysAuthModel) tableName() string {
	return m.table
}
