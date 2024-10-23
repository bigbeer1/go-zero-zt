package alarm

import (
	"fmt"
	"tpmt-zt/service/tpmt/model"
)

// 告警级别和规则
type AlarmRuleYc struct {
	AlarmUpValue     float64 // 告警上限
	AlarmDownValue   float64 // 告警下限
	WarningUpValue   float64 // 预警上限
	WarningDownValue float64 // 预警下限
}

// 告警内容
type AlarmRuleInfo struct {
	Level    int64  `json:"level"`     // 告警等级 0 正常  1预警 2告警
	RuleType int64  `json:"rule_type"` // 0正常,1越上限,2越下限 3正常
	RuleData string `json:"rule_data"` // 数值带单位
}

// 是否触发规则
func (a AlarmRuleYc) CheckAlarmRule(value float64) (res *AlarmRuleInfo) {

	// 首先判断告警
	if value > a.AlarmUpValue {
		ruleData := fmt.Sprintf("越上线%v,当前值:%v", a.AlarmUpValue, value)
		res = &AlarmRuleInfo{
			Level:    2,
			RuleType: 1,
			RuleData: ruleData,
		}

		return res
	}

	if value < a.AlarmDownValue {
		ruleData := fmt.Sprintf("越下线%v,当前值:%v", a.AlarmDownValue, value)
		res = &AlarmRuleInfo{
			Level:    2,
			RuleType: 2,
			RuleData: ruleData,
		}
		return res

	}

	if value > a.WarningUpValue {
		ruleData := fmt.Sprintf("越上线%v,当前值:%v", a.WarningUpValue, value)
		res = &AlarmRuleInfo{
			Level:    1,
			RuleType: 1,
			RuleData: ruleData,
		}
		return res

	}

	if value < a.WarningDownValue {
		ruleData := fmt.Sprintf("越下线%v,当前值:%v", a.WarningDownValue, value)
		res = &AlarmRuleInfo{
			Level:    1,
			RuleType: 2,
			RuleData: ruleData,
		}
		return res

	}

	return res

}

// 判断 数值是否告警
func CheckAlarmRuleYc(item *model.TpmtMonitorPoint, resultValueFloat64 float64) (data *AlarmRuleInfo) {

	alarmRuleYc := &AlarmRuleYc{
		AlarmUpValue:     item.AlarmUpValue,
		AlarmDownValue:   item.AlarmDownValue,
		WarningUpValue:   item.WarningUpValue,
		WarningDownValue: item.WarningDownValue,
	}

	// 判断内容数值内容是否告警
	switch item.PointCategory {
	case 2:
		// 判断是否触发告警
		alarmRuleRes := alarmRuleYc.CheckAlarmRule(resultValueFloat64)
		return alarmRuleRes
	case 3:
		// 判断是否触发告警
		alarmRuleRes := alarmRuleYc.CheckAlarmRule(resultValueFloat64)
		return alarmRuleRes
	}

	return data

}
