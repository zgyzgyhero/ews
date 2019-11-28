package main

import (
	"fmt"
	"github.com/mhewedy/ews/ews"
	"log"
	"time"
)

func main() {

	c := ews.NewClientWithConfig(
		"https://outlook.office365.com/EWS/Exchange.asmx",
		"example@mhewedy.onmicrosoft.com",
		"systemsystem@123",
		&ews.Config{Dump: true},
	)

	//err := testSendEmail(c)

	err := testCreateCalendarItem(c)

	if err != nil {
		log.Fatal("err: ", err.Error())
	}

	fmt.Println("--- sent ---")
}

func testSendEmail(c *ews.Client) error {
	return ews.SendEmail(c,
		[]string{"mhewedy@gmail.com", "someone@else.com"},
		"An email subject",
		"The email body, as plain text",
	)
}

func testCreateCalendarItem(c *ews.Client) error {
	attendee := make([]ews.Attendee, 0)
	attendee = append(attendee,
		ews.Attendee{Mailbox: ews.Mailbox{EmailAddress: "mhewedy@gmail.com"}},
	)
	attendees := make([]ews.Attendees, 0)
	attendees = append(attendees, ews.Attendees{Attendee: attendee})

	return ews.CreateCalendarItem(c, ews.CalendarItem{
		Subject: "Planning Meeting",
		Body: ews.Body{
			BodyType: "Text",
			Body:     []byte("Plan the agenda for next week's meeting."),
		},
		ReminderIsSet:              true,
		ReminderMinutesBeforeStart: 60,
		Start:                      time.Now(),
		End:                        time.Now().Add(30 * time.Minute),
		IsAllDayEvent:              false,
		LegacyFreeBusyStatus:       "Busy",
		Location:                   "Conference Room 721",
		RequiredAttendees:          attendees,
	})
}
