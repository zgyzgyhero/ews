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
			StandardTime: TimeZoneTime{
				Bias:      0,
				Time:      "02:00:00",
				DayOrder:  5,
				Month:     10,
				DayOfWeek: "Sunday",
			},
			DaylightTime: TimeZoneTime{
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

func Test_unmarshal_GetUserAvailabilityResponse(t *testing.T) {

	soapResp := `
<?xml version="1.0" encoding="utf-8"?>
<s:Envelope
    xmlns:s="http://schemas.xmlsoap.org/soap/envelope/">
    <s:Header>
        <h:ServerVersionInfo MajorVersion="15" MinorVersion="20" MajorBuildNumber="2495" MinorBuildNumber="21" Version="V2018_01_08"
            xmlns:h="http://schemas.microsoft.com/exchange/services/2006/types"
            xmlns:xsd="http://www.w3.org/2001/XMLSchema"
            xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"/>
        </s:Header>
        <s:Body>
            <GetUserAvailabilityResponse
                xmlns="http://schemas.microsoft.com/exchange/services/2006/messages"
                xmlns:xsd="http://www.w3.org/2001/XMLSchema"
                xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
                <FreeBusyResponseArray>
                    <FreeBusyResponse>
                        <ResponseMessage ResponseClass="Error">
                            <MessageText>Microsoft.Exchange.InfoWorker.Common.Availability.MailRecipientNotFoundException: Unable to resolve e-mail address someone@ExServer.example.com to an Active Directory object.&#xD;
. Name of the server where exception originated: PR3PR01MB6473. LID: 57660</MessageText>
                            <ResponseCode>ErrorMailRecipientNotFound</ResponseCode>
                            <DescriptiveLinkKey>0</DescriptiveLinkKey>
                            <MessageXml>
                                <ExceptionType
                                    xmlns="http://schemas.microsoft.com/exchange/services/2006/errors">MailRecipientNotFoundException
                                </ExceptionType>
                                <ExceptionCode
                                    xmlns="http://schemas.microsoft.com/exchange/services/2006/errors">5009
                                </ExceptionCode>
                                <ExceptionServerName
                                    xmlns="http://schemas.microsoft.com/exchange/services/2006/errors">PR3PR01MB6473
                                </ExceptionServerName>
                                <ExceptionMessage
                                    xmlns="http://schemas.microsoft.com/exchange/services/2006/errors">Unable to resolve e-mail address someone@ExServer.example.com to an Active Directory object. LID: 57660</ExceptionMessage>
                            </MessageXml>
                        </ResponseMessage>
                        <FreeBusyView>
                            <FreeBusyViewType
                                xmlns="http://schemas.microsoft.com/exchange/services/2006/types">None
                            </FreeBusyViewType>
                        </FreeBusyView>
                    </FreeBusyResponse>
                </FreeBusyResponseArray>
            </GetUserAvailabilityResponse>
        </s:Body>
    </s:Envelope>
`

	var resp getUserAvailabilityResponseEnvelop
	err := xml.Unmarshal([]byte(soapResp), &resp)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t,
		ResponseClassError,
		resp.Body.GetUserAvailabilityResponse.FreeBusyResponseArray.FreeBusyResponse[0].
			ResponseMessage.ResponseClass,
	)

	assert.Equal(t,
		`Unable to resolve e-mail address someone@ExServer.example.com to an Active Directory object. LID: 57660`,
		resp.Body.GetUserAvailabilityResponse.FreeBusyResponseArray.FreeBusyResponse[0].
			ResponseMessage.MessageXml.ExceptionMessage,
	)

}
