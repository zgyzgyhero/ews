package ews

import "encoding/xml"

type BaseShape string

const (
	BaseShapeIdOnly        BaseShape = "IdOnly"
	BaseShapeDefault       BaseShape = "Default"
	BaseShapeAllProperties BaseShape = "AllProperties"
)

type BasePoint string

const (
	BasePointBeginning BasePoint = "Beginning"
	BasePointEnd       BasePoint = "End"
)

type FilePeopleRequest struct {
	XMLName             struct{}            `xml:"m:FindPeople"`
	PersonaShape        *PersonaShape       `xml:"m:PersonaShape,omitempty"`
	IndexedPageItemView IndexedPageItemView `xml:"m:IndexedPageItemView"`
	ParentFolderId      ParentFolderId      `xml:"m:ParentFolderId"`
	QueryString         string              `xml:"m:QueryString,omitempty"`
	// add additional fields
}

type PersonaShape struct {
	BaseShape            BaseShape              `xml:"t:BaseShape,omitempty"`
	AdditionalProperties []AdditionalProperties `xml:"t:AdditionalProperties,omitempty"`
}

type AdditionalProperties struct {
	FieldURI FieldURI `xml:"t:FieldURI,omitempty"`
	// add additional fields
}

type FieldURI struct {
	// List of possible values:
	// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/fielduri
	FieldURI string `xml:"FieldURI,attr,omitempty"`
}

type IndexedPageItemView struct {
	MaxEntriesReturned int       `xml:"MaxEntriesReturned,attr,omitempty"`
	Offset             int       `xml:"Offset,attr"`
	BasePoint          BasePoint `xml:"BasePoint,attr"`
}

type ParentFolderId struct {
	DistinguishedFolderId DistinguishedFolderId `xml:"t:DistinguishedFolderId"`
}

// GetUserAvailability
//https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/findpeople-operation
func FindPeople(c *Client, r *FilePeopleRequest) ([]byte, error) {

	xmlBytes, err := xml.MarshalIndent(r, "", "  ")
	if err != nil {
		return nil, err
	}

	bb, err := c.sendAndReceive(xmlBytes)
	if err != nil {
		return nil, err
	}

	return bb, nil
}

/*
Error:

<FindPeopleResponse ResponseClass="Error"
	xmlns="http://schemas.microsoft.com/exchange/services/2006/messages"
	xmlns:xsd="http://www.w3.org/2001/XMLSchema"
	xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
	<MessageText>The distinguished folder name is unrecognized.</MessageText>
	<ResponseCode>ErrorInvalidOperation</ResponseCode>
	<DescriptiveLinkKey>0</DescriptiveLinkKey>
	<TotalNumberOfPeopleInView>0</TotalNumberOfPeopleInView>
	<FirstMatchingRowIndex>0</FirstMatchingRowIndex>
	<FirstLoadedRowIndex>0</FirstLoadedRowIndex>
</FindPeopleResponse>
*/
