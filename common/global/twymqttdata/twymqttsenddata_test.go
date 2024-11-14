package twymqttdata

import (
	"fmt"
	"testing"
	"tpmt-zt/common/datax"
	"tpmt-zt/common/jsonx"
)

func TestSchedulerTaskNumber2Data(t *testing.T) {

	var data = &SchedulerTaskNumber2Data{
		GwSn:      "88888",
		MsgInfo:   "11111",
		MqttHost:  "127.0.0.1:1883",
		MqttUser:  "admin",
		MqttPass:  "1qaz2wsx",
		SendTopic: "dev_send/f39a72ec-f45d-45eb-bf12-790aecdfec81",
	}

	ggg := datax.ToString(data)
	fmt.Println(ggg)

}

func TestSchedulerTaskNumber2Data99(t *testing.T) {

	var data TwyData

	dataString := "{\"gw_sn\":\"88888\",\"msg_type\":1001,\"msg_info\":\"11111\",\"time_stamp\":1710832662064,\"data\":[{\"72\":\"0\"},{\"28\":\"0\"},{\"82\":\"0\"},{\"91\":\"0\"},{\"53\":\"0\"},{\"62\":\"0\"},{\"13\":\"0\"},{\"30\":\"0\"},{\"10\":\"0\"},{\"80\":\"0\"},{\"43\":\"0\"},{\"64\":\"0\"},{\"95\":\"0\"},{\"56\":\"0\"},{\"9\":\"0\"},{\"84\":\"0\"},{\"91\":\"0\"},{\"50\":\"0\"},{\"81\":\"0\"},{\"102\":\"0\"},{\"80\":\"0\"},{\"28\":\"0\"},{\"76\":\"0\"},{\"96\":\"0\"},{\"94\":\"0\"},{\"100\":\"0\"},{\"79\":\"0\"},{\"6\":\"0\"},{\"103\":\"0\"},{\"99\":\"0\"},{\"44\":\"0\"},{\"77\":\"0\"},{\"42\":\"0\"},{\"36\":\"0\"},{\"801\":\"0\"},{\"33\":\"0\"},{\"64\":\"0\"},{\"8\":\"0\"},{\"19\":\"0\"},{\"41\":\"0\"}]}"

	jsonx.Str2Struct(dataString, &data)

	fmt.Println(data)

}
