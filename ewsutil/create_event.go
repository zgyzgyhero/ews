package ewsutil

import (
	"github.com/mhewedy/ews"
	"time"
)

// CreateEvent helper method to send Message
func CreateEvent(
	c *ews.Client, to []string, subject, body, location string, from time.Time, duration time.Duration,
) error {

	attendee := make([]ews.Attendee, len(to))
	for i, tt := range to {
		attendee[i] = ews.Attendee{Mailbox: ews.Mailbox{EmailAddress: tt}}
	}

	m := ews.CalendarItem{
		Subject: subject,
		Body: ews.Body{
			BodyType: "Text",
			Body:     []byte(body),
		},
		ReminderIsSet:              true,
		ReminderMinutesBeforeStart: 60,
		Start:                      from,
		End:                        from.Add(duration),
		IsAllDayEvent:              false,
		LegacyFreeBusyStatus:       "Busy",
		Location:                   location,
		RequiredAttendees:          []ews.Attendees{{Attendee: attendee}},
	}

	return ews.CreateCalendarItem(c, m)
}
