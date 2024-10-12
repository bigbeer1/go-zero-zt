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
	tpmtGatewayFieldNames          = builder.RawFieldNames(&TpmtGateway{})
	tpmtGatewayRows                = strings.Join(tpmtGatewayFieldNames, ",")
	tpmtGatewayRowsExpectAutoSet   = strings.Join(stringx.Remove(tpmtGatewayFieldNames, "`create_at`", "`create_time`", "`update_at`", "`update_time`"), ",")
	tpmtGatewayRowsWithPlaceHolder = strings.Join(stringx.Remove(tpmtGatewayFieldNames, "`id`", "`create_at`", "`create_time`", "`update_at`", "`update_time`"), "=?,") + "=?"

	cacheTpmtGatewayIdPrefix = "cache:tpmtGateway:id:"
)

type (
	tpmtGatewayModel interface {
		Insert(ctx context.Context, data *TpmtGateway) (sql.Result, error)
		FindOne(ctx context.Context, id string) (*TpmtGateway, error)
		Update(ctx context.Context, data *TpmtGateway) error
		Delete(ctx context.Context, id string) error
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
		FindList(ctx context.Context, rowBuilder squirrel.SelectBuilder, current, pageSize int64) ([]*TpmtGateway, error)
		RowBuilder() squirrel.SelectBuilder
		CountBuilder(field string) squirrel.SelectBuilder
		SumBuilder(field string) squirrel.SelectBuilder
		TransCtx(ctx context.Context, fn func(ctx context.Context, sqlx sqlx.Session) error) error
	}

	defaultTpmtGatewayModel struct {
		sqlc.CachedConn
		table string
	}

	TpmtGateway struct {
		Id           string         `db:"id"`            // 采集器ID/网关
		CreatedAt    time.Time      `db:"created_at"`    // 创建时间
		UpdatedAt    sql.NullTime   `db:"updated_at"`    // 更新时间
		CreatedName  string         `db:"created_name"`  // 创建人
		UpdatedName  sql.NullString `db:"updated_name"`  // 更新人
		GatewayName  string         `db:"gateway_name"`  // 网关名称
		GatewayModel string         `db:"gateway_model"` // 网关型号
		ManuFacturer string         `db:"manu_facturer"` // 生产厂家
		Agreement    int64          `db:"agreement"`     // 协议 默认1:modbus
		BaudRate     int64          `db:"baud_rate"`     // 波特率
		Parity       string         `db:"parity"`        // 校验
		DataBits     int64          `db:"data_bits"`     // 数据位
		StopBits     int64          `db:"stop_bits"`     // 停止位
		ComPort      string         `db:"com_port"`      // com端口
		AddressCode  int64          `db:"address_code"`  // 地址码
	}
)

func newTpmtGatewayModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultTpmtGatewayModel {
	return &defaultTpmtGatewayModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`tpmt_gateway`",
	}
}

func (m *defaultTpmtGatewayModel) Delete(ctx context.Context, id string) error {
	tpmtGatewayIdKey := fmt.Sprintf("%s%v", cacheTpmtGatewayIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, tpmtGatewayIdKey)
	return err
}

func (m *defaultTpmtGatewayModel) FindOne(ctx context.Context, id string) (*TpmtGateway, error) {
	tpmtGatewayIdKey := fmt.Sprintf("%s%v", cacheTpmtGatewayIdPrefix, id)
	var resp TpmtGateway
	err := m.QueryRowCtx(ctx, &resp, tpmtGatewayIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", tpmtGatewayRows, m.table)
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

func (m *defaultTpmtGatewayModel) Insert(ctx context.Context, data *TpmtGateway) (sql.Result, error) {
	tpmtGatewayIdKey := fmt.Sprintf("%s%v", cacheTpmtGatewayIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, tpmtGatewayRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Id, data.CreatedAt, data.UpdatedAt, data.CreatedName, data.UpdatedName, data.GatewayName, data.GatewayModel, data.ManuFacturer, data.Agreement, data.BaudRate, data.Parity, data.DataBits, data.StopBits, data.ComPort, data.AddressCode)
	}, tpmtGatewayIdKey)
	return ret, err
}

func (m *defaultTpmtGatewayModel) Update(ctx context.Context, data *TpmtGateway) error {
	tpmtGatewayIdKey := fmt.Sprintf("%s%v", cacheTpmtGatewayIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, tpmtGatewayRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.CreatedAt, data.UpdatedAt, data.CreatedName, data.UpdatedName, data.GatewayName, data.GatewayModel, data.ManuFacturer, data.Agreement, data.BaudRate, data.Parity, data.DataBits, data.StopBits, data.ComPort, data.AddressCode, data.Id)
	}, tpmtGatewayIdKey)
	return err
}

func (m *defaultTpmtGatewayModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(tpmtGatewayRows).From(m.table)
}

func (m *defaultTpmtGatewayModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(m.table)
}

func (m *defaultTpmtGatewayModel) SumBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("IFNULL(SUM(" + field + "),0)").From(m.table)
}

func (m *defaultTpmtGatewayModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {

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

func (m *defaultTpmtGatewayModel) FindList(ctx context.Context, rowBuilder squirrel.SelectBuilder, current, pageSize int64) ([]*TpmtGateway, error) {

	if current < 1 {
		current = 1
	}
	offset := (current - 1) * pageSize

	query, values, err := rowBuilder.Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*TpmtGateway
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

func (m *defaultTpmtGatewayModel) TransCtx(ctx context.Context, fn func(ctx context.Context, sqlx sqlx.Session) error) error {
	return m.Transact(func(s sqlx.Session) error {
		return fn(ctx, s)
	})
}

func (m *defaultTpmtGatewayModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheTpmtGatewayIdPrefix, primary)
}

func (m *defaultTpmtGatewayModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", tpmtGatewayRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultTpmtGatewayModel) tableName() string {
	return m.table
}
