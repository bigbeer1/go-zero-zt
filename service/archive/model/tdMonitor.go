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

type TdMonitor struct {
	Ts   time.Time `json:"ts"`   // 创建时间
	Data float64   `json:"data"` // 数据
}

type TdAVGTimeMonitor struct {
	Wend time.Time `json:"_wend"`     // 创建时间
	Data float64   `json:"avg(data)"` // 数据
}

type TdAVGMonitor struct {
	Data float64 `json:"avg(data)"` // 数据
}

type TdLASTTimeMonitor struct {
	Wstart time.Time `json:"_wstart"`    // 创建时间
	Data   float64   `json:"last(data)"` // 数据
}

type TdRangeMonitor struct {
	Wstart    time.Time `json:"_wstart"`
	StartTs   time.Time `json:"first(ts)"`
	StartData float64   `json:"first(data)"`
	LastTs    time.Time `json:"last(ts)"`
	LastData  float64   `json:"last(data)"`
	RangeData float64   `json:"last(data)-first(data)"`
}

type TdRangeMonitorOnce struct {
	StartTs   time.Time `json:"first(ts)"`
	StartData float64   `json:"first(data)"`
	LastTs    time.Time `json:"last(ts)"`
	LastData  float64   `json:"last(data)"`
	RangeData float64   `json:"last(data)-first(data)"`
}

func (t *TdMonitor) Insert(ctx context.Context, taos *sql.DB, tddb *TdDb) error {

	// 拼接请求数据库和表
	dbName := fmt.Sprintf("INSERT INTO %s USING %s ", tddb.DbName, tddb.TableName)

	// 拼接参数
	tableData := fmt.Sprintf(" Tags('%v') (`ts`,`data`)", 1)

	value := fmt.Sprintf(" values ('%v',%v);", t.Ts.Format(time.RFC3339Nano), t.Data)

	sqlx := dbName + tableData + value

	_, err := taos.ExecContext(ctx, sqlx)
	if err != nil {
		return err
	}
	return nil
}

func (t *TdMonitor) FindAll(ctx context.Context, taos *sql.DB, tddb *TdDb, startTime, endTime int64, every string) (resp []*TdAVGTimeMonitor, err error) {

	// 拼接请求数据库和表
	sqlData := squirrel.Select(" _WEND, AVG(data)").From(tddb.DbName)
	// 拼接参数

	sqlx, values, _ := t.BetweenTime(sqlData, startTime, endTime).ToSql()

	sqlx = sqlx + fmt.Sprintf(" INTERVAL(%s)  FILL(value,0);", every)

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

func (t *TdMonitor) FindAvgByTime(ctx context.Context, taos *sql.DB, tddb *TdDb, startTime, endTime int64) (resp []*TdAVGMonitor, err error) {

	// 拼接请求数据库和表
	sqlData := squirrel.Select("AVG(data)").From(tddb.DbName)
	// 拼接参数

	sqlx, values, _ := t.BetweenTime(sqlData, startTime, endTime).ToSql()

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

func (t *TdMonitor) BetweenTime(sql squirrel.SelectBuilder, startTime, endTime int64) squirrel.SelectBuilder {
	if startTime != 0 {
		sql = sql.Where(fmt.Sprintf(" ts >= '%v'", time.UnixMilli(startTime).Format(time.RFC3339Nano)))
	} else {
		sql = sql.Where(fmt.Sprintf(" ts >= '%v'", time.UnixMilli(time.Now().UnixMilli()-43200000).Format(time.RFC3339Nano)))
	}

	if endTime != 0 {
		sql = sql.Where(fmt.Sprintf(" ts <= '%v'", time.UnixMilli(endTime).Format(time.RFC3339Nano)))
	} else {
		sql = sql.Where(fmt.Sprintf(" ts <= '%v'", time.Now().Format(time.RFC3339Nano)))
	}

	return sql
}

func (t *TdMonitor) EveryDayLastData(ctx context.Context, taos *sql.DB, tddb *TdDb, startTime, endTime int64, every string) (resp []*TdLASTTimeMonitor, err error) {

	// 拼接请求数据库和表
	sqlData := squirrel.Select(" _WSTART, LAST(data)").From(tddb.DbName)
	//减去结束一秒 来确保不查出明天数据
	endTime = endTime - 1000
	// 拼接参数
	sqlx, values, _ := t.BetweenTime(sqlData, startTime, endTime).ToSql()

	sqlx = sqlx + fmt.Sprintf(" INTERVAL(%s)  FILL(value,0);", every)

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

func (t *TdMonitor) HistoricalRangeMonitor(ctx context.Context, taos *sql.DB, tddb *TdDb, startTime, endTime int64, every string) (resp []*TdRangeMonitor, err error) {

	// 拼接请求数据库和表
	sqlData := squirrel.Select(" _WSTART, LAST(ts) , FIRST(ts) , FIRST(data) , LAST(data) ,LAST(data)-FIRST(data)").From(tddb.DbName)
	// 拼接参数

	sqlx, values, _ := t.BetweenTime(sqlData, startTime, endTime).ToSql()

	sqlx = sqlx + fmt.Sprintf(" INTERVAL(%s) FILL(NULL)", every)

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

func (t *TdMonitor) HistoricalRangeMonitorOnce(ctx context.Context, taos *sql.DB, tddb *TdDb, startTime, endTime int64) (resp TdRangeMonitorOnce, err error) {

	// 拼接请求数据库和表
	sqlData := squirrel.Select(" LAST(ts) , FIRST(ts) , FIRST(data) , LAST(data) ,LAST(data)-FIRST(data)").From(tddb.DbName)
	// 拼接参数

	sqlx, values, _ := t.BetweenTime(sqlData, startTime, endTime).ToSql()

	rows, err := taos.QueryContext(ctx, sqlx, values...)
	if err != nil {
		return resp, err
	}

	var datas []map[string]any
	err = tdenginex.Scan(rows, &datas)
	if err != nil {
		return resp, err
	}

	for _, data := range datas {
		databytes, _ := jsonx.Marshal(data)
		err = jsonx.Unmarshal(databytes, &resp)
		if err != nil {
			return resp, err
		}
	}

	return resp, err
}

func (t *TdMonitor) Count(ctx context.Context, taos *sql.DB, tddb *TdDb, startTime, endTime int64) int64 {

	// 拼接请求数据库和表
	sqlData := squirrel.Select("COUNT('ts')").From(tddb.DbName)
	// 拼接参数

	sqlData = t.BetweenTime(sqlData, startTime, endTime)

	sqlx, values, _ := sqlData.ToSql()

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
