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

type (
	TdAppLogModel interface {
		Insert(ctx context.Context, taos *sql.DB, tddb *TdDb) error
		FindList(ctx context.Context, taos *sql.DB, tddb *TdDb, current, pageSize, startTime, endTime int64) (resp []*TdAppLog, err error)
	}

	TdAppLog struct {
		CreatedTime      time.Time `json:"created_time"`                // 创建人名称
		Uid              string    `json:"uid"`                         // 操作人ID
		CreatedName      string    `json:"created_name"`                // 创建人名称
		Ip               string    `json:"ip"`                          // 请求Ip
		InterfaceType    string    `json:"interface_type"`              // 请求方法
		InterfaceAddress string    `json:"interface_address,omitempty"` // 请求地址
		IsRequest        int64     `json:"is_request,omitempty"`        // 请求结果
		RequestData      string    `json:"request_data,omitempty"`      // 返回内容
		ResponseData     string    `json:"response_data,omitempty"`     // 返回参数
		Timed            int64     `json:"timed,omitempty"`             // 运算时间
	}
)

func (t *TdAppLog) Insert(ctx context.Context, taos *sql.DB, tddb *TdDb) error {

	// 拼接请求数据库和表
	dbName := fmt.Sprintf("INSERT INTO %s USING %s ", tddb.DbName, tddb.TableName)

	// 拼接参数
	tableData := fmt.Sprintf("Tags('%v') (`created_time`,`uid`,`created_name`,`ip`,`interface_type`,`interface_address` "+
		",`is_request`,`request_data`,`response_data`,`timed`)", "1")

	value := fmt.Sprintf(" values ('%v','%v','%v','%v','%v','%v',%v,'%v','%v',%v);", t.CreatedTime.Format(time.RFC3339Nano), t.Uid, t.CreatedName, t.Ip, t.InterfaceType, t.InterfaceAddress, t.IsRequest, t.RequestData, t.ResponseData, t.Timed)

	sqlx := dbName + tableData + value

	_, err := taos.ExecContext(ctx, sqlx)
	if err != nil {
		return err
	}
	return nil
}

func (t *TdAppLog) FindList(ctx context.Context, taos *sql.DB, tddb *TdDb, current, pageSize, startTime, endTime int64) (resp []*TdAppLog, err error) {

	// 拼接请求数据库和表
	sqlData := squirrel.Select("*").From(tddb.DbName).OrderBy("`created_time` desc")
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

func (t *TdAppLog) Count(ctx context.Context, taos *sql.DB, tddb *TdDb, startTime, endTime int64) int64 {

	// 拼接请求数据库和表
	sqlData := squirrel.Select("COUNT('uid')").From(tddb.DbName)
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

func (t *TdAppLog) fillFilter(sql squirrel.SelectBuilder) squirrel.SelectBuilder {
	if len(t.Uid) != 0 {
		sql = sql.Where(fmt.Sprintf(" `uid`= '%v' ", t.Uid))
	}
	if len(t.CreatedName) != 0 {
		sql = sql.Where(fmt.Sprintf(" `created_name` like '%v' ", "%"+t.CreatedName+"%"))
	}
	if len(t.Ip) != 0 {
		sql = sql.Where(fmt.Sprintf(" `ip` like '%v' ", "%"+t.Ip+"%"))
	}
	if len(t.InterfaceType) != 0 {
		sql = sql.Where(fmt.Sprintf(" `interface_type` = '%v' ", t.InterfaceType))
	}
	if len(t.InterfaceAddress) != 0 {
		sql = sql.Where(fmt.Sprintf(" `interface_address` like '%v' ", "%"+t.InterfaceAddress+"%"))
	}
	if t.IsRequest != 99 {
		sql = sql.Where(fmt.Sprintf(" `is_request` = '%v' ", t.IsRequest))
	}
	return sql
}

func (t *TdAppLog) BetweenTime(sql squirrel.SelectBuilder, startTime, endTime int64) squirrel.SelectBuilder {
	if startTime != 0 {
		sql = sql.Where(fmt.Sprintf(" created_time >= '%v'", time.UnixMilli(startTime).Format(time.RFC3339Nano)))
	}

	if endTime != 0 {
		sql = sql.Where(fmt.Sprintf(" created_time <= '%v'", time.UnixMilli(endTime).Format(time.RFC3339Nano)))
	}

	if startTime == 0 && endTime == 0 {
		sql = sql.Where(fmt.Sprintf(" created_time >= '%v'", time.UnixMilli(time.Now().UnixMilli()-43200000).Format(time.RFC3339Nano)))
	}
	return sql
}
