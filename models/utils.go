package models

import (
	"bytes"
	"crypto/rand"
	"encoding/base32"
	"time"

	"github.com/rs/xid"
)

// GetMilliseconds a convenience method to get milliseconds since epoch.
func GetMilliseconds() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// GetMillisecondsForSpecificTime is a convenience method to get milliseconds since epoch for provided Time.
func GetMillisecondsForSpecificTime(thisTime time.Time) int64 {
	return thisTime.UnixNano() / int64(time.Millisecond)
}

// NewID is a globally unique identifier.
func NewID() string {
	return xid.New().String()
}

//NewRandomString is method to generate secured random string
func NewRandomString(length int) string {
	var encoding = base32.NewEncoding("qaz12wsx3edc4rfv5tgb6yhn7ujm8ik9olp")
	var b bytes.Buffer
	str := make([]byte, length+8)
	rand.Read(str)
	encoder := base32.NewEncoder(encoding, &b)
	encoder.Write(str)
	encoder.Close()
	b.Truncate(length) // removes the '==' padding
	return b.String()
}

func ConvertFromMilliseconds(ms int64) time.Time {
	return time.Unix(0, ms*int64(time.Millisecond))
}
