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

type ScheduledTasksLog struct {
	Ts               time.Time `json:"ts"`                 // 创建时间
	ScheduledTasksId string    `json:"scheduled_tasks_id"` // 自定义任务ID
	IsRequest        int64     `json:"is_request"`         // 状态  1:成功  其他失败
	RequestData      string    `json:"request_data"`       // 数据  发送内容
	ResponseData     string    `json:"response_data"`      // 数据  返回内容
	TenantId         string    `json:"tenant_id"`          // 租户ID
}

func (t *ScheduledTasksLog) Insert(ctx context.Context, taos *sql.DB, tddb *TdDb) error {

	if t.TenantId == "" {
		t.TenantId = "0"
	}
	// 拼接请求数据库和表
	dbName := fmt.Sprintf("INSERT INTO %s USING %s ", tddb.DbName, tddb.TableName)

	// 拼接参数
	tableData := fmt.Sprintf(" Tags('%v') (`ts`,`scheduled_tasks_id`,`is_request`,`request_data`,`response_data`)", t.TenantId)

	value := fmt.Sprintf(" values ('%v','%v','%v','%v','%v');", t.Ts.Format(time.RFC3339Nano), t.ScheduledTasksId, t.IsRequest, t.RequestData, t.ResponseData)

	sqlx := dbName + tableData + value

	_, err := taos.ExecContext(ctx, sqlx)
	if err != nil {
		return err
	}
	return nil
}

func (t *ScheduledTasksLog) FindList(ctx context.Context, taos *sql.DB, tddb *TdDb, current, pageSize, startTime, endTime int64) (resp []*ScheduledTasksLog, err error) {

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

func (t *ScheduledTasksLog) Count(ctx context.Context, taos *sql.DB, tddb *TdDb, startTime, endTime int64) int64 {

	// 拼接请求数据库和表
	sqlData := squirrel.Select("COUNT('scheduled_tasks_id')").From(tddb.DbName)
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

func (t *ScheduledTasksLog) BetweenTime(sql squirrel.SelectBuilder, startTime, endTime int64) squirrel.SelectBuilder {
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

func (t *ScheduledTasksLog) fillFilter(sql squirrel.SelectBuilder) squirrel.SelectBuilder {
	if len(t.ScheduledTasksId) != 0 {
		sql = sql.Where(fmt.Sprintf(" `scheduled_tasks_id`= '%v' ", t.ScheduledTasksId))
	}

	if t.IsRequest != 99 {
		sql = sql.Where(fmt.Sprintf(" `is_request` = '%v' ", t.IsRequest))
	}
	return sql
}
