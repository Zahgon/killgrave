package http

import (
	"time"
)

// ResponseDelay represent time delay before server responds.
type ResponseDelay struct {
	delay  int64
	offset int64
}

// Delay return random time.Duration with respect to specified time range.
func (d *ResponseDelay) Delay() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// UnmarshalYAML of yaml.Unmarshaler interface.
// Input should be string, consisting of substring that can be parsed by time.ParseDuration,
// or two similar substrings seperated by ":".
func (d *ResponseDelay) UnmarshalYAML(unmarshal func(interface{}) error) error {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalJSON of json.Unmarshaler interface.
// Input should be string, consisting of substring that can be parsed by time.ParseDuration,
// or two similar substrings seperated by ":".
func (d *ResponseDelay) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (d *ResponseDelay) parseDelay(input string) error { _ = "STUB: not implemented"; return nil }
