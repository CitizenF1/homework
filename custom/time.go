package custom

import (
	"encoding/json"
	"time"
)

type TimeISO8601 struct {
	time.Time
}

const iso8601Format = "2006-01-02T15:04:05Z07:00"

func (t *TimeISO8601) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	parsedTime, err := time.Parse(time.RFC3339, str)
	if err != nil {
		return err
	}
	t.Time = parsedTime
	return nil
}

func (t TimeISO8601) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Format(iso8601Format))
}

type TimeUnix struct {
	time.Time
}

func (t *TimeUnix) UnmarshalJSON(data []byte) error {
	var unixTime int64
	if err := json.Unmarshal(data, &unixTime); err != nil {
		return err
	}
	t.Time = time.Unix(unixTime, 0)
	return nil
}

func (t TimeUnix) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Time.Unix())
}
