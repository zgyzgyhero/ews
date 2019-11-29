package ews

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

func Test_marshal_GetUserAvailabilityRequest(t *testing.T) {

	mb := make([]MailboxData, 0)
	mb = append(mb, MailboxData{
		Email: Email{
			Name:        "",
			Address:     "someone@ExServer.example.com",
			RoutingType: "SMTP",
		},
		AttendeeType:     AttendeeTypeOrganizer,
		ExcludeConflicts: false,
	})

	start, _ := time.Parse(time.RFC3339, "2006-02-06T00:00:00Z")
	end, _ := time.Parse(time.RFC3339, "2006-02-25T23:59:59Z")

	req := GetUserAvailabilityRequest{
		TimeZone: TimeZone{
			Bias: 480,
			StandardTime: Time{
				Bias:      0,
				Time:      "02:00:00",
				DayOrder:  5,
				Month:     10,
				DayOfWeek: "Sunday",
			},
			DaylightTime: Time{
				Bias:      -60,
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
			MergedFreeBusyIntervalInMinutes: 60,
			RequestedView:                   RequestedViewFreeBusyMerged,
		},
	}

	xmlBytes, err := xml.MarshalIndent(req, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, `<m:GetUserAvailabilityRequest>
  <t:TimeZone>
    <t:Bias>480</t:Bias>
    <t:StandardTime>
      <t:Bias>0</t:Bias>
      <t:Time>02:00:00</t:Time>
      <t:DayOrder>5</t:DayOrder>
      <t:Month>10</t:Month>
      <t:DayOfWeek>Sunday</t:DayOfWeek>
    </t:StandardTime>
    <t:DaylightTime>
      <t:Bias>-60</t:Bias>
      <t:Time>02:00:00</t:Time>
      <t:DayOrder>1</t:DayOrder>
      <t:Month>4</t:Month>
      <t:DayOfWeek>Sunday</t:DayOfWeek>
    </t:DaylightTime>
  </t:TimeZone>
  <m:MailboxDataArray>
    <t:MailboxData>
      <t:Email>
        <t:Name></t:Name>
        <t:Address>someone@ExServer.example.com</t:Address>
        <t:RoutingType>SMTP</t:RoutingType>
      </t:Email>
      <t:AttendeeType>Organizer</t:AttendeeType>
      <t:ExcludeConflicts>false</t:ExcludeConflicts>
    </t:MailboxData>
  </m:MailboxDataArray>
  <t:FreeBusyViewOptions>
    <t:TimeWindow>
      <t:StartTime>2006-02-06T00:00:00Z</t:StartTime>
      <t:EndTime>2006-02-25T23:59:59Z</t:EndTime>
    </t:TimeWindow>
    <t:MergedFreeBusyIntervalInMinutes>60</t:MergedFreeBusyIntervalInMinutes>
    <t:RequestedView>FreeBusyMerged</t:RequestedView>
  </t:FreeBusyViewOptions>
</m:GetUserAvailabilityRequest>`, string(xmlBytes))
}
