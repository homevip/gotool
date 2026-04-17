package htime

import (
	carbon "github.com/golang-module/carbon/v2"
)

// 获取北京时间当前时间
func Now() *carbon.Carbon {
	// return carbon.Now(carbon.Shanghai) // 上海时间
	return carbon.Now(carbon.PRC) // 中国时间
}

// 获取当前时间戳/秒
func Unix() int64 {
	return Now().Timestamp()
}

// 获取当前毫秒
func UnixMilli() int64 {
	return Now().TimestampMilli()
}

// 获取当前纳秒
func UnixNano() int64 {
	return Now().TimestampNano()
}

// 获取当前日期
// Y-m-d H:i:s,Y年m月d日 H时i分s秒...
func Date(format string) string {
	return Now().Format(format)
}

// 时间戳转日期
func UnixToTime(timestamp int64) string {
	return carbon.CreateFromTimestamp(timestamp, carbon.PRC).Format("Y-m-d H:i:s")
}

// 日期转时间戳
func DateToUnix(date string) int64 {
	return carbon.Parse(date, carbon.PRC).Timestamp()
}

// 今日结束的时间
func EndOfDay() string {
	return Now().EndOfDay().ToDateTimeString()
}

// 今日结束的时间戳
func EndOfDayTimestamp() int64 {
	return Now().EndOfDay().Timestamp()
}

// 获取2个时间间隔的 秒/分/时/天/周/月/年
// start_time:起始日期 (2023-12-12 16:11:11)
// end_time: 结束日期 (...)
// step 返回说明(1:秒 2:分 3:小时 4:天 5:周 6:月 7:年)
func DateBetweenDiffStep(start_time string, end_time string, step int64) int64 {

	var (
		start = carbon.Parse(start_time, carbon.PRC)
		end   = carbon.Parse(end_time, carbon.PRC)
	)

	switch step {
	case 1:
		// 秒

		return end.DiffAbsInSeconds(start)

	case 2:
		// 分

		return end.DiffAbsInMinutes(start)

	case 3:
		// 时

		return end.DiffAbsInHours(start)

	case 4:
		// 天

		return end.DiffAbsInDays(start)

	case 5:
		// 周

		return end.DiffAbsInWeeks(start)

	case 6:
		// 月

		return end.DiffAbsInMonths(start)

	case 7:
		// 年

		return end.DiffAbsInYears(start)

	}

	return 0
}
