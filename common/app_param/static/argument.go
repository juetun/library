package static

import (
	"fmt"
	"github.com/juetun/base-wrapper/lib/utils"
	"strconv"
	"time"
)

const (
	DataTypeDate = iota + 1
	DataTypeHour
	DataTypeMonth
	DataTypeYear
)

type ArgStaticChartWithTime struct {
	DateType       int      `json:"date_type" form:"date_type"` // 1:年月日;2:年月日 时; 3:年月;4:年
	DateFrom       string   `json:"date_from" form:"date_from"` // 起始日期
	DateTo         string   `json:"date_to" form:"date_to"`     // 终止日期
	DateTypeString string   `json:"-"`                          // 1:年月日;2:年月日 时; 3:年月;4:年
	DateArea       []string `json:"-"`
}

// GetTimeArea
func (r *ArgStaticChartWithTime) getTimeArea() (res []string, err error) {
	res = make([]string, 0, 100)
	var (
		timeFrom, timeTo, timeCurrent time.Time
	)
	if timeFrom, err = utils.DateParse(r.DateFrom, utils.DateTimeGeneral); err != nil {
		return
	}
	if timeTo, err = utils.DateParse(r.DateTo, utils.DateTimeGeneral); err != nil {
		return
	}
	timeTo = timeTo.Add(-1 * time.Second)
	switch r.DateType {
	case DataTypeDate:
		res = append(res, timeFrom.Format(utils.DateGeneral))
		timeCurrent = timeFrom
		for {
			if timeCurrent = timeCurrent.Add(24 * time.Hour); timeCurrent.After(timeTo) {
				break
			}
			res = append(res, timeCurrent.Format(utils.DateGeneral))
		}

	case DataTypeHour:
		const timeFormat = "2006.01.02 15"
		res = append(res, timeFrom.Format(timeFormat))
		timeCurrent = timeFrom
		for {
			if timeCurrent = timeCurrent.Add(time.Hour); timeCurrent.After(timeTo) {
				break
			}
			res = append(res, timeCurrent.Format(timeFormat))
		}
	case DataTypeMonth:
		const timeFormat = "2006.01"
		res = append(res, timeFrom.Format(timeFormat))
		timeCurrent = timeFrom
		for {
			if timeCurrent, err = r.getNextTime(timeCurrent, r.DateType); err != nil || timeCurrent.After(timeTo) {
				break
			}
			res = append(res, timeCurrent.Format(timeFormat))
		}

	case DataTypeYear:
		const timeFormat = "2006"
		res = append(res, timeFrom.Format(timeFormat))
		timeCurrent = timeFrom
		for {
			if timeCurrent, err = r.getNextTime(timeCurrent, r.DateType); err != nil || timeCurrent.After(timeTo) {
				break
			}
			res = append(res, timeCurrent.Format(timeFormat))
		}

	default:
		err = fmt.Errorf("当前不支持您要查看的数据类型")
		return
	}

	return
}

func (r *ArgStaticChartWithTime) getNextTime(timeCurrent time.Time, dataType int) (timeRes time.Time, err error) {
	switch dataType {
	case DataTypeMonth:
		var (
			yearNum, monthNum int
			year              = timeCurrent.Format("2006")
			month             = timeCurrent.Format("01")
		)

		if yearNum, err = strconv.Atoi(year); err != nil {
			return
		}
		if monthNum, err = strconv.Atoi(month); err != nil {
			return
		}
		if monthNum >= 12 {
			monthNum = 1
			yearNum = yearNum + 1
		} else {
			monthNum += 1
		}
		timeRes = time.Date(yearNum, time.Month(monthNum), 1, 0, 0, 0, 0, time.Local)
	case DataTypeYear:
		var (
			yearNum int
		)
		year := timeCurrent.Format("2006")
		if yearNum, err = strconv.Atoi(year); err != nil {
			return
		}
		timeRes = time.Date(yearNum, 1, 1, 0, 0, 0, 0, time.Local)
	default:
		err = fmt.Errorf("当前不支持你选择的时间类型")
	}
	return
}

func (r *ArgStaticChartWithTime) SetDefaultFrom(from string) (res *ArgStaticChartWithTime) {
	res = r
	if r.DateFrom == "" {
		r.DateFrom = from
	}
	return
}

func (r *ArgStaticChartWithTime) getTimeString(timeStamp time.Time, formatString string, suffixs ...string) (res string) {
	var suffix string
	if len(suffixs) > 0 {
		suffix = suffixs[0]
	}
	if suffix == "" {
		res = timeStamp.Format(formatString)
		return
	}
	res = timeStamp.Format(formatString) + suffix
	return

}

func (r *ArgStaticChartWithTime) Default() (err error) {
	if r.DateType == 0 {
		r.DateType = 1
	}

	switch r.DateType {
	case DataTypeDate:
		if r.DateTo == "" {
			r.DateTo = r.getTimeString(time.Now(), "2006-01-02", " 00:00:00")
		}
		if r.DateFrom == "" {
			r.DateFrom = r.getTimeString(time.Now().Add(-15*24*time.Hour), "2006-01-02", " 00:00:00")
		}
		r.DateTypeString = "%Y-%m-%d"
	case DataTypeHour:
		if r.DateTo == "" {
			//2006-01-02 15:04:05
			r.DateTo = r.getTimeString(time.Now(), "2006-01-02 15", "00:00")

		}
		if r.DateFrom == "" {
			r.DateFrom = r.getTimeString(time.Now().Add(-15*time.Hour), "2006-01-02 15", "00:00")
		}
		r.DateTypeString = "%Y-%m-%d %H"
	case DataTypeMonth:
		if r.DateTo == "" {
			r.DateTo = r.getTimeString(time.Now(), "2006-01", "-02 00:00:00")
		}
		if r.DateFrom == "" {
			r.DateFrom = r.getTimeString(time.Now().AddDate(0, -15, 0), "2006-01", "-02 00:00:00")
		}
		r.DateTypeString = "%Y-%m"
	case DataTypeYear:
		if r.DateTo == "" {
			r.DateTo = r.getTimeString(time.Now(), "2006", "-01-02 00:00:00")
		}
		if r.DateFrom == "" {
			r.DateFrom = r.getTimeString(time.Now().AddDate(-15, 0, 0), "2006", "-01-02 00:00:00")
		}
		r.DateTypeString = "%Y"
	default:
		err = fmt.Errorf("当前不支持你输入的类型统计(%d)", r.DateType)
	}
	if r.DateArea, err = r.getTimeArea(); err != nil {
		return
	}
	return
}
