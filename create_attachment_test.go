package ews

import (
	"encoding/xml"
	"fmt"
	"testing"
)

func Test_createAttachmentResponse(t *testing.T) {
	xmlStr := `<?xml version="1.0" encoding="utf-8"?><s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/"><s:Header><h:ServerVersionInfo MajorVersion="15" MinorVersion="1" MajorBuildNumber="2375" MinorBuildNumber="28" Version="V2017_07_11" xmlns:h="http://schemas.microsoft.com/exchange/services/2006/types" xmlns="http://schemas.microsoft.com/exchange/services/2006/types" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"/></s:Header><s:Body xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema"><m:CreateAttachmentResponse xmlns:m="http://schemas.microsoft.com/exchange/services/2006/messages" xmlns:t="http://schemas.microsoft.com/exchange/services/2006/types"><m:ResponseMessages><m:CreateAttachmentResponseMessage ResponseClass="Success"><m:ResponseCode>NoError</m:ResponseCode><m:Attachments><t:FileAttachment><t:AttachmentId Id="AAMkADExZTI1MjVmLWQ4MzEtNDBjYS1hMWEwLTYzMDIyNGQ5MWFjNwBGAAAAAAAsm1xVH2inSpA0JbGZEhm1BwBb0+EYLBFiQqC1D2X3n3odAAAA1zbaAABUGpQQZwqURZb0pqd+SpczAAJQbUulAAABEgAQAFmnqat27N5JqLjo/76gzxY=" RootItemId="AAMkADExZTI1MjVmLWQ4MzEtNDBjYS1hMWEwLTYzMDIyNGQ5MWFjNwBGAAAAAAAsm1xVH2inSpA0JbGZEhm1BwBb0+EYLBFiQqC1D2X3n3odAAAA1zbaAABUGpQQZwqURZb0pqd+SpczAAJQbUulAAA=" RootItemChangeKey="CQAAABYAAABUGpQQZwqURZb0pqd+SpczAAJQdXxM"/><t:LastModifiedTime>2023-03-13T07:23:46</t:LastModifiedTime><t:IsInline>false</t:IsInline></t:FileAttachment></m:Attachments></m:CreateAttachmentResponseMessage><m:CreateAttachmentResponseMessage ResponseClass="Success"><m:ResponseCode>NoError</m:ResponseCode><m:Attachments><t:FileAttachment><t:AttachmentId Id="AAMkADExZTI1MjVmLWQ4MzEtNDBjYS1hMWEwLTYzMDIyNGQ5MWFjNwBGAAAAAAAsm1xVH2inSpA0JbGZEhm1BwBb0+EYLBFiQqC1D2X3n3odAAAA1zbaAABUGpQQZwqURZb0pqd+SpczAAJQbUulAAABEgAQAOw7JLNg/1dGrQsh2ZFBmek=" RootItemId="AAMkADExZTI1MjVmLWQ4MzEtNDBjYS1hMWEwLTYzMDIyNGQ5MWFjNwBGAAAAAAAsm1xVH2inSpA0JbGZEhm1BwBb0+EYLBFiQqC1D2X3n3odAAAA1zbaAABUGpQQZwqURZb0pqd+SpczAAJQbUulAAA=" RootItemChangeKey="CQAAABYAAABUGpQQZwqURZb0pqd+SpczAAJQdXxM"/><t:LastModifiedTime>2023-03-13T07:23:46</t:LastModifiedTime><t:IsInline>false</t:IsInline></t:FileAttachment></m:Attachments></m:CreateAttachmentResponseMessage></m:ResponseMessages></m:CreateAttachmentResponse></s:Body></s:Envelope>`
	resp := CreateAttachmentResponseBodyEnvelop{}
	err := xml.Unmarshal([]byte(xmlStr), &resp)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(resp)
}