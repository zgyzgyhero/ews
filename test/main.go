package main

import (
	"fmt"
	"github.com/mhewedy/ews"
	"github.com/mhewedy/ews/ewsutil"
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

	//err := testCreateCalendarItem(c)

	//err := testGetUserAvailability(c)

	err := testGetAvailability(c)

	if err != nil {
		log.Fatal("err>: ", err.Error())
	}

	fmt.Println("--- success ---")
}

func testSendEmail(c *ews.Client) error {
	return ewsutil.SendEmail(c,
		[]string{"mhewedy@gmail.com", "someone@else.com"},
		"An email subject",
		"The email body, as plain text",
	)
}

func testCreateCalendarItem(c *ews.Client) error {
	attendee := make([]ews.Attendee, 0)
	attendee = append(attendee,
		ews.Attendee{Mailbox: ews.Mailbox{EmailAddress: "mhewedy@mhewedy.onmicrosoft.com"}},
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
		Start:                      time.Now().Add(24 * time.Hour),
		End:                        time.Now().Add(24 * time.Hour).Add(30 * time.Minute),
		IsAllDayEvent:              false,
		LegacyFreeBusyStatus:       "Busy",
		Location:                   "Conference Room 721",
		RequiredAttendees:          attendees,
	})
}

func testGetUserAvailability(c *ews.Client) error {

	mb := make([]ews.MailboxData, 0)
	mb = append(mb, ews.MailboxData{
		Email: ews.Email{
			Name:        "",
			Address:     "mhewedy@mhewedy.onmicrosoft.com",
			RoutingType: "SMTP",
		},
		AttendeeType:     ews.AttendeeTypeRequired,
		ExcludeConflicts: false,
	}, ews.MailboxData{
		Email: ews.Email{
			Name:        "",
			Address:     "example2@mhewedy.onmicrosoft.com",
			RoutingType: "SMTP",
		},
		AttendeeType:     ews.AttendeeTypeRoom,
		ExcludeConflicts: false,
	})

	start, _ := time.Parse(time.RFC3339, "2019-11-28T00:00:00Z")
	end, _ := time.Parse(time.RFC3339, "2019-12-10T00:00:00Z")

	req := &ews.GetUserAvailabilityRequest{
		TimeZone: ews.TimeZone{
			Bias: 180,
			StandardTime: ews.TimeZoneTime{
				Bias:      0,
				Time:      "02:00:00",
				DayOrder:  5,
				Month:     10,
				DayOfWeek: "Sunday",
			},
			DaylightTime: ews.TimeZoneTime{
				Bias:      0,
				Time:      "02:00:00",
				DayOrder:  1,
				Month:     4,
				DayOfWeek: "Sunday",
			},
		},
		MailboxDataArray: ews.MailboxDataArray{MailboxData: mb},
		FreeBusyViewOptions: ews.FreeBusyViewOptions{
			TimeWindow: ews.TimeWindow{
				StartTime: start,
				EndTime:   end,
			},
			MergedFreeBusyIntervalInMinutes: 60,
			RequestedView:                   ews.RequestedViewFreeBusy,
		},
	}

	response, err := ews.GetUserAvailability(c, req)
	if err != nil {
		return err
	}

	fmt.Println(response)

	return nil
}

func testGetAvailability(c *ews.Client) error {

	eventUsers := []ewsutil.EventUser{
		{
			Email:        "mhewedy@mhewedy.onmicrosoft.com",
			AttendeeType: ews.AttendeeTypeRequired,
		},
		{
			Email:        "example2@mhewedy.onmicrosoft.com",
			AttendeeType: ews.AttendeeTypeRequired,
		},
	}
	start, _ := time.Parse(time.RFC3339, "2019-11-29T00:00:00Z")

	events, err := ewsutil.GetAvailability(c, eventUsers, start, 24*time.Hour)

	if err != nil {
		return err
	}

	fmt.Println(events)

	return nil
}
