package ews

import (
	"fmt"
	"time"
)

type DistinguishedFolderId struct {
	// List of values:
	// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/distinguishedfolderid
	Id string `xml:"Id,attr"`
}

type Time string

func (t Time) ToTime() (time.Time, error) {
	offset, err := getRFC3339Offset(time.Now())
	if err != nil {
		return time.Time{}, err
	}
	return time.Parse(time.RFC3339, string(t)+offset)

}

// return RFC3339 formatted offset, ex: +03:00 -03:30
func getRFC3339Offset(t time.Time) (string, error) {

	_, offset := t.Zone()
	i := int(float32(offset) / 36)

	sign := "+"
	if i < 0 {
		i = -i
		sign = "-"
	}
	hour := i / 100
	min := i % 100
	min = (60 * min) / 100

	return fmt.Sprintf("%s%02d:%02d", sign, hour, min), nil
}
