package ewsutil

import (
	"github.com/mhewedy/ews"
	"time"
)

// CreateEvent helper method to send Message
func CreateEvent(
	c ews.Client, to, optional []string, subject, body, location string, from time.Time, duration time.Duration,
) error {

	requiredAttendees := make([]ews.Attendee, len(to))
	for i, tt := range to {
		requiredAttendees[i] = ews.Attendee{Mailbox: ews.Mailbox{EmailAddress: tt}}
	}

	optionalAttendees := make([]ews.Attendee, len(optional))
	for i, tt := range optional {
		optionalAttendees[i] = ews.Attendee{Mailbox: ews.Mailbox{EmailAddress: tt}}
	}

	m := ews.CalendarItem{
		Subject: subject,
		Body: ews.Body{
			BodyType: "Text",
			Body:     []byte(body),
		},
		ReminderIsSet:              true,
		ReminderMinutesBeforeStart: 15,
		Start:                      from,
		End:                        from.Add(duration),
		IsAllDayEvent:              false,
		LegacyFreeBusyStatus:       ews.BusyTypeBusy,
		Location:                   location,
		RequiredAttendees:          []ews.Attendees{{Attendee: requiredAttendees}},
		OptionalAttendees:          []ews.Attendees{{Attendee: optionalAttendees}},
	}

	return ews.CreateCalendarItem(c, m)
}
