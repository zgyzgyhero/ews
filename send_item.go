package ews

import "encoding/xml"

type SendItem struct {
	XMLName          struct{} `xml:"m:SendItem"`
	SaveItemToFolder string     `xml:"m:SaveItemToFolder,attr"`
	ItemIds          ItemIds  `xml:"m:ItemIds"`
}

type ItemIds struct {
	ItemId []ItemId `xml:"t:ItemId"`
}

// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/senditem-operation
func SendSavedItem(c Client,itemId *ItemIds) error {
	item := SendItem{
		SaveItemToFolder: "false",
		ItemIds:        *itemId,
	}
	xmlBytes, err := xml.MarshalIndent(item, "", "  ")
	if err != nil {
		return err
	}
	bb, err := c.SendAndReceive(xmlBytes)
	if err != nil {
		return err
	}

	if _, err := checkCreateItemResponse(bb); err != nil {
		return err
	}

	return nil
}