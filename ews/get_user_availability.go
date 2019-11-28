package ews

import "time"

const (
	AttendeeTypeOrganizer = "Organizer"
	AttendeeTypeRequired  = "Required"
	AttendeeTypeOptional  = "Optional"
	AttendeeTypeRoom      = "Room"
	AttendeeTypeResource  = "Resource"
)

const (
	RequestedViewNone           = "None"
	RequestedViewMergedOnly     = "MergedOnly"
	RequestedViewFreeBusy       = "FreeBusy"
	RequestedViewFreeBusyMerged = "FreeBusyMerged"
	RequestedViewDetailed       = "Detailed"
	RequestedViewDetailedMerged = "DetailedMerged"
)

type GetUserAvailabilityRequest struct {
	XMLName             struct{}            `xml:"m:GetUserAvailabilityRequest"`
	TimeZone            TimeZone            `xml:"t:TimeZone"`
	MailboxDataArray    MailboxDataArray    `xml:"m:MailboxDataArray"`
	FreeBusyViewOptions FreeBusyViewOptions `xml:"t:FreeBusyViewOptions"`
}

type FreeBusyViewOptions struct {
	TimeWindow                      TimeWindow `xml:"t:TimeWindow"`
	MergedFreeBusyIntervalInMinutes int        `xml:"t:MergedFreeBusyIntervalInMinutes"`
	RequestedView                   string     `xml:"t:RequestedView"`
}

type TimeWindow struct {
	StartTime time.Time `xml:"t:StartTime"`
	EndTime   time.Time `xml:"t:EndTime"`
}

type TimeZone struct {
	Bias         int  `xml:"t:Bias"`
	StandardTime Time `xml:"t:StandardTime"`
	DaylightTime Time `xml:"t:DaylightTime"`
}

type Time struct {
	Bias      int    `xml:"t:Bias"`
	Time      string `xml:"t:Time"`
	DayOrder  int16  `xml:"t:DayOrder"`
	Month     int16  `xml:"t:Month"`
	DayOfWeek string `xml:"t:DayOfWeek"`
	Year      string `xml:"Year"`
}

type MailboxDataArray struct {
	MailboxData []MailboxData `xml:"t:MailboxData"`
}

type MailboxData struct {
	Email            Email  `xml:"t:Email"`
	AttendeeType     string `xml:"t:AttendeeType"`
	ExcludeConflicts bool   `xml:"t:ExcludeConflicts"`
}

type Email struct {
	Name        string `xml:"t:Name"`
	Address     string `xml:"t:Address"`
	RoutingType string `xml:"t:RoutingType"`
}

type GetUserAvailabilityResponse struct {
}

// GetUserAvailability
//https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/getuseravailability-operation
func GetUserAvailability(c *Client, r *GetUserAvailabilityRequest) (*GetUserAvailabilityResponse, error) {

	return nil, nil
}
