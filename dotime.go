package publice

import (
	"github.com/astaxie/beego"
	"strconv"
	"time"
)

func init() {
	//GetBeforeTime()
}

func GetDates(now string, num int, before int, format string) string {
	var strs = "'" + Gettimes("2006-01-02") + "','" + now + "'"
	for i := 0; i < num; i++ {
		now = TimeBeforeTime(now, format, before, false)
		strs += ",'" + now + "'"
	}
	return strs
}

func GetBeforeTime(before int, format string, lessHour int, toTimestemp bool) string {
	times, _ := time.Parse(format, time.Now().AddDate(0, 0, before).Format(format))
	less, _ := strconv.ParseInt(strconv.Itoa(60*60*lessHour), 10, 64)
	if toTimestemp {
		result := strconv.FormatInt(times.Unix()-less, 10)
		return result
	} else {
		result := time.Unix(times.Unix()-less, 10).Format(format)
		return result
	}
}
func TimeBeforeTime(now string, format string, lessHour int, toTimestemp bool) string {
	times, _ := time.Parse(format, now)
	less, _ := strconv.ParseInt(strconv.Itoa(60*60*lessHour), 10, 64)
	if toTimestemp {
		result := strconv.FormatInt(times.Unix()-less, 10)
		return result
	} else {
		result := time.Unix(times.Unix()-less, 10).Format(format)
		return result
	}
}

func TimeFormatChange(times string, nowformat string, format string) string {
	timess, _ := time.Parse(nowformat, times)
	return time.Unix(timess.Unix(), 10).Format(format)
}

//时间
func Gettimes(format string) (timeres string) {
	timestemp := time.Now().Unix()
	tm := time.Unix(timestemp, 10)
	return tm.Format(format)
}

//时间戳
func Gettimestemp() (timestemp int64) {
	timestemp = time.Now().Unix()
	return
}
func Gettimestemps() (timestemps string) {
	timestemp := time.Now().Unix()
	timestemps = strconv.FormatInt(timestemp, 10)
	return
}

func Timestringtoint64(timestring string) (timeint int64) {
	var times time.Time
	var err error
	if len(timestring) == 19 {
		times, err = time.Parse("2006-01-02 15:04:05", timestring)
	} else if len(timestring) == 16 {
		times, err = time.Parse("2006-01-02 15:04", timestring)
	} else {
		times, err = time.Parse("2006-01-02", timestring)
	}
	if err != nil {
		beego.Debug("时间字符砖转换时间戳失败:", err)
	}
	return times.Unix()
}
