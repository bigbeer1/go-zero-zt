package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/jsonx"
	"time"
	"tpmt-zt/common/tdenginex"
)

type AlarmLog struct {
	Ts           time.Time `json:"ts"`            // 创建时间
	Id           string    `json:"id"`            // 告警ID
	Mid          string    `json:"mid"`           // mid 预警设备ID
	Name         string    `json:"name"`          // 设备名称
	AlarmType    int64     `json:"alarm_type"`    // 类型：1 越上限/2 越下限/ 3 变位 / 4 网关下线
	AlarmGrade   int64     `json:"alarm_grade"`   // 等级：1 预警/2 告警 /3 提醒
	AlarmContent string    `json:"alarm_content"` // 数据  返回内容
	AssetId      string    `json:"asset_id"`      // 资产ID
	AlarmState   int64     `json:"alarm_state"`   // 状态 0 未读  1已读  2已确认
}

func (t *AlarmLog) Insert(ctx context.Context, taos *sql.DB, tddb *TdDb) error {

	// 拼接请求数据库和表
	dbName := fmt.Sprintf("INSERT INTO %s USING %s ", tddb.DbName, tddb.TableName)

	// 拼接参数
	tableData := fmt.Sprintf("Tags('%v') (`ts`,`id`,`mid`,`name`,`alarm_type`,`alarm_grade`,`alarm_content`,`asset_id`,`alarm_state`)", "1")

	value := fmt.Sprintf(" values ('%v','%v','%v','%v','%v','%v','%v','%v','%v');", t.Ts.Format(time.RFC3339Nano), t.Id, t.Mid, t.Name, t.AlarmType, t.AlarmGrade,
		t.AlarmContent, t.AssetId, t.AlarmState)

	sqlx := dbName + tableData + value

	_, err := taos.ExecContext(ctx, sqlx)
	if err != nil {
		return err
	}

	return nil
}

func (t *AlarmLog) FindList(ctx context.Context, taos *sql.DB, tddb *TdDb, current, pageSize, startTime, endTime int64) (resp []*AlarmLog, err error) {

	// 拼接请求数据库和表
	sqlData := squirrel.Select("*").From(tddb.DbName).OrderBy("`ts` desc")
	// 拼接参数

	if current < 1 {
		current = 1
	}
	offset := (current - 1) * pageSize
	sqlData = t.BetweenTime(sqlData, startTime, endTime)

	sqlx, values, _ := t.fillFilter(sqlData).Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()

	rows, err := taos.QueryContext(ctx, sqlx, values...)
	if err != nil {
		return nil, err
	}

	var datas []map[string]any
	err = tdenginex.Scan(rows, &datas)
	if err != nil {
		return nil, err
	}

	databytes, _ := jsonx.Marshal(datas)
	err = jsonx.Unmarshal(databytes, &resp)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (t *AlarmLog) FindOne(ctx context.Context, taos *sql.DB, tddb *TdDb) (resp []*AlarmLog, err error) {

	// 拼接请求数据库和表
	sqlData := squirrel.Select("*").From(tddb.DbName).OrderBy("`ts` desc")
	// 拼接参数

	current := 1
	offset := (current - 1) * 1

	sqlx, values, _ := t.fillFilter(sqlData).Offset(uint64(offset)).Limit(uint64(1)).ToSql()

	rows, err := taos.QueryContext(ctx, sqlx, values...)
	if err != nil {
		return nil, err
	}

	var datas []map[string]any
	err = tdenginex.Scan(rows, &datas)
	if err != nil {
		return nil, err
	}

	databytes, _ := jsonx.Marshal(datas)
	err = jsonx.Unmarshal(databytes, &resp)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (t *AlarmLog) Count(ctx context.Context, taos *sql.DB, tddb *TdDb, startTime, endTime int64) int64 {

	// 拼接请求数据库和表
	sqlData := squirrel.Select("COUNT('id')").From(tddb.DbName)
	// 拼接参数

	sqlData = t.BetweenTime(sqlData, startTime, endTime)

	sqlx, values, _ := t.fillFilter(sqlData).ToSql()

	rows, err := taos.QueryContext(ctx, sqlx, values...)
	if err != nil {
		return 0
	}

	var count int64
	err = tdenginex.Scan(rows, &count)
	if err != nil {
		return 0
	}

	return count
}

func (t *AlarmLog) fillFilter(sql squirrel.SelectBuilder) squirrel.SelectBuilder {
	if len(t.Id) > 0 {
		sql = sql.Where(fmt.Sprintf(" `id`= '%v' ", t.Id))
	}
	if len(t.Mid) > 0 {
		sql = sql.Where(fmt.Sprintf(" `mid`= '%v' ", t.Mid))
	}
	if t.AlarmType != 99 {
		sql = sql.Where(fmt.Sprintf(" `alarm_type`= '%v' ", t.AlarmType))
	}
	if t.AlarmGrade != 99 {
		sql = sql.Where(fmt.Sprintf(" `alarm_grade`= '%v' ", t.AlarmGrade))
	}
	if len(t.AssetId) > 0 {
		sql = sql.Where(fmt.Sprintf(" `asset_code`= '%v' ", t.AssetId))
	}
	if t.AlarmState != 99 {
		sql = sql.Where(fmt.Sprintf(" `alarm_state`= '%v' ", t.AlarmState))
	}

	return sql

}

func (t *AlarmLog) BetweenTime(sql squirrel.SelectBuilder, startTime, endTime int64) squirrel.SelectBuilder {
	if startTime != 0 {
		sql = sql.Where(fmt.Sprintf(" ts >= '%v'", time.UnixMilli(startTime).Format(time.RFC3339Nano)))
	}

	if endTime != 0 {
		sql = sql.Where(fmt.Sprintf(" ts <= '%v'", time.UnixMilli(endTime).Format(time.RFC3339Nano)))
	}

	if startTime == 0 && endTime == 0 {
		sql = sql.Where(fmt.Sprintf(" ts >= '%v'", time.UnixMilli(time.Now().UnixMilli()-43200000).Format(time.RFC3339Nano)))
	}
	return sql
}
