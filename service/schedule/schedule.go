package schedule

import (
	"errors"
	"github.com/xgpc/dsg/exce"
	"time"
)

type timerTime struct {
	// 是否设置过周期类型和值
	timeType string

	// 当类型是 HourlyAt、DailyAt、WeeklyAt、MonthlyAt、YearlyAt
	minute int
	hour   int
	week   string
	day    int
	month  string
}

type TimerData struct {
	// 调度周期
	timeTime timerTime

	// 调度重复：在一个周期内没执行完，新的周期是否要执行
	repeat bool

	// 调度函数
	handler func()

	// 任务是否正在执行
	isRunning bool
}

var timerList []*TimerData

const errStr = "定时器设置错误，请检查！"

// 设置定时任务
func NewSchedule(handler func()) *TimerData {
	var data = &TimerData{
		repeat:    false,
		handler:   handler,
		isRunning: false,
		timeTime:  timerTime{},
	}
	timerList = append(timerList, data)
	return data
}

// 设置是否重复，默认不重复，如果不重复的话，在周期内该函数尚未执行完毕，则此周期就不再执行
func (t *TimerData) SetRepeat(repeat bool) *TimerData {
	t.repeat = repeat
	return t
}

// 每分钟
func (t *TimerData) EveryMinute() {
	t.timeTime.timeType = "EveryMinute"
}

// 每五分钟
func (t *TimerData) EveryFiveMinutes() {
	t.timeTime.timeType = "EveryFiveMinutes"
}

// 每十分钟
func (t *TimerData) EveryTenMinutes() {
	t.timeTime.timeType = "EveryTenMinutes"
}

// 每半小时
func (t *TimerData) EveryThirtyMinutes() {
	t.timeTime.timeType = "EveryThirtyMinutes"
}

// 每小时的x分
/*
 *		minute:		[0,59]
 */
func (t *TimerData) HourlyAt(minute int) {
	if minute < 0 || minute > 59 {
		panic(errStr)
	}
	t.timeTime.timeType = "HourlyAt"
	t.timeTime.minute = minute
}

// 每天的x时x分
/*
 *		hour:		[0, 23]
 *		minute:		[0,59]
 */
func (t *TimerData) DailyAt(hour int, minute int) {
	if hour < 0 || hour > 23 {
		panic(errStr)
	}
	if minute < 0 || minute > 59 {
		panic(errStr)
	}
	t.timeTime.timeType = "DailyAt"
	t.timeTime.minute = minute
	t.timeTime.hour = hour
}

// 每周的星期几x时x分
/*
 *		week:		[0,6]	(0表示星期天)
 *		hour:		[0, 23]
 *		minute:		[0,59]
 */
func (t *TimerData) WeeklyAt(week int, hour int, minute int) {
	if week < 0 || week > 6 {
		panic(errStr)
	}
	if hour < 0 || hour > 23 {
		panic(errStr)
	}
	if minute < 0 || minute > 59 {
		panic(errStr)
	}
	t.timeTime.timeType = "WeeklyAt"
	t.timeTime.minute = minute
	t.timeTime.hour = hour
	t.timeTime.week = map[int]string{
		0: "Sunday",
		1: "Monday",
		2: "Tuesday",
		3: "Wednesday",
		4: "Thursday",
		5: "Friday",
		6: "Saturday",
	}[week]
}

// 每月的x日x时x分
/*
 *		day:		[1,31]
 *		hour:		[0, 23]
 *		minute:		[0,59]
 */
func (t *TimerData) MonthlyAt(day int, hour int, minute int) {
	if day < 1 || day > 31 {
		panic(errStr)
	}
	if hour < 0 || hour > 23 {
		panic(errStr)
	}
	if minute < 0 || minute > 59 {
		panic(errStr)
	}
	t.timeTime.timeType = "MonthlyAt"
	t.timeTime.minute = minute
	t.timeTime.hour = hour
	t.timeTime.day = day
}

// 每年的x月x日x时x分
/*
 *		month：		[1,12]
 *		day:		[1,31]
 *		hour:		[0, 23]
 *		minute:		[0,59]
 */
func (t *TimerData) YearlyAt(month int, day int, hour int, minute int) {
	if hour < 0 || hour > 23 {
		panic(errStr)
	}
	if minute < 0 || minute > 59 {
		panic(errStr)
	}
	if month < 1 || month > 12 {
		panic(errStr)
	}
	if day < 1 || day > 31 {
		panic(errStr)
	}
	for _, v := range []int{4, 6, 9, 11} {
		if month == v && day > 30 {
			panic(errStr)
		}
	}
	if month == 2 && day > 29 {
		panic(errStr)
	}
	if month == 2 && day == 29 {
		if day == 29 {
			var year = time.Now().Year()
			if year%100 == 0 {
				if year%400 != 0 {
					panic("世纪年，不能被400整除，不是闰年，二月没有29天")
				}
			} else {
				if year%4 != 0 {
					panic("普通年，不能被4整除，不是闰年，二月没有29天")
				}
			}
		}
	}

	t.timeTime.timeType = "YearlyAt"
	t.timeTime.minute = minute
	t.timeTime.hour = hour
	t.timeTime.day = day
	t.timeTime.month = map[int]string{
		1:  "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}[month]
}

// 开始执行任务
func StartSchedules() {

	// 校验是否每个任务都设置了周期
	for _, v := range timerList {
		if v.timeTime.timeType == "" {
			panic(errStr)
		}
	}

	go func() {

		// 首次执行，在下一分钟
		time.Sleep(time.Second * time.Duration(60-time.Now().Second()))

		for {
			// 每分钟的00秒，执行一次这里的逻辑，相当于linux的crontab最细到每秒执行一次
			var now = time.Now()
			var currSecond = now.Second()
			runMain(now)
			time.Sleep(time.Second * time.Duration(60-currSecond))
		}
	}()
}

func runMain(now time.Time) {
	for _, t := range timerList {

		if !t.isShowTime(now) {
			continue
		}

		t.myHandler()
	}
}

func (t *TimerData) isShowTime(now time.Time) bool {

	// 1. 任务不可重复，但此时正在执行的，就不能执行
	if !t.repeat && t.isRunning == true {
		return false
	}

	// 2. 是否符合执行的时间点
	switch t.timeTime.timeType {
	case "EveryMinute":
		return now.Minute()%1 == 0
	case "EveryFiveMinutes":
		return now.Minute()%5 == 0
	case "EveryTenMinutes":
		return now.Minute()%10 == 0
	case "EveryThirtyMinutes":
		return now.Minute()%30 == 0
	case "HourlyAt":
		return now.Minute() == t.timeTime.minute
	case "DailyAt":
		return now.Hour() == t.timeTime.hour && now.Minute() == t.timeTime.minute
	case "WeeklyAt":
		return now.Weekday().String() == t.timeTime.week && now.Hour() == t.timeTime.hour && now.Minute() == t.timeTime.minute
	case "MonthlyAt":
		return now.Day() == t.timeTime.day && now.Hour() == t.timeTime.hour && now.Minute() == t.timeTime.minute
	case "YearlyAt":
		return now.Month().String() == t.timeTime.month && now.Day() == t.timeTime.day && now.Hour() == t.timeTime.hour && now.Minute() == t.timeTime.minute
	}
	return false
}

func (t *TimerData) myHandler() {
	t.isRunning = true
	go func() {
		// 捕捉异常，避免业务失败导致整个服务停掉
		defer func() {
			if e := recover(); e != nil {
				switch v := e.(type) {
				case exce.SysException:
				case string:
					exce.ThrowSys(errors.New("定时器处理错误：" + v))
				case error:
					exce.ThrowSys(e.(error))
				default:
					exce.ThrowSys(errors.New("定时器处理错误"))
				}
			}
		}()

		defer func() {
			t.isRunning = false
		}()
		// 执行业务
		t.handler()
	}()
}
