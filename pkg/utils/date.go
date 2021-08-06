package utils

import "time"

/**
 * @Author: ula
 * @Date: 2021/1/21
 */

func GetIntervalDays(start, end string) ([]string, error) {
	sDate, err := time.Parse("2006-01-02", start)
	if err != nil {
		return nil, err
	}
	eDate, err := time.Parse("2006-01-02", end)
	if err != nil {
		return nil, err
	}
	if eDate.Before(sDate) {
		sDate = eDate
		end = start
	}
	dateList := make([]string, 0)
	for sDate.Format("2006-01-02") != end {
		format := sDate.Format("2006-01-02")
		dateList = append(dateList, format)
		sDate = sDate.AddDate(0, 0, 1)
	}
	dateList = append(dateList, end)
	return dateList, nil
}

func GetStartAndEndDate() (string, string) {
	e := time.Now()
	s := e.AddDate(0, 0, -6)
	return s.Format("2006-01-02"), e.Format("2006-01-02")
}
