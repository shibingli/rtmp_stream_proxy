package utils

import (
	"fmt"
	"strconv"
	"time"
)

const (
	NanosecondLen  = 19
	MillisecondLen = 13
	SecondLen      = 10
)

func ConvertToTime(timestamp int64) (t time.Time, err error) {
	tmpStr := strconv.FormatInt(timestamp, 10)
	tmpStrLen := len(tmpStr)

	switch tmpStrLen {
	case NanosecondLen:
		t = time.Unix(0, timestamp)
		return
	case MillisecondLen:
		t = time.Unix(0, timestamp*int64(time.Millisecond))
		return
	case SecondLen:
		t = time.Unix(timestamp, 0)
		return
	}

	err = fmt.Errorf("%s", "Invalid parameter.")
	return
}
