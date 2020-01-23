package go_jp_national_holiday

import (
	"encoding/csv"
	"io"
	"os"
	"time"
)

const dateFormat = "2020/1/1"

func IsBusinessDay(t time.Time) bool {
	// isHoliday
	if t.Weekday() == time.Saturday || t.Weekday() == time.Sunday {
		return false
	}

	// isNationalHoliday
	t = t.In(time.FixedZone("Asia/Tokyo", 9*60*60))
	csvFile, err := os.Open("nationalHoliday.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	r := csv.NewReader(csvFile)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return true
		}
		if t.Format(dateFormat) == record[0] {
			return false
		}
	}
	return true
}
