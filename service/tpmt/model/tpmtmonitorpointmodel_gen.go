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
	tpmtMonitorPointFieldNames          = builder.RawFieldNames(&TpmtMonitorPoint{})
	tpmtMonitorPointRows                = strings.Join(tpmtMonitorPointFieldNames, ",")
	tpmtMonitorPointRowsExpectAutoSet   = strings.Join(stringx.Remove(tpmtMonitorPointFieldNames, "`id`", "`create_at`", "`create_time`", "`update_at`", "`update_time`"), ",")
	tpmtMonitorPointRowsWithPlaceHolder = strings.Join(stringx.Remove(tpmtMonitorPointFieldNames, "`id`", "`create_at`", "`create_time`", "`update_at`", "`update_time`"), "=?,") + "=?"

	cacheTpmtMonitorPointIdPrefix = "cache:tpmtMonitorPoint:id:"
)

type (
	tpmtMonitorPointModel interface {
		Insert(ctx context.Context, data *TpmtMonitorPoint) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*TpmtMonitorPoint, error)
		Update(ctx context.Context, data *TpmtMonitorPoint) error
		Delete(ctx context.Context, id int64) error
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
		FindList(ctx context.Context, rowBuilder squirrel.SelectBuilder, current, pageSize int64) ([]*TpmtMonitorPoint, error)
		RowBuilder() squirrel.SelectBuilder
		CountBuilder(field string) squirrel.SelectBuilder
		SumBuilder(field string) squirrel.SelectBuilder
		TransCtx(ctx context.Context, fn func(ctx context.Context, sqlx sqlx.Session) error) error
	}

	defaultTpmtMonitorPointModel struct {
		sqlc.CachedConn
		table string
	}

	TpmtMonitorPoint struct {
		Id                        int64          `db:"id"`                          // 监测点ID
		CreatedAt                 time.Time      `db:"created_at"`                  // 创建时间
		UpdatedAt                 sql.NullTime   `db:"updated_at"`                  // 更新时间
		CreatedName               string         `db:"created_name"`                // 创建人
		UpdatedName               sql.NullString `db:"updated_name"`                // 更新人
		SerialNumber              string         `db:"serial_number"`               // 编号
		Name                      string         `db:"name"`                        // 监测点名称
		RegisterAddress           string         `db:"register_address"`            // 寄存器地址
		PointCollectorInstruction int64          `db:"point_collector_instruction"` // 采集器指令  1: 01  2: 02  3:03  4:04
		PointAnalysisRule         int64          `db:"point_analysis_rule"`         // 采集器解析规则 1: 16位无符号/2:单精度浮点数
		PointType                 int64          `db:"point_type"`                  // 类型：1:综保/2:局放/3:测温/4:微水/5:油色谱/6:机器人/7:其他
		PointCategory             int64          `db:"point_category"`              // 类别：1:遥信/2:遥测/3:遥脉
		PointGroup                int64          `db:"point_group"`                 // 分组
		CircuitType               int64          `db:"circuit_type"`                // 回路类型
		YxDecode                  sql.NullString `db:"yx_decode"`                   // 遥信解译
		DataBits                  int64          `db:"data_bits"`                   // 数据位
		Coefficient               float64        `db:"coefficient"`                 // 系数
		RetainDecimals            int64          `db:"retain_decimals"`             // 保留小数位
		Unit                      sql.NullString `db:"unit"`                        // 单位
		AlarmDuration             int64          `db:"alarm_duration"`              // 持续时间
		AlarmUpValue              float64        `db:"alarm_up_value"`              // 告警上限
		AlarmDownValue            float64        `db:"alarm_down_value"`            // 告警下限
		WarningUpValue            float64        `db:"warning_up_value"`            // 预警上限
		WarningDownValue          float64        `db:"warning_down_value"`          // 预警下限
		IsDisplacementWarning     int64          `db:"is_displacement_warning"`     // 变位预警 0 不启用 1:启用
		TpmtGatewayId             string         `db:"tpmt_gateway_id"`             // 网关ID
		AssetId                   string         `db:"asset_id"`                    // 资产ID
		Sort                      int64          `db:"sort"`                        // 排序
	}
)

func newTpmtMonitorPointModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultTpmtMonitorPointModel {
	return &defaultTpmtMonitorPointModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`tpmt_monitor_point`",
	}
}

func (m *defaultTpmtMonitorPointModel) Delete(ctx context.Context, id int64) error {
	tpmtMonitorPointIdKey := fmt.Sprintf("%s%v", cacheTpmtMonitorPointIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, tpmtMonitorPointIdKey)
	return err
}

func (m *defaultTpmtMonitorPointModel) FindOne(ctx context.Context, id int64) (*TpmtMonitorPoint, error) {
	tpmtMonitorPointIdKey := fmt.Sprintf("%s%v", cacheTpmtMonitorPointIdPrefix, id)
	var resp TpmtMonitorPoint
	err := m.QueryRowCtx(ctx, &resp, tpmtMonitorPointIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", tpmtMonitorPointRows, m.table)
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

func (m *defaultTpmtMonitorPointModel) Insert(ctx context.Context, data *TpmtMonitorPoint) (sql.Result, error) {
	tpmtMonitorPointIdKey := fmt.Sprintf("%s%v", cacheTpmtMonitorPointIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, tpmtMonitorPointRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.CreatedAt, data.UpdatedAt, data.CreatedName, data.UpdatedName, data.SerialNumber, data.Name, data.RegisterAddress, data.PointCollectorInstruction, data.PointAnalysisRule, data.PointType, data.PointCategory, data.PointGroup, data.CircuitType, data.YxDecode, data.DataBits, data.Coefficient, data.RetainDecimals, data.Unit, data.AlarmDuration, data.AlarmUpValue, data.AlarmDownValue, data.WarningUpValue, data.WarningDownValue, data.IsDisplacementWarning, data.TpmtGatewayId, data.AssetId, data.Sort)
	}, tpmtMonitorPointIdKey)
	return ret, err
}

func (m *defaultTpmtMonitorPointModel) Update(ctx context.Context, data *TpmtMonitorPoint) error {
	tpmtMonitorPointIdKey := fmt.Sprintf("%s%v", cacheTpmtMonitorPointIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, tpmtMonitorPointRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.CreatedAt, data.UpdatedAt, data.CreatedName, data.UpdatedName, data.SerialNumber, data.Name, data.RegisterAddress, data.PointCollectorInstruction, data.PointAnalysisRule, data.PointType, data.PointCategory, data.PointGroup, data.CircuitType, data.YxDecode, data.DataBits, data.Coefficient, data.RetainDecimals, data.Unit, data.AlarmDuration, data.AlarmUpValue, data.AlarmDownValue, data.WarningUpValue, data.WarningDownValue, data.IsDisplacementWarning, data.TpmtGatewayId, data.AssetId, data.Sort, data.Id)
	}, tpmtMonitorPointIdKey)
	return err
}

func (m *defaultTpmtMonitorPointModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(tpmtMonitorPointRows).From(m.table)
}

func (m *defaultTpmtMonitorPointModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(m.table)
}

func (m *defaultTpmtMonitorPointModel) SumBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("IFNULL(SUM(" + field + "),0)").From(m.table)
}

func (m *defaultTpmtMonitorPointModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {

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

func (m *defaultTpmtMonitorPointModel) FindList(ctx context.Context, rowBuilder squirrel.SelectBuilder, current, pageSize int64) ([]*TpmtMonitorPoint, error) {

	if current < 1 {
		current = 1
	}
	offset := (current - 1) * pageSize

	query, values, err := rowBuilder.Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*TpmtMonitorPoint
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

func (m *defaultTpmtMonitorPointModel) TransCtx(ctx context.Context, fn func(ctx context.Context, sqlx sqlx.Session) error) error {
	return m.Transact(func(s sqlx.Session) error {
		return fn(ctx, s)
	})
}

func (m *defaultTpmtMonitorPointModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheTpmtMonitorPointIdPrefix, primary)
}

func (m *defaultTpmtMonitorPointModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", tpmtMonitorPointRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultTpmtMonitorPointModel) tableName() string {
	return m.table
}
