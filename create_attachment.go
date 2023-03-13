package ews

import (
	"encoding/xml"
	"errors"
)

type CreateAttachment struct {
	XMLName      struct{}    `xml:"m:CreateAttachment"`
	ParentItemId ItemId      `xml:"m:ParentItemId"`
	Attachments  Attachments `xml:"m:Attachments"`
}

type CreateAttachmentResponseBodyEnvelop struct {
	XMLName struct{}                     `xml:"Envelope"`
	Body    CreateAttachmentResponseBody `xml:"Body"`
}

type CreateAttachmentResponseBody struct {
	CreateAttachmentResponse []createAttachmentResponseMessage `xml:"CreateAttachmentResponse>ResponseMessages>CreateAttachmentResponseMessage"`
}

type createAttachmentResponseMessage struct {
	ResponseCode     ResponseClass `xml:"ResponseCode"`
	MessageText      string        `xml:"MessageText,omitempty"`
	ResponseClass    string        `xml:"ResponseClass,attr"`
	FileAttachmentId AttachmentId  `xml:"Attachments>FileAttachment>AttachmentId,omitempty"`
}

type AttachmentId struct {
	Id                string `xml:"Id,attr"`
	RootItemId        string `xml:"RootItemId,attr"`
	RootItemChangeKey string `xml:"RootItemChangeKey,attr"`
}

// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/createattachment-operation
func SaveCreateAttachment(c Client, a *Attachments, itemId *ItemId) (*[]createAttachmentResponseMessage, error) {
	ca := CreateAttachment{
		ParentItemId: *itemId,
		Attachments:  *a,
	}
	xmlBytes, err := xml.MarshalIndent(ca, "", "  ")
	if err != nil {
		return nil, err
	}
	bb, err := c.SendAndReceive(xmlBytes)
	if err != nil {
		return nil, err
	}
	resp, err := checkCreateAttachmentResponse(bb)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func checkCreateAttachmentResponse(bb []byte) (*[]createAttachmentResponseMessage, error) {
	var soapResp CreateAttachmentResponseBodyEnvelop
	if err := xml.Unmarshal(bb, &soapResp); err != nil {
		return nil, err
	}
	resp := soapResp.Body.CreateAttachmentResponse
	if len(resp) > 0 && resp[0].ResponseCode == ResponseClassError {
		return nil, errors.New(resp[0].MessageText)
	}
	return &resp, nil
}
