package global

import "time"

type MonitorCacheUpData struct {
	Id                    string    `json:"id"`
	Name                  string    `json:"name"`                    // 监测点名称
	PointCategory         int64     `json:"point_category"`          // 类别：1:遥信/2:遥测/3:遥脉
	IsDisplacementWarning int64     `json:"is_displacement_warning"` // 变位预警 0 不启用 1:启用
	Ts                    time.Time `json:"ts"`
	AssetId               string    `json:"asset_id"` //资产ID
	Data                  string    `json:"data"`
}

type GateWayOnlineUpData struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	OnlineStatus string `json:"online_status"`
}
