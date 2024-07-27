package custom_test

import (
	"encoding/json"
	"testing"
	"time"
	cs "wb/custom"
)

func TestTimeISO8601_MarshalJSON(t *testing.T) {
	iso8601Time := cs.TimeISO8601{Time: time.Date(2021, 12, 15, 10, 0, 0, 0, time.UTC)}
	expected := `"2021-12-15T10:00:00Z"`
	actual, err := json.Marshal(iso8601Time)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if string(actual) != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestTimeISO8601_UnmarshalJSON(t *testing.T) {
	data := []byte(`"2021-12-15T10:00:00Z"`)
	var iso8601Time cs.TimeISO8601
	err := json.Unmarshal(data, &iso8601Time)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := time.Date(2021, 12, 15, 10, 0, 0, 0, time.UTC)
	if !iso8601Time.Time.Equal(expected) {
		t.Errorf("Expected %v, got %v", expected, iso8601Time.Time)
	}
}

func TestTimeUnix_MarshalJSON(t *testing.T) {
	unixTime := cs.TimeUnix{Time: time.Unix(1639564800, 0)}
	expected := `1639564800`
	actual, err := json.Marshal(unixTime)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if string(actual) != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestTimeUnix_UnmarshalJSON(t *testing.T) {
	data := []byte(`1639564800`)
	var unixTime cs.TimeUnix
	err := json.Unmarshal(data, &unixTime)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := time.Unix(1639564800, 0)
	if !unixTime.Time.Equal(expected) {
		t.Errorf("Expected %v, got %v", expected, unixTime.Time)
	}
}
