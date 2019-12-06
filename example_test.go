package ews_test

import (
	"fmt"
	. "github.com/mhewedy/ews"
	"github.com/mhewedy/ews/ewsutil"
	"io/ioutil"
	"log"
	"math"
	"os"
	"testing"
	"time"
)

func Test_Example(t *testing.T) {

	c := NewClient(
		"https://outlook.office365.com/EWS/Exchange.asmx",
		"example@mhewedy.onmicrosoft.com",
		"systemsystem@123",
		&Config{Dump: true},
	)

	//err := testSendEmail(c)

	//err := testCreateCalendarItem(c)

	//err := testGetUserAvailability(c)

	//err := testListUsersEvents(c)

	//err := testCreateEvent(c)

	//err := testGetRoomLists(c)

	//err := testFindPeople(c)

	err := testGetUserPhoto(c)

	if err != nil {
		log.Fatal("err>: ", err.Error())
	}

	fmt.Println("--- success ---")
}

func testSendEmail(c Client) error {
	return ewsutil.SendEmail(c,
		[]string{"mhewedy@gmail.com", "someone@else.com"},
		"An email subject",
		"The email body, as plain text",
	)
}

func testCreateCalendarItem(c Client) error {
	attendee := make([]Attendee, 0)
	attendee = append(attendee,
		Attendee{Mailbox: Mailbox{EmailAddress: "mhewedy@mhewedy.onmicrosoft.com"}},
	)
	attendees := make([]Attendees, 0)
	attendees = append(attendees, Attendees{Attendee: attendee})

	return CreateCalendarItem(c, CalendarItem{
		Subject: "Planning Meeting",
		Body: Body{
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

func testGetUserAvailability(c Client) error {

	mb := make([]MailboxData, 0)
	mb = append(mb, MailboxData{
		Email: Email{
			Name:        "",
			Address:     "mhewedy@mhewedy.onmicrosoft.com",
			RoutingType: "SMTP",
		},
		AttendeeType:     AttendeeTypeRequired,
		ExcludeConflicts: false,
	}, MailboxData{
		Email: Email{
			Name:        "",
			Address:     "example2@mhewedy.onmicrosoft.com",
			RoutingType: "SMTP",
		},
		AttendeeType:     AttendeeTypeRoom,
		ExcludeConflicts: false,
	})

	start, _ := time.Parse(time.RFC3339, "2019-11-28T00:00:00Z")
	end, _ := time.Parse(time.RFC3339, "2019-12-10T00:00:00Z")

	req := &GetUserAvailabilityRequest{
		TimeZone: TimeZone{
			Bias: -180,
			StandardTime: TimeZoneTime{
				Bias:      0,
				Time:      "02:00:00",
				DayOrder:  5,
				Month:     10,
				DayOfWeek: "Sunday",
			},
			DaylightTime: TimeZoneTime{
				Bias:      0,
				Time:      "02:00:00",
				DayOrder:  1,
				Month:     4,
				DayOfWeek: "Sunday",
			},
		},
		MailboxDataArray: MailboxDataArray{MailboxData: mb},
		FreeBusyViewOptions: FreeBusyViewOptions{
			TimeWindow: TimeWindow{
				StartTime: start,
				EndTime:   end,
			},
			MergedFreeBusyIntervalInMinutes: 30,
			RequestedView:                   RequestedViewFreeBusy,
		},
	}

	response, err := GetUserAvailability(c, req)
	if err != nil {
		return err
	}

	fmt.Println(response)

	return nil
}

func testListUsersEvents(c Client) error {

	eventUsers := []ewsutil.EventUser{
		{
			Email:        "mhewedy@mhewedy.onmicrosoft.com",
			AttendeeType: AttendeeTypeRequired,
		},
		{
			Email:        "example2@mhewedy.onmicrosoft.com",
			AttendeeType: AttendeeTypeRequired,
		},
	}
	start, _ := time.Parse(time.RFC3339, "2019-11-29T00:00:00Z")

	events, err := ewsutil.ListUsersEvents(c, eventUsers, start, 48*time.Hour)

	if err != nil {
		return err
	}

	fmt.Println(events)

	return nil
}

func testCreateEvent(c Client) error {

	return ewsutil.CreateEvent(c,
		[]string{"mhewedy@mhewedy.onmicrosoft.com", "example2@mhewedy.onmicrosoft.com"},
		"An Event subject",
		"An Event body, as plain text",
		"Room 55",
		time.Now().Add(24*time.Hour),
		30*time.Minute,
	)
}

func testGetRoomLists(c Client) error {
	response, err := GetRoomLists(c)
	if err != nil {
		return err
	}
	fmt.Println(response)

	return nil
}

func testFindPeople(c Client) error {

	req := &FindPeopleRequest{IndexedPageItemView: IndexedPageItemView{
		MaxEntriesReturned: math.MaxInt32,
		Offset:             0,
		BasePoint:          BasePointBeginning,
	}, ParentFolderId: ParentFolderId{
		DistinguishedFolderId: DistinguishedFolderId{Id: "directory"}},
		PersonaShape: &PersonaShape{BaseShape: BaseShapeIdOnly,
			AdditionalProperties: AdditionalProperties{
				FieldURI: []FieldURI{
					{FieldURI: "persona:DisplayName"},
					{FieldURI: "persona:Title"},
					{FieldURI: "persona:EmailAddress"},
				},
			}},
		QueryString: "ex",
	}

	resp, err := FindPeople(c, req)

	if err != nil {
		return err
	}

	fmt.Println(resp)

	return nil
}

func testGetUserPhoto(c Client) error {

	bytes, err := ewsutil.GetUserPhoto(c, "mhewedy@mhewedy.onmicrosoft.com")

	if err != nil {
		return err
	}

	err = ioutil.WriteFile("/tmp/file.png", bytes, os.ModePerm)
	fmt.Println("written to: /tmp/file.png")

	return err
}
