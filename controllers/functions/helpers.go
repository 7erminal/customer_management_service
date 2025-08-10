package functions

import (
	"crypto/rand"
	"io"
	"time"

	"github.com/beego/beego/v2/core/logs"
)

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func EncodeToString(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

func ConvertStringToDateTime(date_ string) time.Time {
	var tDateTime time.Time

	var allowedDateList [6]string = [6]string{"2006-01-02", "2006/01/02", "2006-01-02 15:04:05.000", "2006/01/02 15:04:05.000", "2006-01-02T15:04:05.000Z", "2006-01-02 15:04:05.000000 -0700 MST"}

	for _, cdate_ := range allowedDateList {
		logs.Debug("About to convert ", date_)
		// Convert dob string to date
		tdt, error := time.Parse(cdate_, date_)

		if error != nil {
			logs.Error("Error parsing date", error)
		} else {
			logs.Error("Date converted to time successfully", tdt)
			tDateTime = tdt
			break
		}
	}

	return tDateTime
}
