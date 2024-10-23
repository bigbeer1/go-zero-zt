package eval

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"tpmt-zt/common/datax"
)

func MonitorValueEval(temp string, coefficient float64) (resultValue string, updateTime int64) {

	//获取到的数据分割处理 11|2022-16-16 14:16:40
	arr := strings.Split(temp, "|")

	if len(arr) == 2 {
		monitorValueFloat, err := datax.ToFloat64(arr[0])
		if err != nil {
			resultValue = "-"
			updateTime, _ = datax.ToInt64(arr[1])
			return
		} else {
			// 乘以系数
			monitorValueFloatBig := big.NewFloat(monitorValueFloat)
			monitorValueFloatBig.Mul(monitorValueFloatBig, big.NewFloat(coefficient))
			bit := "%." + "2" + "f"
			monitorValue, _ := strconv.ParseFloat(fmt.Sprintf(bit, monitorValueFloatBig), 64)

			resultValue = datax.ToString(monitorValue)
			updateTime, _ = datax.ToInt64(arr[1])
			return
		}
	} else {
		resultValue = "-"
		updateTime = 0
		return
	}

}
