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
	tpmtAssetFieldNames          = builder.RawFieldNames(&TpmtAsset{})
	tpmtAssetRows                = strings.Join(tpmtAssetFieldNames, ",")
	tpmtAssetRowsExpectAutoSet   = strings.Join(stringx.Remove(tpmtAssetFieldNames, "`create_at`", "`create_time`", "`update_at`", "`update_time`"), ",")
	tpmtAssetRowsWithPlaceHolder = strings.Join(stringx.Remove(tpmtAssetFieldNames, "`id`", "`create_at`", "`create_time`", "`update_at`", "`update_time`"), "=?,") + "=?"

	cacheTpmtAssetIdPrefix = "cache:tpmtAsset:id:"
)

type (
	tpmtAssetModel interface {
		Insert(ctx context.Context, data *TpmtAsset) (sql.Result, error)
		FindOne(ctx context.Context, id string) (*TpmtAsset, error)
		Update(ctx context.Context, data *TpmtAsset) error
		Delete(ctx context.Context, id string) error
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
		FindList(ctx context.Context, rowBuilder squirrel.SelectBuilder, current, pageSize int64) ([]*TpmtAsset, error)
		RowBuilder() squirrel.SelectBuilder
		CountBuilder(field string) squirrel.SelectBuilder
		SumBuilder(field string) squirrel.SelectBuilder
		TransCtx(ctx context.Context, fn func(ctx context.Context, sqlx sqlx.Session) error) error
	}

	defaultTpmtAssetModel struct {
		sqlc.CachedConn
		table string
	}

	TpmtAsset struct {
		Id           string         `db:"id"`            // 资产ID
		AssetType    int64          `db:"asset_type"`    // 资产类型
		AssetCode    string         `db:"asset_code"`    // 资产编号
		AssetName    string         `db:"asset_name"`    // 资产名称
		AssetModel   string         `db:"asset_model"`   // 资产型号
		ManuFacturer string         `db:"manu_facturer"` // 生产厂家
		Voltage      string         `db:"voltage"`       // 电压
		Capacity     string         `db:"capacity"`      // 容量
		CreatedAt    time.Time      `db:"created_at"`    // 创建时间
		CreatedName  string         `db:"created_name"`  // 创建人
		UpdatedAt    sql.NullTime   `db:"updated_at"`    // 更新时间
		UpdatedName  sql.NullString `db:"updated_name"`  // 更新人
	}
)

func newTpmtAssetModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultTpmtAssetModel {
	return &defaultTpmtAssetModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`tpmt_asset`",
	}
}

func (m *defaultTpmtAssetModel) Delete(ctx context.Context, id string) error {
	tpmtAssetIdKey := fmt.Sprintf("%s%v", cacheTpmtAssetIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, tpmtAssetIdKey)
	return err
}

func (m *defaultTpmtAssetModel) FindOne(ctx context.Context, id string) (*TpmtAsset, error) {
	tpmtAssetIdKey := fmt.Sprintf("%s%v", cacheTpmtAssetIdPrefix, id)
	var resp TpmtAsset
	err := m.QueryRowCtx(ctx, &resp, tpmtAssetIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", tpmtAssetRows, m.table)
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

func (m *defaultTpmtAssetModel) Insert(ctx context.Context, data *TpmtAsset) (sql.Result, error) {
	tpmtAssetIdKey := fmt.Sprintf("%s%v", cacheTpmtAssetIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, tpmtAssetRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Id, data.AssetType, data.AssetCode, data.AssetName, data.AssetModel, data.ManuFacturer, data.Voltage, data.Capacity, data.CreatedAt, data.CreatedName, data.UpdatedAt, data.UpdatedName)
	}, tpmtAssetIdKey)
	return ret, err
}

func (m *defaultTpmtAssetModel) Update(ctx context.Context, data *TpmtAsset) error {
	tpmtAssetIdKey := fmt.Sprintf("%s%v", cacheTpmtAssetIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, tpmtAssetRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.AssetType, data.AssetCode, data.AssetName, data.AssetModel, data.ManuFacturer, data.Voltage, data.Capacity, data.CreatedAt, data.CreatedName, data.UpdatedAt, data.UpdatedName, data.Id)
	}, tpmtAssetIdKey)
	return err
}

func (m *defaultTpmtAssetModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(tpmtAssetRows).From(m.table)
}

func (m *defaultTpmtAssetModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(m.table)
}

func (m *defaultTpmtAssetModel) SumBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("IFNULL(SUM(" + field + "),0)").From(m.table)
}

func (m *defaultTpmtAssetModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {

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

func (m *defaultTpmtAssetModel) FindList(ctx context.Context, rowBuilder squirrel.SelectBuilder, current, pageSize int64) ([]*TpmtAsset, error) {

	if current < 1 {
		current = 1
	}
	offset := (current - 1) * pageSize

	query, values, err := rowBuilder.Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*TpmtAsset
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

func (m *defaultTpmtAssetModel) TransCtx(ctx context.Context, fn func(ctx context.Context, sqlx sqlx.Session) error) error {
	return m.Transact(func(s sqlx.Session) error {
		return fn(ctx, s)
	})
}

func (m *defaultTpmtAssetModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheTpmtAssetIdPrefix, primary)
}

func (m *defaultTpmtAssetModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", tpmtAssetRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultTpmtAssetModel) tableName() string {
	return m.table
}
