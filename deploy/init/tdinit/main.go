package main

import (
	"database/sql"
	"fmt"
	_ "github.com/taosdata/driver-go/v3/taosRestful"
	"log"
	"time"
)

func main() {

	var taosUri = "root:POOTACA93V@http(127.0.0.1:6041)/"
	taos, err := sql.Open("taosRestful", taosUri)
	if err != nil {
		fmt.Println("failed to connect TDengine, err:", err)
		return
	}

	defer taos.Close()
	createMonitorPoint(taos, 180)
	createLog(taos, "app_log", 90)
	createLog(taos, "alarm_log", 90)
	createLog(taos, "scheduled_tasks_log", 30)

	time.Sleep(time.Second * 1)
	a := fmt.Sprintf("CREATE STABLE monitor_point.tpmt_monitor_point (ts TIMESTAMP, data FLOAT) TAGS (tenant_id BINARY(36))")

	CreateStable(taos, a)
	time.Sleep(time.Second * 1)

	a = fmt.Sprintf("CREATE STABLE app_log.tpmt_app_log (created_time TIMESTAMP, uid BINARY(36), created_name BINARY(36)" +
		",ip BINARY(36), interface_type BINARY(36), interface_address BINARY(36), request_data VARCHAR(30000), " +
		"is_request INT, response_data VARCHAR(30000) ,timed BIGINT) TAGS (tenant_id BINARY(36))")

	CreateStable(taos, a)
	time.Sleep(time.Second * 1)

	a = fmt.Sprintf("CREATE STABLE scheduled_tasks_log.ptm_scheduled_tasks (ts TIMESTAMP, scheduled_tasks_id BINARY(100)" +
		", is_request INT" +
		", request_data VARCHAR(30000)" +
		", response_data VARCHAR(30000))" +
		" TAGS (tenant_id BINARY(36))")
	CreateStable(taos, a)

	time.Sleep(time.Second * 1)

	a = fmt.Sprintf("CREATE STABLE scheduled_tasks_log.ptm_scheduled_tasks_failure_record (ts TIMESTAMP, scheduled_tasks_id BINARY(100)" +
		", is_request INT" +
		", request_data VARCHAR(30000)" +
		", response_data VARCHAR(30000))" +
		" TAGS (tenant_id BINARY(36))")
	CreateStable(taos, a)

	time.Sleep(time.Second * 1)

	a = fmt.Sprintf("CREATE STABLE alarm_log.ptm_alarm_log (ts TIMESTAMP, id VARCHAR(40)" +
		", mid VARCHAR(40)" +
		", name VARCHAR(100)" +
		", alarm_type INT" +
		", alarm_grade INT" +
		", alarm_content VARCHAR(30000)" +
		", asset_code VARCHAR(30000)" +
		", alarm_state INT)" +
		" TAGS (tenant_id BINARY(36))")
	CreateStable(taos, a)

	time.Sleep(time.Second * 1)

	fmt.Println("初始化成功")

}

func createMonitorPoint(taos *sql.DB, number int) {
	dataString := fmt.Sprintf("CREATE DATABASE monitor_point BUFFER 50 KEEP %vd  VGROUPS 5  ", number)
	_, err := taos.Exec(dataString)
	if err != nil {
		log.Fatalln("failed to create database, err:", err)
	}
}

func createLog(taos *sql.DB, name string, number int) {
	dataString := fmt.Sprintf("CREATE DATABASE %s BUFFER 16 KEEP %vd  VGROUPS 2   ", name, number)
	_, err := taos.Exec(dataString)
	if err != nil {
		log.Fatalln("failed to create database, err:", err)
	}
}

func CreateStable(taos *sql.DB, dataString string) {
	_, err := taos.Exec(dataString)
	if err != nil {
		log.Fatalln("failed to create stable, err:", err)
	}
}
