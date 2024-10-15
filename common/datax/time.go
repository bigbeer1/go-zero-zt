package datax

import (
	"time"
)

// getTimeLoc
/**
 * @Description: 获取中国时区
 * @return l
 */
func getTimeLoc() (l *time.Location) {
	l, _ = time.LoadLocation("Asia/Shanghai")
	return
}

// NowStr
/**
 * @Description: 获取当时时间字符串（2006-04-02 15:04:05）
 * @return timeStr
 */
func NowStr() (timeStr string) {
	l := getTimeLoc()
	t := time.Now().In(l)
	timeStr = t.Format("2006-04-02 15:04:05")
	return
}
