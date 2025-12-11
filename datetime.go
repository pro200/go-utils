package utils

import (
	"fmt"
	"strings"
	"time"
)

// "19771029"형식 받아 time.Time으로 변환
func ParseTime(str string) (time.Time, error) {
	// 지원할 날짜 포맷들 (Go는 레퍼런스 타임 "Mon Jan 2 15:04:05 MST 2006" 기반)
	layouts := []string{
		"060102",              // yymmdd
		"060102150405",        // yymmddHHMMSS
		"20060102150405",      // yyyymmddHHMMSS
		"2006-01-02",          // yyyy-mm-dd
		"2006/01/02",          // yyyy/mm/dd
		"20060102",            // yyyymmdd
		"2006-01-02 15:04:05", // yyyy-mm-dd HH:MM:SS
		time.RFC3339,          // 2006-01-02T15:04:05Z07:00
	}

	var t time.Time
	var err error
	for _, layout := range layouts {
		t, err = time.Parse(layout, str)
		if err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("지원되지 않는 날짜 형식: %s", str)
}

func TimeFormat(t time.Time, layout string) string {
	replacer := strings.NewReplacer(
		"%Y", "2006", // 4자리 연도
		"%y", "06", // 2자리 연도
		"%m", "01", // 월
		"%d", "02", // 일
		"%H", "15", // 시 (00~23)
		"%I", "03", // 시 (01~12)
		"%M", "04", // 분
		"%S", "05", // 초
	)
	layout = replacer.Replace(layout)

	return t.Format(layout)
}

// 날짜 문자열을 받아 time.Time으로 변환 후 layout으로 포멧팅
// "060102", "20060102", "060102150405", "20060102150405" 중 하나의 형식을 value로 받음
func ParseTimeFormat(str, layout string) (string, error) {
	t, err := ParseTime(str)
	if err != nil {
		return "", err
	}

	return TimeFormat(t, layout), nil
}
