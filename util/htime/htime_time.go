package htime

import (
	"time"

	"github.com/gogf/gf/v2/os/gtime"
	carbon "github.com/golang-module/carbon/v2"
)

// 获取当前时间戳/秒
func GetUnix() int64 {
	return time.Now().Unix()
}

// 获取当前毫秒
func GetUnixMilli() int64 {
	return time.Now().UnixMilli()
}

// 获取当前纳秒
func GetUnixNano() int64 {
	return time.Now().UnixNano()
}

// 获取当前日期
func GetDate() string {
	template := "2006-01-02 15:04:05"
	return time.Now().Format(template)
}

// 获取年月日
func GetDay() string {
	template := "20060102"
	return time.Now().Format(template)
}

// 时间戳转日期
func UnixToTime(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	return t.Format("2006-01-02 15:04:05")
}

// 日期转时间戳
func DateToUnix(date string) int64 {
	template := "2006-01-02 15:04:05"
	t, err := time.ParseInLocation(template, date, time.Local)
	if err != nil {
		return 0
	}
	return t.Unix()
}

// 今日结束的时间
func EndOfDay() string {
	return gtime.Now().EndOfDay().String()
}

// 当前时间戳
func NowTimestamp() int64 {
	return gtime.Now().Timestamp()
}

// 获取2个时间间隔的 秒/分/时/天/周/月/年
// start_time:起始日期 (2023-12-12 16:11:11)
// end_time: 结束日期 (...)
// step 返回说明(1:秒 2:分 3:小时 4:天 5:周 6:月 7:年)
func DateBetweenDiffStep(start_time string, end_time string, step int64) int64 {

	switch step {
	case 1:
		// 秒

		return carbon.Parse(end_time).DiffAbsInSeconds(carbon.Parse(start_time))

	case 2:
		// 分

		return carbon.Parse(end_time).DiffAbsInMinutes(carbon.Parse(start_time))

	case 3:
		// 时

		return carbon.Parse(end_time).DiffAbsInHours(carbon.Parse(start_time))

	case 4:
		// 天

		return carbon.Parse(end_time).DiffAbsInDays(carbon.Parse(start_time))

	case 5:
		// 周

		return carbon.Parse(end_time).DiffAbsInWeeks(carbon.Parse(start_time))

	case 6:
		// 月

		return carbon.Parse(end_time).DiffAbsInMonths(carbon.Parse(start_time))

	case 7:
		// 年

		return carbon.Parse(end_time).DiffAbsInYears(carbon.Parse(start_time))

	}

	return 0
}
