package ews

import (
	"encoding/xml"
)

// https://msdn.microsoft.com/en-us/library/office/aa563009(v=exchg.140).aspx

type CreateItem struct {
	XMLName            struct{}          `xml:"m:CreateItem"`
	MessageDisposition string            `xml:"MessageDisposition,attr"`
	SavedItemFolderId  SavedItemFolderId `xml:"m:SavedItemFolderId"`
	Items              Messages          `xml:"m:Items"`
}

type Messages struct {
	Message []Message `xml:"t:Message"`
}

type SavedItemFolderId struct {
	DistinguishedFolderId DistinguishedFolderId `xml:"t:DistinguishedFolderId"`
}

type DistinguishedFolderId struct {
	Id string `xml:"Id,attr"`
}

type Message struct {
	ItemClass    string     `xml:"t:ItemClass"`
	Subject      string     `xml:"t:Subject"`
	Body         Body       `xml:"t:Body"`
	Sender       OneMailbox `xml:"t:Sender"`
	ToRecipients XMailbox   `xml:"t:ToRecipients"`
}

type Body struct {
	BodyType string `xml:"BodyType,attr"`
	Body     []byte `xml:",chardata"`
}

type OneMailbox struct {
	Mailbox Mailbox `xml:"t:Mailbox"`
}

type XMailbox struct {
	Mailbox []Mailbox `xml:"t:Mailbox"`
}

type Mailbox struct {
	EmailAddress string `xml:"t:EmailAddress"`
}

// CreateMessageItem
// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/createitem-operation-email-message
func CreateMessageItem(c *Client, item *CreateItem) error {
	xmlBytes, err := xml.MarshalIndent(item, "", "  ")
	if err != nil {
		return err
	}

	_, err = c.sendAndReceive(xmlBytes)
	if err != nil {
		return err
	}
	return nil
}

// SendEmail helper method to send Message
func SendEmail(c *Client, to []string, subject, body string) error {

	item := &CreateItem{
		MessageDisposition: "SendAndSaveCopy",
		SavedItemFolderId:  SavedItemFolderId{DistinguishedFolderId{Id: "sentitems"}},
	}
	m := &Message{
		ItemClass: "IPM.Note",
		Subject:   subject,
		Body: Body{
			BodyType: "Text",
			Body:     []byte(body),
		},
		Sender: OneMailbox{
			Mailbox: Mailbox{
				EmailAddress: c.Username,
			},
		},
	}
	mb := make([]Mailbox, len(to))
	for i, addr := range to {
		mb[i].EmailAddress = addr
	}
	m.ToRecipients.Mailbox = append(m.ToRecipients.Mailbox, mb...)
	item.Items.Message = append(item.Items.Message, *m)

	return CreateMessageItem(c, item)
}
